#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
Markdown 转 docsify 文档系统
将项目中的所有 README.md 文件转换为 docsify 格式的文档
"""

import os
import re
from pathlib import Path
import shutil
import sys

# 检查是否安装了 markdown 库
try:
    import markdown
except ImportError:
    print("错误: 未安装 markdown 库")
    print("请运行: pip install markdown")
    sys.exit(1)

# docsify 配置文件模板
INDEX_HTML_TEMPLATE = """<!DOCTYPE html>

<html lang="zh-CN">
<head>
<meta charset="utf-8"/>
<title>AutoGo ScriptEngine</title>
<link href="icon.svg" rel="icon" type="image/x-icon"/>
<meta content="never" name="referrer"/>
<meta content="IE=edge,chrome=1" http-equiv="X-UA-Compatible">
<meta content="AutoGo ScriptEngine - 为 AutoGo 提供 JavaScript 和 Lua 脚本引擎支持" name="description"/>
<meta content="width=device-width, user-scalable=no, initial-scale=1.0, maximum-scale=1.0, minimum-scale=1.0" name="viewport"/>
<link href="style/vue.css" rel="stylesheet"/>
<link href="style/myStyle.css" rel="stylesheet" type="text/css">
<style>
    .cover {
      background: linear-gradient(to left bottom, hsl(216, 100%, 85%) 0%, hsl(107, 100%, 85%) 100%) !important;
    }
  </style>
</head>
<body>
<div id="app">Loading...</div>
<a class="to-top">Top</a>
<script>
    window.$docsify = {
      el: '#app',
      name: 'AutoGo ScriptEngine',
      repo: 'https://github.com/ZingYao/autogo_scriptengine',
      coverpage: true,
      loadSidebar: true,
      auto2top: true,
      subMaxLevel: 2,
      maxLevel: 4,
      homepage: 'README.md',

      alias: {
        '/API/_sidebar.md': '/_sidebar.md'
      },
      
      search: {
        paths: 'auto',
        placeholder: {
          '/':'🔍 搜索'
        },
        noData: {
          '/':'😒 找不到结果',
        },
        depth: 4,
        maxAge: 86400000,
      },

      footer: {
        copy: '<span>MIT License</span>',
        auth: ' <strong>AutoGo ScriptEngine</strong>',
        pre: '<hr/>',
        style: 'text-align: center;',
      },

      copyCode: {
          buttonText: '复制',
          errorText: '错误',
          successText: '成功!'
      },

      'flexible-alerts': {
        style: 'flat'
      }

    }
  </script>
<script src="docsify.min.js"></script>
<script src="prism-go.min.js"></script>
<script src="docsify-copy-code"></script>
<script src="search.min.js"></script>
<script src="docsify-footer.min.js"></script>
<script src="docsify-plugin-flexible-alerts.min.js"></script>
<script src="zoom-image.js"></script>
<script src="style/jquery-1.11.3.min.js"></script>
<script src="style/jquery.toTop.min.js"></script>
<script>
    $('.to-top').toTop();
  </script>
</body>
</html>
"""

# _sidebar.md 模板
SIDEBAR_TEMPLATE = """- 前言
  - [简介](README.md)
  - [更新日志](changelog.md)

- 引擎文档
  - [JavaScript 引擎](js_engine/README.md)
  - [Lua 引擎](lua_engine/README.md)

- API 文档
{api_docs}
"""

# README.md 模板
README_TEMPLATE = """---
type: cover
---

# AutoGo ScriptEngine

## 为 AutoGo 提供 JavaScript 和 Lua 脚本引擎支持

> 让开发者可以用熟悉的脚本语言编写自动化任务

![color](#f0f4f8)

<div style="display: flex; justify-content: center; margin: 40px 0;">
  <img src="https://trae-api-cn.mchost.guru/api/ide/v1/text_to_image?prompt=cute%20cartoon%20robot%20programmer%20with%20code%20symbols%20around%20it%2C%20friendly%20expression%2C%20simple%20style%2C%20blue%20and%20green%20colors&image_size=square" alt="ScriptEngine Mascot" style="width: 200px; height: 200px; border-radius: 50%; box-shadow: 0 4px 12px rgba(0,0,0,0.1);">
</div>

## 核心特点

| 特点 | 描述 |
|------|------|
| 🚀 **双引擎支持** | 同时支持 JavaScript 和 Lua 脚本语言 |
| 📚 **丰富的 API** | 提供应用管理、设备控制、图像识别、OCR 等多种功能 |
| 🔧 **方法注册系统** | 支持动态注册、重写和恢复方法 |
| 🔄 **协程支持** | Lua 引擎支持协程操作，提升并发能力 |
| 📖 **文档生成** | 可自动生成 API 文档，方便查阅 |
| 🔒 **代码保护** | 脚本代码易于混淆加密，有效保护业务逻辑 |
| 🔥 **热更新支持** | 脚本可动态加载，无需重新编译即可更新功能 |
| 🔄 **无痛迁移** | 可以无痛迁移其他平台的代码，复用现有的脚本代码库 |

## 快速开始

<div style="display: flex; justify-content: center; gap: 20px; margin: 40px 0;">
  <a href="#/js_engine/README.md" style="padding: 12px 30px; background: linear-gradient(135deg, #667eea 0%, #764ba2 100%); color: white; text-decoration: none; border-radius: 8px; font-weight: 600; box-shadow: 0 4px 12px rgba(102, 126, 234, 0.3); transition: all 0.3s ease;">JavaScript 引擎</a>
  <a href="#/lua_engine/README.md" style="padding: 12px 30px; background: linear-gradient(135deg, #f093fb 0%, #f5576c 100%); color: white; text-decoration: none; border-radius: 8px; font-weight: 600; box-shadow: 0 4px 12px rgba(240, 147, 251, 0.3); transition: all 0.3s ease;">Lua 引擎</a>
  <a href="#/changelog.md" style="padding: 12px 30px; background: linear-gradient(135deg, #4facfe 0%, #00f2fe 100%); color: white; text-decoration: none; border-radius: 8px; font-weight: 600; box-shadow: 0 4px 12px rgba(79, 172, 254, 0.3); transition: all 0.3s ease;">更新日志</a>
</div>

## 安装

```bash
go get github.com/ZingYao/autogo_scriptengine@v0.0.9
```

## 与 AutoGo 的关系

本项目是 AutoGo 的扩展方案，通过封装 AutoGo 提供的原生 API，为开发者提供更灵活的脚本编写方式：

- **AutoGo** - 提供 Android 自动化的核心能力（无障碍服务、图像识别、触摸模拟等）
- **ScriptEngine** - 为 AutoGo 添加脚本语言支持，让开发者可以用 JavaScript 或 Lua 编写自动化脚本

## 许可证

MIT License

---

<style>
  :root {
    --theme-color: #667eea;
  }
  
  .cover {
    background: linear-gradient(135deg, #f5f7fa 0%, #c3cfe2 100%) !important;
  }
  
  h1 {
    color: #2c3e50 !important;
    font-size: 2.5rem !important;
    margin-bottom: 1rem !important;
  }
  
  h2 {
    color: #34495e !important;
    font-size: 1.5rem !important;
    margin-bottom: 2rem !important;
  }
  
  blockquote {
    border-left: 4px solid var(--theme-color) !important;
    background: rgba(102, 126, 234, 0.1) !important;
    padding: 15px 20px !important;
    border-radius: 0 8px 8px 0 !important;
    margin: 20px 0 !important;
  }
  
  table {
    width: 100% !important;
    border-collapse: collapse !important;
    margin: 30px 0 !important;
  }
  
  th {
    background: linear-gradient(135deg, #667eea 0%, #764ba2 100%) !important;
    color: white !important;
    padding: 12px !important;
    text-align: left !important;
    font-weight: 600 !important;
  }
  
  td {
    padding: 12px !important;
    border: 1px solid #e0e0e0 !important;
  }
  
  tr:nth-child(even) {
    background: #f8f9fa !important;
  }
  
  tr:hover {
    background: #e3f2fd !important;
  }
  
  a:hover {
    transform: translateY(-2px) !important;
  }
</style>
"""

# changelog.md 模板
CHANGELOG_TEMPLATE = """# 更新日志

## v0.0.9 (2026-03-16)

### 文档更新与脚本引擎支持

- 支持 JavaScript 和 Lua 脚本引擎
- 支持 require 功能，实现脚本模块化
- 完善文档结构和 API 参考
- 修复各种兼容性问题
- 添加模块白名单功能，解决 Windows 编译命令过长问题
- 优化代码结构和性能
- 添加无痛迁移特性，支持其他平台代码的无缝迁移

## v0.0.5 (2026-03-13)

### 完成全量测试并修复完善

- 对所有模块进行全量测试
- 修复测试中发现的问题
- 完善代码实现和文档
- 优化代码结构和性能

## v0.0.4 (2026-03-13)

### 功能增强与修复

- 修复初始化问题
- 优化方法注册机制
- 完善错误处理
- 提升代码稳定性

## v0.0.3 (2026-03-13)

### 修复编译问题

- 修复编译错误
- 优化依赖管理
- 提升构建稳定性

## v0.0.2 (2026-03-13)

### 完善 API

- 完善 API 实现
- 修复参数错误
- 优化代码结构

## v0.0.1 (2026-03-13)

### 项目初始化

- 初始化 AutoGo ScriptEngine 项目
- 实现基本架构
- 添加核心模块
"""

def find_all_readmes(root_dir):
    """查找所有 README.md 文件"""
    readmes = []
    for root, dirs, files in os.walk(root_dir):
        # 跳过隐藏目录和特定目录
        dirs[:] = [d for d in dirs if not d.startswith('.') and d not in ['docs', 'node_modules', '.git', 'scripts']]
        
        for file in files:
            if file.lower() == 'readme.md':
                readmes.append(os.path.join(root, file))
    return readmes

def generate_sidebar_content(project_root, readmes):
    """生成侧边栏内容"""
    # 按目录分组
    grouped = {}
    for md_path in readmes:
        rel_path = os.path.relpath(md_path, project_root)
        dir_name = os.path.dirname(rel_path)
        
        if dir_name not in grouped:
            grouped[dir_name] = []
        grouped[dir_name].append(rel_path)
    
    # 生成 API 文档部分
    api_docs = []
    
    # 处理 js_engine 模块
    if 'js_engine/model' in grouped:
        for file_path in sorted(grouped['js_engine/model']):
            module_name = os.path.basename(os.path.dirname(file_path))
            api_docs.append(f"  - [js/{module_name}]({file_path})")
    
    # 处理 lua_engine 模块
    if 'lua_engine/model' in grouped:
        for file_path in sorted(grouped['lua_engine/model']):
            module_name = os.path.basename(os.path.dirname(file_path))
            api_docs.append(f"  - [lua/{module_name}]({file_path})")
    
    return '\n'.join(api_docs)

def copy_docsify_files(docs_dir, docs_copy_dir):
    """复制 docsify 所需的文件"""
    # 复制样式文件
    style_src = os.path.join(docs_copy_dir, 'style')
    style_dst = os.path.join(docs_dir, 'style')
    if os.path.exists(style_src):
        shutil.copytree(style_src, style_dst, dirs_exist_ok=True)
        print(f"✓ 复制样式文件: {style_src} -> {style_dst}")
    
    # 复制图标文件
    icon_src = os.path.join(docs_copy_dir, 'icon.svg')
    icon_dst = os.path.join(docs_dir, 'icon.svg')
    if os.path.exists(icon_src):
        shutil.copy2(icon_src, icon_dst)
        print(f"✓ 复制图标文件: {icon_src} -> {icon_dst}")
    
    # 复制 JavaScript 文件
    js_files = [
        ('docsify@4', 'docsify.min.js'),
        ('docsify-copy-code', 'docsify-copy-code'),
        ('search.min.js', 'search.min.js'),
        ('prism-go.min.js', 'prism-go.min.js'),
        ('docsify-footer.min.js', 'docsify-footer.min.js'),
        ('docsify-plugin-flexible-alerts.min.js', 'docsify-plugin-flexible-alerts.min.js'),
        ('zoom-image.js', 'zoom-image.js')
    ]
    
    for src_name, dst_name in js_files:
        js_src = os.path.join(docs_copy_dir, src_name)
        js_dst = os.path.join(docs_dir, dst_name)
        if os.path.exists(js_src):
            shutil.copy2(js_src, js_dst)
            print(f"✓ 复制 JavaScript 文件: {src_name} -> {dst_name}")

def main():
    """主函数"""
    print("=" * 60)
    print("Markdown 转 docsify 文档系统")
    print("=" * 60)
    
    # 获取项目根目录
    script_dir = os.path.dirname(os.path.abspath(__file__))
    project_root = os.path.dirname(script_dir)
    
    print(f"\n项目根目录: {project_root}")
    
    # 查找所有 README.md 文件
    print("\n正在查找所有 README.md 文件...")
    readmes = find_all_readmes(project_root)
    print(f"找到 {len(readmes)} 个 README.md 文件\n")
    
    # 准备 docs 目录
    docs_dir = os.path.join(project_root, 'docs')
    os.makedirs(docs_dir, exist_ok=True)
    
    # 复制 docsify 所需文件
    docs_copy_dir = os.path.join(project_root, 'docs copy')
    if os.path.exists(docs_copy_dir):
        print("正在复制 docsify 所需文件...")
        copy_docsify_files(docs_dir, docs_copy_dir)
    else:
        print("警告: docs copy 目录不存在，跳过复制 docsify 文件")
    
    # 生成 index.html
    print("\n正在生成 index.html...")
    index_html_path = os.path.join(docs_dir, 'index.html')
    with open(index_html_path, 'w', encoding='utf-8') as f:
        f.write(INDEX_HTML_TEMPLATE)
    print(f"✓ 生成: {index_html_path}")
    
    # 生成 _sidebar.md
    print("正在生成 _sidebar.md...")
    sidebar_content = generate_sidebar_content(project_root, readmes)
    sidebar_content = SIDEBAR_TEMPLATE.format(api_docs=sidebar_content)
    sidebar_path = os.path.join(docs_dir, '_sidebar.md')
    with open(sidebar_path, 'w', encoding='utf-8') as f:
        f.write(sidebar_content)
    print(f"✓ 生成: {sidebar_path}")
    
    # 生成 README.md
    print("正在生成 README.md...")
    readme_path = os.path.join(docs_dir, 'README.md')
    with open(readme_path, 'w', encoding='utf-8') as f:
        f.write(README_TEMPLATE)
    print(f"✓ 生成: {readme_path}")
    
    # 生成 changelog.md
    print("正在生成 changelog.md...")
    changelog_path = os.path.join(docs_dir, 'changelog.md')
    with open(changelog_path, 'w', encoding='utf-8') as f:
        f.write(CHANGELOG_TEMPLATE)
    print(f"✓ 生成: {changelog_path}")
    
    # 复制所有 README.md 文件到 docs 目录
    print("\n正在复制 README.md 文件...")
    for md_path in readmes:
        rel_path = os.path.relpath(md_path, project_root)
        dest_path = os.path.join(docs_dir, rel_path)
        
        # 创建目标目录
        os.makedirs(os.path.dirname(dest_path), exist_ok=True)
        
        # 复制文件
        shutil.copy2(md_path, dest_path)
        print(f"✓ 复制: {rel_path}")
    
    # 输出统计信息
    print("\n" + "=" * 60)
    print("转换完成！")
    print(f"文档数量: {len(readmes)} 个")
    print("=" * 60)
    print(f"\ndocsify 文档已生成到: {docs_dir}")
    print("\n使用浏览器打开以下文件查看文档：")
    print(f"  {os.path.join(docs_dir, 'index.html')}")
    print("\n或使用 docsify 命令启动本地服务器：")
    print("  cd docs && docsify serve")

if __name__ == '__main__':
    main()