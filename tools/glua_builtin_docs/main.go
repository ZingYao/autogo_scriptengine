package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const schema = "glua builtin function docs. Generated from AutoGo ScriptEngine static registrations."

type localizedText map[string]string
type localizedList map[string][]string

type builtinDoc struct {
	Signature   localizedText `json:"signature"`
	Returns     localizedText `json:"returns"`
	Params      localizedList `json:"params"`
	Description localizedText `json:"description"`
	Example     localizedText `json:"example"`
}

type builtinCatalog struct {
	Schema    string             `json:"_schema"`
	Functions orderedBuiltinDocs `json:"functions"`
}

type orderedBuiltinDocs map[string]builtinDoc

type report struct {
	GeneratedFiles []generatedFile `json:"generatedFiles"`
	Directories    []bucketReport  `json:"directories"`
	Conflicts      []conflict      `json:"conflicts,omitempty"`
	Unresolved     []unresolved    `json:"unresolved,omitempty"`
	Exclusions     []exclusion     `json:"exclusions,omitempty"`
	Fallbacks      []fallback      `json:"fallbacks,omitempty"`
}

type generatedFile struct {
	Name  string `json:"name"`
	Path  string `json:"path"`
	Count int    `json:"count"`
}

type bucketReport struct {
	Name       string `json:"name"`
	Source     string `json:"source"`
	Count      int    `json:"count"`
	Duplicates int    `json:"duplicates"`
}

type conflict struct {
	Name      string   `json:"name"`
	Bucket    string   `json:"bucket"`
	Kept      string   `json:"kept"`
	Discarded []string `json:"discarded"`
}

type unresolved struct {
	File   string `json:"file"`
	Line   int    `json:"line"`
	Reason string `json:"reason"`
}

type exclusion struct {
	File   string `json:"file"`
	Line   int    `json:"line"`
	Reason string `json:"reason"`
}

type fallback struct {
	File   string `json:"file"`
	Line   int    `json:"line"`
	Name   string `json:"name"`
	Reason string `json:"reason"`
}

type functionDoc struct {
	Name        string
	Description string
	Signature   signature
	Source      sourcePosition
	Bucket      string
	Kind        string
}

type signature struct {
	Params  []param
	Returns []param
	Unknown bool
}

type param struct {
	Name string
	Type string
}

type sourcePosition struct {
	File string
	Line int
}

type functionIndex struct {
	decls   map[string]*ast.FuncType
	methods map[string]*ast.FuncType
}

type sourceFile struct {
	path     string
	astFile  *ast.File
	constMap map[string]string
	imports  map[string]string
	funcs    functionIndex
}

type externalIndexes map[string]functionIndex

func main() {
	root := flag.String("root", ".", "项目根目录")
	outDir := flag.String("out", "docs/glua_builtin_docs", "输出目录")
	flag.Parse()

	if err := run(*root, *outDir); err != nil {
		fmt.Fprintf(os.Stderr, "generate glua builtin docs failed: %v\n", err)
		os.Exit(1)
	}
}

func run(root string, outDir string) error {
	root, err := filepath.Abs(root)
	if err != nil {
		return err
	}
	files, err := parseSourceFiles(root)
	if err != nil {
		return err
	}

	docs, _ := collectDocs(root, files, externalIndexes{})
	if err := os.MkdirAll(filepath.Join(root, outDir), 0o755); err != nil {
		return err
	}
	generated, err := writeCatalogs(filepath.Join(root, outDir), docs)
	if err != nil {
		return err
	}
	return writeReadme(filepath.Join(root, outDir, "README.md"), generated)
}

func parseSourceFiles(root string) ([]sourceFile, error) {
	var files []sourceFile
	err := filepath.WalkDir(filepath.Join(root, "lua_engine"), func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() {
			name := entry.Name()
			if name == "vendor" || strings.HasPrefix(name, ".") {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		fset := token.NewFileSet()
		parsed, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			return fmt.Errorf("parse %s: %w", rel, err)
		}
		constMap := collectStringConsts(parsed)
		funcs := collectFuncDecls(parsed)
		ast.Inspect(parsed, func(node ast.Node) bool {
			if node == nil {
				return true
			}
			return true
		})
		files = append(files, sourceFile{
			path:     rel,
			astFile:  parsed,
			constMap: constMap,
			imports:  collectImports(parsed),
			funcs:    funcs,
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	sort.Slice(files, func(i, j int) bool {
		return files[i].path < files[j].path
	})
	return files, nil
}

func collectStringConsts(file *ast.File) map[string]string {
	values := map[string]string{}
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.CONST {
			continue
		}
		for _, spec := range gen.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			for index, name := range valueSpec.Names {
				if index >= len(valueSpec.Values) {
					continue
				}
				if value, ok := stringValue(valueSpec.Values[index], values); ok {
					values[name.Name] = value
				}
			}
		}
	}
	return values
}

func collectFuncDecls(file *ast.File) functionIndex {
	index := functionIndex{decls: map[string]*ast.FuncType{}, methods: map[string]*ast.FuncType{}}
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Type == nil {
			continue
		}
		if fn.Recv == nil || len(fn.Recv.List) == 0 {
			index.decls[fn.Name.Name] = fn.Type
			continue
		}
		recvName := receiverTypeName(fn.Recv.List[0].Type)
		if recvName != "" {
			index.methods[recvName+"."+fn.Name.Name] = fn.Type
		}
	}
	return index
}

func collectImports(file *ast.File) map[string]string {
	imports := map[string]string{}
	for _, item := range file.Imports {
		path, err := strconv.Unquote(item.Path.Value)
		if err != nil {
			continue
		}
		name := filepath.Base(path)
		if item.Name != nil && item.Name.Name != "." && item.Name.Name != "_" {
			name = item.Name.Name
		}
		imports[name] = path
	}
	return imports
}

func receiverTypeName(expr ast.Expr) string {
	switch value := expr.(type) {
	case *ast.Ident:
		return value.Name
	case *ast.StarExpr:
		return receiverTypeName(value.X)
	case *ast.SelectorExpr:
		return value.Sel.Name
	case *ast.ParenExpr:
		return receiverTypeName(value.X)
	default:
		return ""
	}
}

func collectDocs(root string, files []sourceFile, external externalIndexes) (map[string][]functionDoc, report) {
	result := map[string][]functionDoc{
		"common":    {},
		"android":   {},
		"ios":       {},
		"lrappsoft": {},
	}
	rep := report{}
	fset := token.NewFileSet()

	for _, src := range files {
		bucket := bucketForPath(src.path)
		ast.Inspect(src.astFile, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}
			position := positionForNode(fset, root, src.path, src.astFile, call)
			switch {
			case isRegisterMethodCall(call):
				doc, ok, reason := buildRegisterMethodDoc(src, external, call, bucket, position)
				if !ok {
					if excluded, excludeReason := excludedDynamicRegistration(src.path, position.Line, reason); excluded {
						rep.Exclusions = append(rep.Exclusions, exclusion{File: src.path, Line: position.Line, Reason: excludeReason})
						return true
					}
					rep.Unresolved = append(rep.Unresolved, unresolved{File: src.path, Line: position.Line, Reason: reason})
					return true
				}
				if doc.Signature.Unknown {
					rep.Fallbacks = append(rep.Fallbacks, fallback{File: src.path, Line: position.Line, Name: doc.Name, Reason: "无法静态解析 Go 函数签名，使用 any 回退"})
				}
				result[bucket] = append(result[bucket], doc)
			case isRegisterValueCall(call):
				doc, ok, reason := buildRegisterValueDoc(src, call, bucket, position)
				if !ok {
					rep.Unresolved = append(rep.Unresolved, unresolved{File: src.path, Line: position.Line, Reason: reason})
					return true
				}
				result[bucket] = append(result[bucket], doc)
			case isGLuaRegisterCall(call):
				doc, ok, reason := buildGLuaRegisterDoc(src, external, call, bucket, position)
				if !ok {
					rep.Unresolved = append(rep.Unresolved, unresolved{File: src.path, Line: position.Line, Reason: reason})
					return true
				}
				result[bucket] = append(result[bucket], doc)
			case isInstallCoreTableCall(call):
				docs, ok, reason := buildCoreTableDocs(src, external, call, bucket, position)
				if !ok {
					rep.Unresolved = append(rep.Unresolved, unresolved{File: src.path, Line: position.Line, Reason: reason})
					return true
				}
				result[bucket] = append(result[bucket], docs...)
			}
			return true
		})
	}

	for bucket, items := range result {
		merged, duplicates, conflicts := dedupeDocs(bucket, items)
		result[bucket] = merged
		rep.Directories = append(rep.Directories, bucketReport{
			Name:       bucket,
			Source:     sourceDescription(bucket),
			Count:      len(merged),
			Duplicates: duplicates,
		})
		rep.Conflicts = append(rep.Conflicts, conflicts...)
	}
	sort.Slice(rep.Directories, func(i, j int) bool {
		return rep.Directories[i].Name < rep.Directories[j].Name
	})
	sort.Slice(rep.Unresolved, func(i, j int) bool {
		if rep.Unresolved[i].File == rep.Unresolved[j].File {
			return rep.Unresolved[i].Line < rep.Unresolved[j].Line
		}
		return rep.Unresolved[i].File < rep.Unresolved[j].File
	})
	sort.Slice(rep.Exclusions, func(i, j int) bool {
		if rep.Exclusions[i].File == rep.Exclusions[j].File {
			return rep.Exclusions[i].Line < rep.Exclusions[j].Line
		}
		return rep.Exclusions[i].File < rep.Exclusions[j].File
	})
	sort.Slice(rep.Fallbacks, func(i, j int) bool {
		if rep.Fallbacks[i].File == rep.Fallbacks[j].File {
			return rep.Fallbacks[i].Line < rep.Fallbacks[j].Line
		}
		return rep.Fallbacks[i].File < rep.Fallbacks[j].File
	})
	return result, rep
}

func excludedDynamicRegistration(path string, line int, reason string) (bool, string) {
	if reason == "" {
		return false, ""
	}
	switch {
	case path == "lua_engine/lua_engine.go" && (line == 274 || line == 697):
		return true, "运行时由宿主或 Lua 脚本动态注册，builtin docs 不收录用户动态方法。"
	case strings.Contains(path, "/imgui/imgui_reflect_bridge.go"):
		return true, "ImGui 反射桥循环注册入口已由生成的具体 imgui.* 注册项覆盖。"
	case path == "lua_engine/model/lrappsoft/dynamicui/dynamicui_inject.go":
		return true, "dynamicui 占位循环方法按运行时兼容策略注册，不作为静态 builtin docs 条目。"
	default:
		return false, ""
	}
}

func bucketForPath(path string) string {
	switch {
	case strings.HasPrefix(path, "lua_engine/model/autogo_ios/"):
		return "ios"
	case strings.HasPrefix(path, "lua_engine/model/autogo/"):
		return "android"
	case strings.HasPrefix(path, "lua_engine/model/lrappsoft/"):
		return "lrappsoft"
	default:
		return "common"
	}
}

func sourceDescription(bucket string) string {
	switch bucket {
	case "android":
		return "lua_engine/model/autogo"
	case "ios":
		return "lua_engine/model/autogo_ios"
	case "lrappsoft":
		return "lua_engine/model/lrappsoft"
	default:
		return "lua_engine core"
	}
}

func positionForNode(_ *token.FileSet, root string, rel string, file *ast.File, node ast.Node) sourcePosition {
	content, err := os.ReadFile(filepath.Join(root, rel))
	if err != nil {
		return sourcePosition{File: rel, Line: 0}
	}
	offset := int(node.Pos() - file.Pos())
	line := 1
	if offset > 0 && offset <= len(content) {
		line = bytes.Count(content[:offset], []byte("\n")) + 1
	}
	return sourcePosition{File: rel, Line: line}
}

func isRegisterMethodCall(call *ast.CallExpr) bool {
	switch fn := call.Fun.(type) {
	case *ast.SelectorExpr:
		return fn.Sel.Name == "RegisterMethod"
	case *ast.Ident:
		return fn.Name == "RegisterMethod"
	default:
		return false
	}
}

func isRegisterValueCall(call *ast.CallExpr) bool {
	fn, ok := call.Fun.(*ast.SelectorExpr)
	return ok && fn.Sel.Name == "RegisterValue"
}

func isGLuaRegisterCall(call *ast.CallExpr) bool {
	fn, ok := call.Fun.(*ast.SelectorExpr)
	return ok && fn.Sel.Name == "Register" && selectorRoot(fn.X) == "glua"
}

func isInstallCoreTableCall(call *ast.CallExpr) bool {
	fn, ok := call.Fun.(*ast.SelectorExpr)
	return ok && fn.Sel.Name == "installVMCoreTable"
}

func selectorRoot(expr ast.Expr) string {
	switch value := expr.(type) {
	case *ast.Ident:
		return value.Name
	case *ast.SelectorExpr:
		return selectorRoot(value.X)
	default:
		return ""
	}
}

func buildRegisterMethodDoc(src sourceFile, external externalIndexes, call *ast.CallExpr, bucket string, position sourcePosition) (functionDoc, bool, string) {
	if len(call.Args) < 3 {
		return functionDoc{}, false, "RegisterMethod 参数数量不足"
	}
	name, ok := stringValue(call.Args[0], src.constMap)
	if !ok || name == "" {
		return functionDoc{}, false, "RegisterMethod 方法名不是可解析字符串"
	}
	description, _ := stringValue(call.Args[1], src.constMap)
	sig := signatureForExpr(src, external, call.Args[2])
	return functionDoc{
		Name:        name,
		Description: nonEmpty(description, "AutoGo ScriptEngine 注册方法。"),
		Signature:   sig,
		Source:      position,
		Bucket:      bucket,
		Kind:        "function",
	}, true, ""
}

func buildRegisterValueDoc(src sourceFile, call *ast.CallExpr, bucket string, position sourcePosition) (functionDoc, bool, string) {
	if len(call.Args) < 1 {
		return functionDoc{}, false, "RegisterValue 参数数量不足"
	}
	name, ok := stringValue(call.Args[0], src.constMap)
	if !ok || name == "" {
		return functionDoc{}, false, "RegisterValue 字段名不是可解析字符串"
	}
	return functionDoc{
		Name:        name,
		Description: "AutoGo ScriptEngine 注册字段值。",
		Signature:   signature{Returns: []param{{Name: "value", Type: "any"}}},
		Source:      position,
		Bucket:      bucket,
		Kind:        "value",
	}, true, ""
}

func buildGLuaRegisterDoc(src sourceFile, external externalIndexes, call *ast.CallExpr, bucket string, position sourcePosition) (functionDoc, bool, string) {
	if len(call.Args) < 3 {
		return functionDoc{}, false, "glua.Register 参数数量不足"
	}
	name, ok := stringValue(call.Args[1], src.constMap)
	if !ok || name == "" {
		return functionDoc{}, false, "glua.Register 方法名不是可解析字符串"
	}
	return functionDoc{
		Name:        name,
		Description: "Lua 引擎核心全局函数。",
		Signature:   signatureForExpr(src, external, call.Args[2]),
		Source:      position,
		Bucket:      bucket,
		Kind:        "function",
	}, true, ""
}

func buildCoreTableDocs(src sourceFile, external externalIndexes, call *ast.CallExpr, bucket string, position sourcePosition) ([]functionDoc, bool, string) {
	if len(call.Args) < 2 {
		return nil, false, "installVMCoreTable 参数数量不足"
	}
	tableName, ok := stringValue(call.Args[0], src.constMap)
	if !ok || tableName == "" {
		return nil, false, "installVMCoreTable 表名不是可解析字符串"
	}
	lit, ok := call.Args[1].(*ast.CompositeLit)
	if !ok {
		return nil, false, "installVMCoreTable 方法表不是静态字面量"
	}
	var docs []functionDoc
	for _, elt := range lit.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}
		methodName, ok := stringValue(kv.Key, src.constMap)
		if !ok || methodName == "" {
			continue
		}
		fullName := tableName + "." + methodName
		docs = append(docs, functionDoc{
			Name:        fullName,
			Description: "Lua 引擎核心 table 方法。",
			Signature:   signatureForExpr(src, external, kv.Value),
			Source:      position,
			Bucket:      bucket,
			Kind:        "function",
		})
	}
	return docs, true, ""
}

func stringValue(expr ast.Expr, constMap map[string]string) (string, bool) {
	switch value := expr.(type) {
	case *ast.BasicLit:
		if value.Kind != token.STRING {
			return "", false
		}
		unquoted, err := strconv.Unquote(value.Value)
		return unquoted, err == nil
	case *ast.Ident:
		result, ok := constMap[value.Name]
		return result, ok
	case *ast.BinaryExpr:
		if value.Op != token.ADD {
			return "", false
		}
		left, ok := stringValue(value.X, constMap)
		if !ok {
			return "", false
		}
		right, ok := stringValue(value.Y, constMap)
		if !ok {
			return "", false
		}
		return left + right, true
	case *ast.ParenExpr:
		return stringValue(value.X, constMap)
	default:
		return "", false
	}
}

func signatureForExpr(src sourceFile, external externalIndexes, expr ast.Expr) signature {
	switch value := expr.(type) {
	case *ast.FuncLit:
		return signatureForFuncType(value.Type)
	case *ast.Ident:
		if fnType, ok := src.funcs.decls[value.Name]; ok {
			return signatureForFuncType(fnType)
		}
		return signature{Unknown: true, Params: []param{{Name: "...", Type: "any"}}, Returns: []param{{Name: "result", Type: "any"}}}
	case *ast.SelectorExpr:
		if sig, ok := selectorSignature(src, external, value); ok {
			return sig
		}
		if sig, ok := methodExpressionSignature(src, external, value); ok {
			return sig
		}
		return signature{Unknown: true, Params: []param{{Name: "...", Type: "any"}}, Returns: []param{{Name: "result", Type: "any"}}}
	case *ast.ParenExpr:
		return signatureForExpr(src, external, value.X)
	default:
		if sig, ok := methodExpressionSignature(src, external, value); ok {
			return sig
		}
		return signature{Unknown: true, Params: []param{{Name: "...", Type: "any"}}, Returns: []param{{Name: "result", Type: "any"}}}
	}
}

func selectorSignature(src sourceFile, external externalIndexes, expr *ast.SelectorExpr) (signature, bool) {
	root, ok := expr.X.(*ast.Ident)
	if !ok {
		return signature{}, false
	}
	importPath, ok := src.imports[root.Name]
	if !ok {
		return signature{}, false
	}
	index := external.indexForImport(importPath)
	fnType, ok := index.decls[expr.Sel.Name]
	if !ok {
		return signature{}, false
	}
	return signatureForFuncType(fnType), true
}

func methodExpressionSignature(src sourceFile, external externalIndexes, expr ast.Expr) (signature, bool) {
	selector, ok := expr.(*ast.SelectorExpr)
	if !ok {
		return signature{}, false
	}
	recvImport, recvType := selectorReceiver(selector.X)
	if recvImport == "" || recvType == "" {
		return signature{}, false
	}
	importPath, ok := src.imports[recvImport]
	if !ok {
		return signature{}, false
	}
	index := external.indexForImport(importPath)
	fnType, ok := index.methods[recvType+"."+selector.Sel.Name]
	if !ok {
		return signature{}, false
	}
	sig := signatureForFuncType(fnType)
	sig.Params = append([]param{{Name: "self", Type: "userdata"}}, sig.Params...)
	return sig, true
}

func selectorReceiver(expr ast.Expr) (string, string) {
	switch value := expr.(type) {
	case *ast.ParenExpr:
		return selectorReceiver(value.X)
	case *ast.StarExpr:
		return selectorReceiver(value.X)
	case *ast.SelectorExpr:
		if root, ok := value.X.(*ast.Ident); ok {
			return root.Name, value.Sel.Name
		}
	}
	return "", ""
}

func (indexes externalIndexes) indexForImport(importPath string) functionIndex {
	if index, ok := indexes[importPath]; ok {
		return index
	}
	index := functionIndex{decls: map[string]*ast.FuncType{}, methods: map[string]*ast.FuncType{}}
	dir, err := goListDir(importPath)
	if err != nil {
		indexes[importPath] = index
		return index
	}
	entries, err := os.ReadDir(dir)
	if err != nil {
		indexes[importPath] = index
		return index
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".go") || strings.HasSuffix(entry.Name(), "_test.go") {
			continue
		}
		path := filepath.Join(dir, entry.Name())
		parsed, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if err != nil {
			continue
		}
		fileIndex := collectFuncDecls(parsed)
		for name, fnType := range fileIndex.decls {
			index.decls[name] = fnType
		}
		for name, fnType := range fileIndex.methods {
			index.methods[name] = fnType
		}
	}
	indexes[importPath] = index
	return index
}

func goListDir(importPath string) (string, error) {
	cmd := exec.Command("go", "list", "-f", "{{.Dir}}", importPath)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(output)), nil
}

func signatureForFuncType(fn *ast.FuncType) signature {
	if fn == nil {
		return signature{Unknown: true}
	}
	return signature{
		Params:  fieldsToParams(fn.Params, true),
		Returns: fieldsToParams(fn.Results, false),
	}
}

func fieldsToParams(list *ast.FieldList, input bool) []param {
	if list == nil {
		return nil
	}
	var params []param
	argIndex := 1
	resultIndex := 1
	for _, field := range list.List {
		luaType := luaType(field.Type)
		names := field.Names
		if len(names) == 0 {
			if input {
				params = append(params, param{Name: fmt.Sprintf("arg%d", argIndex), Type: luaType})
				argIndex++
			} else if luaType != "error" {
				params = append(params, param{Name: fmt.Sprintf("result%d", resultIndex), Type: luaType})
				resultIndex++
			}
			continue
		}
		for _, name := range names {
			paramName := name.Name
			if !input && luaType == "error" {
				continue
			}
			if input && strings.HasPrefix(luaType, "...") {
				paramName = "..."
			}
			params = append(params, param{Name: paramName, Type: luaType})
			if input {
				argIndex++
			} else {
				resultIndex++
			}
		}
	}
	return params
}

func luaType(expr ast.Expr) string {
	switch value := expr.(type) {
	case *ast.Ident:
		return identLuaType(value.Name)
	case *ast.SelectorExpr:
		return "userdata"
	case *ast.ArrayType:
		if ident, ok := value.Elt.(*ast.Ident); ok && ident.Name == "byte" {
			return "string"
		}
		return "table"
	case *ast.MapType:
		return "table"
	case *ast.StructType:
		return "table"
	case *ast.InterfaceType:
		return "any"
	case *ast.FuncType:
		return "function"
	case *ast.StarExpr:
		switch value.X.(type) {
		case *ast.StructType, *ast.SelectorExpr, *ast.Ident:
			return "userdata"
		default:
			return "userdata"
		}
	case *ast.Ellipsis:
		return "..." + luaType(value.Elt)
	default:
		return "any"
	}
}

func identLuaType(name string) string {
	switch name {
	case "string":
		return "string"
	case "bool":
		return "boolean"
	case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "uintptr", "float32", "float64", "complex64", "complex128":
		return "number"
	case "byte":
		return "string"
	case "error":
		return "error"
	case "any", "interface":
		return "any"
	default:
		return "userdata"
	}
}

func dedupeDocs(bucket string, docs []functionDoc) ([]functionDoc, int, []conflict) {
	sort.SliceStable(docs, func(i, j int) bool {
		if docs[i].Name == docs[j].Name {
			if docs[i].Source.File == docs[j].Source.File {
				return docs[i].Source.Line < docs[j].Source.Line
			}
			return docs[i].Source.File < docs[j].Source.File
		}
		return docs[i].Name < docs[j].Name
	})
	byName := map[string]functionDoc{}
	var duplicates int
	conflictMap := map[string]*conflict{}
	for _, doc := range docs {
		if previous, exists := byName[doc.Name]; exists {
			duplicates++
			if signatureText(previous) != signatureText(doc) {
				item := conflictMap[doc.Name]
				if item == nil {
					item = &conflict{Name: doc.Name, Bucket: bucket, Kept: sourceText(doc)}
					conflictMap[doc.Name] = item
				}
				item.Discarded = append(item.Discarded, sourceText(previous))
			}
		}
		byName[doc.Name] = doc
	}
	var names []string
	for name := range byName {
		names = append(names, name)
	}
	sort.Strings(names)
	result := make([]functionDoc, 0, len(names))
	for _, name := range names {
		result = append(result, byName[name])
	}
	var conflicts []conflict
	for _, item := range conflictMap {
		sort.Strings(item.Discarded)
		conflicts = append(conflicts, *item)
	}
	sort.Slice(conflicts, func(i, j int) bool {
		return conflicts[i].Name < conflicts[j].Name
	})
	return result, duplicates, conflicts
}

func writeCatalogs(outDir string, docs map[string][]functionDoc) ([]generatedFile, error) {
	names := []string{"common", "android", "ios", "lrappsoft"}
	var generated []generatedFile
	for _, name := range names {
		catalog := builtinCatalog{Schema: schema, Functions: orderedBuiltinDocs{}}
		for _, doc := range docs[name] {
			catalog.Functions[doc.Name] = doc.toBuiltinDoc()
		}
		fileName := "autogo-scriptengine-" + name + ".json"
		path := filepath.Join(outDir, fileName)
		if err := writeJSON(path, catalog); err != nil {
			return nil, err
		}
		generated = append(generated, generatedFile{Name: name, Path: filepath.ToSlash(filepath.Join("docs/glua_builtin_docs", fileName)), Count: len(catalog.Functions)})
	}
	return generated, nil
}

func writeReport(path string, rep report) error {
	return writeJSON(path, rep)
}

func writeJSON(path string, value interface{}) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0o644)
}

func writeReadme(path string, generated []generatedFile) error {
	var builder strings.Builder
	builder.WriteString("# GLua Builtin Docs\n\n")
	builder.WriteString("本目录由 `go run ./tools/glua_builtin_docs` 生成，供 `go-lua-vm`/`gluals`、VS Code 和 JetBrains 导入。\n\n")
	builder.WriteString("## 文件\n\n")
	for _, item := range generated {
		builder.WriteString(fmt.Sprintf("- `%s`：%d 个条目\n", filepath.Base(item.Path), item.Count))
	}
	builder.WriteString("\n")
	builder.WriteString("## 重新生成\n\n")
	builder.WriteString("```bash\n")
	builder.WriteString("go run ./tools/glua_builtin_docs\n")
	builder.WriteString("```\n\n")
	builder.WriteString("## VS Code\n\n")
	builder.WriteString("```json\n")
	builder.WriteString("{\n")
	builder.WriteString("  \"glua.builtinDocs\": [\n")
	builder.WriteString("    \"docs/glua_builtin_docs/autogo-scriptengine-common.json\",\n")
	builder.WriteString("    \"docs/glua_builtin_docs/autogo-scriptengine-android.json\",\n")
	builder.WriteString("    \"docs/glua_builtin_docs/autogo-scriptengine-lrappsoft.json\"\n")
	builder.WriteString("  ]\n")
	builder.WriteString("}\n")
	builder.WriteString("```\n\n")
	builder.WriteString("iOS 项目将 `android` 文件替换为 `autogo-scriptengine-ios.json`。修改 JSON 后需要重新加载 IDE 窗口或重启语言服务。\n")
	return os.WriteFile(path, []byte(builder.String()), 0o644)
}

func (doc functionDoc) toBuiltinDoc() builtinDoc {
	sig := doc.signatureString()
	description := doc.Description
	source := fmt.Sprintf("%s:%d", doc.Source.File, doc.Source.Line)
	return builtinDoc{
		Signature: localizedText{
			"en":    sig,
			"zh-CN": sig,
		},
		Returns: localizedText{
			"en":    returnsText(doc.Signature.Returns, false),
			"zh-CN": returnsText(doc.Signature.Returns, true),
		},
		Params: localizedList{
			"en":    paramText(doc.Signature.Params, false),
			"zh-CN": paramText(doc.Signature.Params, true),
		},
		Description: localizedText{
			"en":    description + " Source: " + source,
			"zh-CN": description + " 来源：" + source,
		},
		Example: localizedText{
			"en":    exampleText(doc.Name),
			"zh-CN": exampleText(doc.Name),
		},
	}
}

func (doc functionDoc) signatureString() string {
	params := make([]string, 0, len(doc.Signature.Params))
	for _, item := range doc.Signature.Params {
		if item.Name == "..." || strings.HasPrefix(item.Type, "...") {
			params = append(params, "...")
			continue
		}
		params = append(params, item.Name)
	}
	if doc.Kind == "value" {
		return doc.Name
	}
	return fmt.Sprintf("%s(%s)", doc.Name, strings.Join(params, ", "))
}

func paramText(params []param, zh bool) []string {
	if len(params) == 0 {
		if zh {
			return []string{"无参数。"}
		}
		return []string{"No parameters."}
	}
	result := make([]string, 0, len(params))
	for _, item := range params {
		name := item.Name
		if name == "..." || strings.HasPrefix(item.Type, "...") {
			name = "..."
		}
		if zh {
			result = append(result, fmt.Sprintf("%s：%s。", name, item.Type))
		} else {
			result = append(result, fmt.Sprintf("%s: %s.", name, item.Type))
		}
	}
	return result
}

func returnsText(params []param, zh bool) string {
	if len(params) == 0 {
		if zh {
			return "返回：无。"
		}
		return "returns: none."
	}
	parts := make([]string, 0, len(params))
	for _, item := range params {
		parts = append(parts, item.Type)
	}
	if zh {
		return "返回：" + strings.Join(parts, ", ") + "。"
	}
	return "returns: " + strings.Join(parts, ", ") + "."
}

func exampleText(name string) string {
	if strings.Contains(name, ".") {
		return "-- " + name + "(...)"
	}
	return name + "(...)"
}

func signatureText(doc functionDoc) string {
	return doc.signatureString() + " -> " + returnsText(doc.Signature.Returns, false)
}

func sourceText(doc functionDoc) string {
	return fmt.Sprintf("%s:%d %s", doc.Source.File, doc.Source.Line, signatureText(doc))
}

func nonEmpty(value string, fallback string) string {
	if strings.TrimSpace(value) == "" {
		return fallback
	}
	return value
}
