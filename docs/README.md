---
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