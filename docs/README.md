# AutoGo Script Engine

## 简介

AutoGo Script Engine 是为 AutoGo 提供 JavaScript 和 Lua 脚本引擎支持的扩展方案，让开发者可以用熟悉的脚本语言编写自动化任务。

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

## 安装

```bash
go get github.com/ZingYao/autogo_scriptengine@v0.0.9
```

## 与 AutoGo 的关系

本项目是 AutoGo 的扩展方案，通过封装 AutoGo 提供的原生 API，为开发者提供更灵活的脚本编写方式：

- **AutoGo** - 提供 Android 自动化的核心能力（无障碍服务、图像识别、触摸模拟等）
- **ScriptEngine** - 为 AutoGo 添加脚本语言支持，让开发者可以用 JavaScript 或 Lua 编写自动化脚本

## 快速开始

1. **安装依赖**：使用上面的命令安装 AutoGo Script Engine
2. **选择引擎**：根据您的喜好选择 JavaScript 或 Lua 引擎
3. **编写脚本**：使用所选脚本语言编写自动化任务
4. **运行脚本**：通过 AutoGo 执行您的脚本

## 许可证

MIT License

---

<style>
  :root {
    --theme-color: #4CAF50;
  }
  
  .cover {
    background: linear-gradient(to left bottom, hsl(216, 100%, 85%) 0%, hsl(107, 100%, 85%) 100%) !important;
  }
  
  h1 {
    color: #2E7D32 !important;
    font-size: 2.5rem !important;
    margin-bottom: 1rem !important;
  }
  
  h2 {
    color: #00897B !important;
    font-size: 1.5rem !important;
    margin-bottom: 2rem !important;
  }
  
  blockquote {
    border-left: 4px solid var(--theme-color) !important;
    background: rgba(76, 175, 80, 0.1) !important;
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
    background: linear-gradient(to left bottom, hsl(216, 100%, 65%) 0%, hsl(107, 100%, 65%) 100%) !important;
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
    background: #f1f8e9 !important;
  }
  
  tr:hover {
    background: #e8f5e8 !important;
  }
  
  a:hover {
    transform: translateY(-2px) !important;
  }
</style>