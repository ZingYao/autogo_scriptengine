package debugger

import (
	"fmt"
	"sort"

	lua "github.com/yuin/gopher-lua"
)

func (d *Debugger) captureFrame(L *lua.LState, pos Position) *Frame {
	frame := &Frame{
		ID:       1,
		Name:     "lua",
		Position: pos,
	}
	if d.config.CollectLocals {
		frame.Locals = d.captureLocals(L)
	}
	if d.config.CollectGlobals {
		frame.Globals = d.captureGlobals(L)
	}
	return frame
}

func (d *Debugger) captureLocals(L *lua.LState) []Variable {
	dbg, ok := L.GetStack(1)
	if !ok {
		return nil
	}
	var variables []Variable
	for idx := 1; idx <= 256; idx++ {
		name, value := L.GetLocal(dbg, idx)
		if name == "" {
			break
		}
		variables = append(variables, Variable{
			Name:  name,
			Type:  value.Type().String(),
			Value: d.luaValueToSnapshot(value, 0),
		})
	}
	return variables
}

func (d *Debugger) captureGlobals(L *lua.LState) []Variable {
	var variables []Variable
	global := L.G.Global
	global.ForEach(func(key lua.LValue, value lua.LValue) {
		name := key.String()
		if name == hitFunction {
			return
		}
		variables = append(variables, Variable{
			Name:  name,
			Type:  value.Type().String(),
			Value: d.luaValueToSnapshot(value, 0),
		})
	})
	sort.Slice(variables, func(i, j int) bool {
		return variables[i].Name < variables[j].Name
	})
	return variables
}

func (d *Debugger) luaValueToSnapshot(value lua.LValue, depth int) interface{} {
	if value == nil {
		return nil
	}
	if depth >= d.config.MaxVariableDepth {
		return value.String()
	}
	switch v := value.(type) {
	case lua.LBool:
		return bool(v)
	case lua.LNumber:
		return float64(v)
	case lua.LString:
		return string(v)
	case *lua.LTable:
		result := map[string]interface{}{}
		count := 0
		v.ForEach(func(key lua.LValue, item lua.LValue) {
			if count >= 128 {
				return
			}
			result[key.String()] = d.luaValueToSnapshot(item, depth+1)
			count++
		})
		return result
	case *lua.LFunction:
		return "function"
	case *lua.LUserData:
		return fmt.Sprintf("userdata:%T", v.Value)
	case *lua.LState:
		return "thread"
	default:
		if value == lua.LNil {
			return nil
		}
		return value.String()
	}
}
