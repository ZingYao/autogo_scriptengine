#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
README.md 文档生成脚本

这个脚本用于根据 *_inject.go 源文件自动生成 README.md 文档。
它会解析 Go 源码中的方法注册信息，生成包含模块简介、常量、方法列表和使用示例的中文文档。

使用方法:
    python generate_readme.py

作者: AutoGo Engine
日期: 2026-03-25
"""

import os
import re
import glob
from typing import Dict, List, Tuple, Optional


class MethodInfo:
    """
    方法信息类
    
    用于存储一个方法的所有相关信息，包括名称、描述、参数和返回值。
    """
    def __init__(self, name: str, description: str, params: List[Tuple[str, str]], 
                 returns: str = "", example: str = ""):
        # 方法名称
        self.name = name
        # 方法中文描述
        self.description = description
        # 参数列表，每个元素是 (参数名, 参数类型) 的元组
        self.params = params
        # 返回值描述
        self.returns = returns
        # 使用示例代码
        self.example = example


class ModuleParser:
    """
    模块解析器类
    
    负责解析 Go 源码文件，提取模块名称、方法和文档信息。
    """
    
    def __init__(self, file_path: str):
        # 源文件路径
        self.file_path = file_path
        # 模块名称（从文件路径提取）
        self.module_name = os.path.basename(os.path.dirname(file_path))
        # 引擎类型（js_engine 或 lua_engine）
        self.engine_type = "js" if "js_engine" in file_path else "lua"
        # 方法列表
        self.methods: List[MethodInfo] = []
        # 模块描述
        self.module_desc = ""
        # 常量列表
        self.constants: List[Tuple[str, str, str]] = []
        
    def parse(self) -> bool:
        """
        解析源文件
        
        读取并解析 Go 源码文件，提取所有方法信息。
        
        返回:
            bool: 解析是否成功
        """
        try:
            with open(self.file_path, 'r', encoding='utf-8') as f:
                content = f.read()
            
            # 提取模块描述（从注释中获取）
            self._extract_module_desc(content)
            
            # 提取常量定义
            self._extract_constants(content)
            
            # 提取方法定义
            self._extract_methods(content)
            
            return True
        except Exception as e:
            print(f"解析文件失败 {self.file_path}: {e}")
            return False
    
    def _extract_module_desc(self, content: str):
        """
        提取模块描述
        
        从文件顶部的注释中提取模块描述信息。
        
        参数:
            content: 文件内容字符串
        """
        # 查找模块结构体的注释
        module_pattern = r'//\s*(\w+)Module\s+(\S+)\s*模块'
        match = re.search(module_pattern, content)
        if match:
            self.module_desc = match.group(2) + "模块"
        else:
            self.module_desc = self.module_name + " 模块"
    
    def _extract_constants(self, content: str):
        """
        提取常量定义
        
        从源码中提取 const 定义的常量。
        
        参数:
            content: 文件内容字符串
        """
        # 查找常量定义块
        const_pattern = r'const\s*\((.*?)\)'
        const_match = re.search(const_pattern, content, re.DOTALL)
        if const_match:
            const_block = const_match.group(1)
            # 提取每个常量
            const_lines = const_block.strip().split('\n')
            for line in const_lines:
                line = line.strip()
                if line and not line.startswith('//'):
                    # 匹配 常量名 = iota // 注释 格式
                    match = re.match(r'(\w+)\s*=\s*iota\s*//\s*(.+)', line)
                    if match:
                        self.constants.append((match.group(1), "iota", match.group(2)))
    
    def _extract_methods(self, content: str):
        """
        提取方法定义
        
        从 Register 方法中提取所有注册的方法信息。
        
        参数:
            content: 文件内容字符串
        """
        # 查找 engine.RegisterMethod 调用
        # 匹配 engine.RegisterMethod("方法名", "描述", ...)
        pattern = r'engine\.RegisterMethod\("([^"]+)",\s*"([^"]+)"'
        matches = re.findall(pattern, content)
        
        for method_name, description in matches:
            # 解析方法名（格式：模块名.方法名）
            parts = method_name.split('.')
            if len(parts) >= 2:
                func_name = parts[-1]
            else:
                func_name = method_name
            
            # 创建方法信息对象
            method = MethodInfo(
                name=func_name,
                description=description,
                params=[],  # 参数信息需要从函数定义中提取
                returns=""
            )
            self.methods.append(method)
        
        # 如果没有找到 RegisterMethod，尝试从 Set 调用中提取
        if not self.methods:
            self._extract_set_methods(content)
    
    def _extract_set_methods(self, content: str):
        """
        从 Set 调用中提取方法
        
        对于直接设置到对象的方法（如 xxxObj.Set("methodName", ...)）
        
        参数:
            content: 文件内容字符串
        """
        # 匹配 xxxObj.Set("方法名", func(...)
        pattern = r'\w+Obj\.Set\("([^"]+)",\s*func\(([^)]*)\)'
        matches = re.findall(pattern, content)
        
        for method_name, params in matches:
            # 尝试从上下文找到描述
            description = self._find_method_description(content, method_name)
            
            method = MethodInfo(
                name=method_name,
                description=description or f"{method_name} 方法",
                params=[],
                returns=""
            )
            self.methods.append(method)
    
    def _find_method_description(self, content: str, method_name: str) -> str:
        """
        查找方法描述
        
        在源码中查找方法的注释描述。
        
        参数:
            content: 文件内容字符串
            method_name: 方法名
            
        返回:
            str: 方法描述，如果没有找到则返回空字符串
        """
        # 查找方法前的注释
        pattern = rf'//\s*{method_name}\s*[-–]\s*(.+?)(?:\n|$)'
        match = re.search(pattern, content)
        if match:
            return match.group(1).strip()
        return ""


class ReadmeGenerator:
    """
    README.md 文档生成器类
    
    负责根据解析的模块信息生成格式化的 Markdown 文档。
    """
    
    def __init__(self, parser: ModuleParser):
        # 模块解析器
        self.parser = parser
    
    def generate(self) -> str:
        """
        生成 README.md 内容
        
        根据解析的模块信息生成完整的 Markdown 文档。
        
        返回:
            str: Markdown 格式的文档内容
        """
        lines = []
        
        # 添加标题
        lines.append(f"# {self.parser.module_name} 模块")
        lines.append("")
        
        # 添加模块简介
        lines.append("## 模块简介")
        lines.append("")
        lines.append(self._get_module_intro())
        lines.append("")
        
        # 添加常量部分（如果有）
        if self.parser.constants:
            lines.append("## 常量")
            lines.append("")
            lines.append(self._generate_constants_table())
            lines.append("")
        
        # 添加方法列表
        lines.append("## 方法列表")
        lines.append("")
        
        for method in self.parser.methods:
            lines.append(self._generate_method_doc(method))
            lines.append("")
        
        # 添加综合使用示例
        lines.append("## 综合使用示例")
        lines.append("")
        lines.append(self._generate_examples())
        
        return '\n'.join(lines)
    
    def _get_module_intro(self) -> str:
        """
        获取模块简介
        
        根据模块名称返回对应的中文描述。
        
        返回:
            str: 模块简介文本
        """
        # 模块描述映射表
        module_intros = {
            'coroutine': 'coroutine 模块提供了 JavaScript 协程（并发）功能的支持。该模块允许在脚本中创建和管理协程，实现异步任务的并发执行、协程池管理以及任务调度等功能。',
            'dotocr': 'dotocr 模块提供了 OCR（光学字符识别）功能，支持从屏幕、图像文件、Base64 编码的图像等多种来源进行文字识别。该模块适用于自动化脚本中需要识别屏幕文字的场景。',
            'hud': 'hud 模块提供了在屏幕上显示浮动窗口（HUD）的功能。可以创建自定义的浮动界面，用于显示调试信息、状态提示等内容。',
            'images': 'images 模块提供了图像处理功能，包括屏幕截图、颜色识别、图像变换（裁剪、缩放、旋转等）以及图像格式转换等功能。',
            'ime': 'ime 模块提供了输入法相关的功能，包括剪贴板操作、文本输入、输入法切换等。',
            'imgui': 'imgui 模块提供了 ImGui 图形用户界面库的绑定，用于创建复杂的交互式 GUI 界面。',
            'media': 'media 模块提供了媒体相关的功能，包括音频播放、文件扫描、短信发送等。',
            'motion': 'motion 模块提供了模拟用户操作的功能，包括点击、滑动、按键等手势操作。',
            'opencv': 'opencv 模块提供了 OpenCV 计算机视觉库的功能，支持图像匹配、特征检测等高级图像处理功能。',
            'plugin': 'plugin 模块提供了加载外部 APK 插件的功能，允许动态加载和执行外部代码。',
            'ppocr': 'ppocr 模块提供了基于 PaddleOCR 的文字识别功能，支持高精度的中文和英文识别。',
            'rhino': 'rhino 模块提供了 Rhino JavaScript 引擎的集成功能，用于执行 JavaScript 代码。',
            'storages': 'storages 模块提供了本地存储功能，支持键值对的持久化存储。',
            'system': 'system 模块提供了系统级别的功能，包括进程管理、内存监控、系统设置等。',
            'uiacc': 'uiacc 模块提供了无障碍服务（Accessibility）的功能，用于查找和操作界面元素。',
            'utils': 'utils 模块提供了各种实用工具函数，包括日志记录、数据转换、加密解密、编码解码等。',
            'vdisplay': 'vdisplay 模块提供了虚拟显示设备的功能，可以创建和管理虚拟屏幕。',
            'yolo': 'yolo 模块提供了 YOLO（You Only Look Once）目标检测功能，用于实时对象检测。',
            'console': 'console 模块提供了控制台窗口的功能，用于显示日志和调试信息。',
            'app': 'app 模块提供了应用管理功能，包括应用启动、安装、卸载、信息获取等。',
            'files': 'files 模块提供了文件系统操作功能，包括文件读写、复制、移动、删除等。',
            'http': 'http 模块提供了 HTTP 网络请求功能，支持 GET、POST 等请求方式。',
            'device': 'device 模块提供了设备信息获取和设置功能，包括屏幕、音量、电池等信息。',
        }
        
        return module_intros.get(self.parser.module_name, 
                f'{self.parser.module_name} 模块提供了相关的功能支持。')
    
    def _generate_constants_table(self) -> str:
        """
        生成常量表格
        
        将常量列表格式化为 Markdown 表格。
        
        返回:
            str: Markdown 表格格式的常量列表
        """
        lines = ["| 常量名 | 值 | 说明 |", "|--------|-----|------|"]
        for name, value, desc in self.parser.constants:
            lines.append(f"| {name} | {value} | {desc} |")
        return '\n'.join(lines)
    
    def _generate_method_doc(self, method: MethodInfo) -> str:
        """
        生成单个方法的文档
        
        将方法信息格式化为 Markdown 文档格式。
        
        参数:
            method: 方法信息对象
            
        返回:
            str: Markdown 格式的方法文档
        """
        lines = []
        
        # 方法标题
        full_name = f"{self.parser.module_name}.{method.name}"
        lines.append(f"### {full_name}")
        lines.append(method.description)
        lines.append("")
        
        # 参数部分
        if method.params:
            lines.append("**参数：**")
            for param_name, param_type in method.params:
                lines.append(f"- `{param_name}` ({param_type}): 参数描述")
            lines.append("")
        
        # 返回值
        if method.returns:
            lines.append(f"**返回值：** {method.returns}")
            lines.append("")
        
        # 使用示例
        lines.append("**使用示例：**")
        lines.append("```javascript")
        if method.example:
            lines.append(method.example)
        else:
            lines.append(f"// 调用 {full_name} 方法")
            if method.params:
                param_examples = []
                for param_name, param_type in method.params:
                    if "string" in param_type.lower():
                        param_examples.append(f'"{param_name}"')
                    elif "int" in param_type.lower() or "number" in param_type.lower():
                        param_examples.append("0")
                    elif "bool" in param_type.lower():
                        param_examples.append("true")
                    else:
                        param_examples.append("null")
                lines.append(f"var result = {full_name}({', '.join(param_examples)});")
            else:
                lines.append(f"{full_name}();")
        lines.append("```")
        lines.append("")
        lines.append("---")
        
        return '\n'.join(lines)
    
    def _generate_examples(self) -> str:
        """
        生成综合使用示例
        
        根据模块类型生成相应的使用示例代码。
        
        返回:
            str: Markdown 格式的示例代码
        """
        # 示例代码映射表
        examples = {
            'coroutine': '''### 示例1：启动协程
```javascript
var coroId = coroutine.launch(function() {
    console.log("协程开始执行");
    coroutine.sleep(1000);
    console.log("协程执行完成");
}, "myCoroutine", 1);
```''',
            'dotocr': '''### 示例1：识别屏幕文字
```javascript
var text = dotocr.ocr(0, 0, 500, 200, "FFFFFF-000000", 5, 5, 0.9, 0, "default", 0);
console.log("识别结果: " + text);
```''',
            'images': '''### 示例1：截图并保存
```javascript
var img = images.captureScreen(0, 0, device.width, device.height);
images.save(img, "/sdcard/screenshot.png", 100);
```''',
            'motion': '''### 示例1：点击屏幕
```javascript
click(500, 500);  // 在坐标(500, 500)处点击
swipe(100, 500, 900, 500, 500);  // 从(100,500)滑动到(900,500)，耗时500ms
```''',
            'utils': '''### 示例1：日志和提示
```javascript
utils.logI("TAG", "这是一条日志");
utils.toast("操作成功");
```''',
            'files': '''### 示例1：读写文件
```javascript
files.write("/sdcard/test.txt", "Hello World");
var content = files.read("/sdcard/test.txt");
console.log(content);
```''',
            'http': '''### 示例1：发送HTTP请求
```javascript
var response = http.get("https://api.example.com/data", 5000);
console.log("状态码: " + response.code);
console.log("数据: " + response.data);
```''',
            'device': '''### 示例1：获取设备信息
```javascript
console.log("屏幕宽度: " + device.width);
console.log("屏幕高度: " + device.height);
console.log("电量: " + device.getBattery() + "%");
```''',
            'app': '''### 示例1：启动应用
```javascript
app.launch("com.example.app", 0);
var currentPkg = app.currentPackage();
console.log("当前应用: " + currentPkg);
```''',
            'storages': '''### 示例1：存储数据
```javascript
storages.put("config", "username", "admin");
var username = storages.get("config", "username");
console.log("用户名: " + username);
```''',
        }
        
        return examples.get(self.parser.module_name, 
                f"### 示例1：基本使用\n```javascript\n// {self.parser.module_name} 模块的基本使用示例\n```")


def find_inject_files(base_path: str) -> List[str]:
    """
    查找所有的 *_inject.go 文件
    
    参数:
        base_path: 基础搜索路径
        
    返回:
        List[str]: 文件路径列表
    """
    pattern = os.path.join(base_path, "**", "*_inject.go")
    return glob.glob(pattern, recursive=True)


def main():
    """
    主函数
    
    程序的入口点，负责协调整个文档生成流程。
    """
    # 获取脚本所在目录
    script_dir = os.path.dirname(os.path.abspath(__file__))
    
    print("=" * 60)
    print("README.md 文档生成工具")
    print("=" * 60)
    print()
    
    # 查找所有 inject 文件
    inject_files = find_inject_files(script_dir)
    print(f"找到 {len(inject_files)} 个模块文件")
    print()
    
    # 统计信息
    success_count = 0
    failed_count = 0
    skipped_count = 0
    
    # 处理每个文件
    for file_path in inject_files:
        # 获取模块目录
        module_dir = os.path.dirname(file_path)
        readme_path = os.path.join(module_dir, "README.md")
        
        # 检查是否已存在 README.md
        if os.path.exists(readme_path):
            print(f"跳过（已存在）: {file_path}")
            skipped_count += 1
            continue
        
        print(f"处理: {file_path}")
        
        # 解析文件
        parser = ModuleParser(file_path)
        if parser.parse():
            # 生成文档
            generator = ReadmeGenerator(parser)
            content = generator.generate()
            
            # 写入文件
            try:
                with open(readme_path, 'w', encoding='utf-8') as f:
                    f.write(content)
                print(f"  已生成: {readme_path}")
                success_count += 1
            except Exception as e:
                print(f"  写入失败: {e}")
                failed_count += 1
        else:
            print(f"  解析失败")
            failed_count += 1
    
    print()
    print("=" * 60)
    print("生成完成!")
    print(f"  成功: {success_count}")
    print(f"  失败: {failed_count}")
    print(f"  跳过: {skipped_count}")
    print("=" * 60)


if __name__ == "__main__":
    main()
