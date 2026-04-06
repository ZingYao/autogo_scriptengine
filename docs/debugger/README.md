# AutoGo 快速调试工具

## 概述

AutoGo 快速调试工具是一个现代化的 AutoGo ScriptEngine 开发调试工具，提供友好的 TUI 界面和强大的调试功能，帮助开发者更高效地开发和调试 AutoGo 脚本。

## 项目地址

```
https://github.com/ZingYao/autogo_scriptengine_debugger.git
```

## 功能特性

### 🎨 现代化 TUI 界面

- 直观的终端用户界面
- 实时日志输出显示
- 支持鼠标操作
- 彩色语法高亮

### 🚀 核心功能

- **项目管理**: 初始化、编译、部署 AutoGo 项目
- **设备管理**: 自动检测设备、连接设备、获取设备信息
- **脚本运行**: 支持 Lua 和 JavaScript 脚本
- **实时调试**: 日志输出、暂停、恢复、停止脚本
- **AG 管理**: 自动下载和更新 AutoGo 工具

### 🔧 开发工具

- **代码风格支持**: AutoGo、LrAppSoft、NodeJS
- **自动生成项目模板**
- **配置文件管理**
- **多设备支持**

### 🌍 跨平台支持

- ✅ Windows (AMD64)
- ✅ macOS (ARM64 - M1/M2/M3)
- ✅ macOS (AMD64 - Intel)

## 安装

### 从 Release 下载

访问 [Releases](https://github.com/ZingYao/autogo_scriptengine_debugger/releases) 页面下载最新版本：

| 平台 | 文件名 | 说明 |
|------|--------|------|
| Windows | AutoGoScriptEngineDebugger-Windows.zip | 包含 .exe 可执行文件 |
| macOS ARM | AutoGoScriptEngineDebugger-macOS-ARM.tar.gz | 适用于 M1/M2/M3 芯片 |
| macOS AMD | AutoGoScriptEngineDebugger-macOS-AMD.tar.gz | 适用于 Intel 芯片 |

### 安装步骤

**Windows:**
```bash
# 1. 解压缩
unzip AutoGoScriptEngineDebugger-Windows.zip

# 2. 运行
./AutoGoScriptEngineDebugger.exe
```

**macOS:**
```bash
# 1. 解压缩
tar -xzf AutoGoScriptEngineDebugger-macOS-*.tar.gz

# 2. 添加执行权限
chmod +x AutoGoScriptEngineDebugger*

# 3. 运行
./AutoGoScriptEngineDebuggerArm  # 或 AutoGoScriptEngineDebuggerAmd
```

### 从源码构建

**前置要求:**
- Go 1.21 或更高版本
- Git

**构建步骤:**
```bash
# 克隆仓库
git clone https://github.com/ZingYao/autogo_scriptengine_debugger.git
cd autogo_scriptengine_debugger

# 编译
go build -o AutoGoScriptEngineDebugger .

# 运行
./AutoGoScriptEngineDebugger
```

## 快速开始

### 1. 启动程序

```bash
# TUI 模式（推荐）
./AutoGoScriptEngineDebugger

# CLI 模式
./AutoGoScriptEngineDebugger --cli

# 直接运行脚本
./AutoGoScriptEngineDebugger script.lua
```

### 2. 初始化项目

1. 启动程序后，按 `i` 键选择"项目初始化"
2. 输入 Module 名称（如：example.com/myproject）
3. 选择目标平台（android/ios）
4. 等待初始化完成

### 3. 连接设备

1. 确保设备通过 USB 连接或网络连接
2. 按 `2` 键进入"设备管理"
3. 选择"查看已连接设备"
4. 选择要使用的设备

### 4. 运行脚本

1. 将脚本文件放到项目目录的 `scripts/` 文件夹
2. 按 `1` 键进入"运行管理"
3. 选择"选择脚本文件"
4. 选择要运行的脚本
5. 选择"运行脚本"

## 快捷键

### 主菜单

| 按键 | 功能 |
|------|------|
| 1-4 | 快速选择菜单项 |
| i | 项目初始化 |
| h | 帮助 |
| q | 退出 |
| l | 查看调试器日志 |
| d | 查看项目运行日志 |
| r | 刷新页面 |
| F9 | 切换鼠标模式 |
| Ctrl+Q | 直接退出 |
| Ctrl+1-4 | 快速执行菜单项 |

### 日志浏览

| 按键 | 功能 |
|------|------|
| ↑/↓ | 滚动 |
| PgUp/PgDn | 翻页 |
| Home/End | 跳转首尾 |
| Tab | 切换日志 |
| r | 刷新页面 |
| c | 清空日志 |
| ESC | 返回菜单 |

## 项目结构

```
autogo_scriptengine_debugger/
├── main.go              # 程序入口
├── agmanager/           # AG 工具管理
│   ├── agmanager.go     # 下载、安装 AG
│   └── version_list.go  # 版本列表
├── config/              # 配置管理
│   └── config.go        # 加载、保存配置
├── device/              # 设备管理
│   └── device.go        # ADB 设备操作
├── interactive/         # 交互式输入
│   └── interactive.go   # CLI 模式输入
├── printer/             # 彩色打印
│   └── printer.go       # 日志输出
├── project/             # 项目管理
│   ├── project.go       # 项目操作
│   ├── embed_files.go   # 嵌入模板文件
│   ├── scripts/         # 示例脚本模板
│   └── main.go.code     # main.go 模板
├── script/              # 脚本操作
│   └── script.go        # 部署脚本
├── tui/                 # TUI 界面
│   └── tui.go           # 终端界面
└── .github/
    └── workflows/
        └── release.yml  # 自动发布配置
```

## 开发

### 构建

```bash
# 本地构建
go build -o AutoGoScriptEngineDebugger .

# 跨平台编译
# Windows
GOOS=windows GOARCH=amd64 go build -o AutoGoScriptEngineDebugger.exe .

# macOS ARM
GOOS=darwin GOARCH=arm64 go build -o AutoGoScriptEngineDebuggerArm .

# macOS AMD
GOOS=darwin GOARCH=amd64 go build -o AutoGoScriptEngineDebuggerAmd .
```

### 发布新版本

1. 更新代码并提交
2. 创建版本标签：
   ```bash
   git tag v1.0.0
   git push origin v1.0.0
   ```
3. GitHub Actions 会自动构建并创建 Release

## 配置

配置文件保存在 `当前执行目录/.autogo_scriptengine_debugger/config.json`

配置项：
- `codeStyle`: 代码风格（autogo/lrappsoft/nodejs）
- `deviceServiceUrl`: 设备服务地址
- `deviceId`: 设备 ID
- `deviceIp`: 设备 IP
- `devicePort`: 设备端口
- `projectPath`: 项目路径
- `agPath`: AG 工具路径

## 密钥管理

### 对称加密密钥（AES）生成指南

#### 什么是 AES？

AES（Advanced Encryption Standard）是一种对称加密算法，加密和解密使用相同的密钥。本系统使用 AES-GCM 模式，提供加密和完整性验证。

#### 支持的密钥长度

- **AES-128**: 16 字节密钥（128 位）
- **AES-192**: 24 字节密钥（192 位）
- **AES-256**: 32 字节密钥（256 位）- **推荐**

#### 方法一：使用 OpenSSL 生成（推荐）

```bash
# 生成 256 位（32 字节）AES 密钥，并转换为 Base64
openssl rand -base64 32

# 输出示例：
# kYp3s6v9y$B&E)H@McQfThWmZq4t7w!z
```

#### 方法二：使用 Python 生成

```python
import base64
import os

# 生成 32 字节（256 位）随机密钥
key = os.urandom(32)

# 转换为 Base64 编码
key_base64 = base64.b64encode(key).decode('utf-8')
print(f"AES 密钥 (Base64): {key_base64}")
```

#### 方法三：使用 Go 生成

```go
package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateAESKey() (string, error) {
	// 生成 32 字节（256 位）密钥
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		return "", err
	}
	
	// 转换为 Base64
	return base64.StdEncoding.EncodeToString(key), nil
}

func main() {
	key, err := generateAESKey()
	if err != nil {
		panic(err)
	}
	fmt.Printf("AES 密钥 (Base64): %s\n", key)
}
```

### 非对称加密密钥对（RSA）生成指南

#### 什么是 RSA？

RSA 是一种非对称加密算法，使用一对密钥：
- **公钥**：用于加密数据
- **私钥**：用于解密数据

#### 支持的密钥长度

- **RSA-2048**: 2048 位密钥 - **推荐**
- **RSA-3072**: 3072 位密钥
- **RSA-4096**: 4096 位密钥（更安全，但性能较低）

#### 方法一：使用 OpenSSL 生成（推荐）

```bash
# 1. 生成 2048 位 RSA 私钥
openssl genrsa -out private_key.pem 2048

# 2. 从私钥提取公钥
openssl rsa -in private_key.pem -pubout -out public_key.pem

# 3. 查看私钥内容
cat private_key.pem

# 4. 查看公钥内容
cat public_key.pem
```

### 安全最佳实践

#### 密钥存储

1. **不要将密钥硬编码在代码中**
   ```go
   // ❌ 错误做法
   const key = "my-secret-key-123"
   
   // ✅ 正确做法：从环境变量或配置文件读取
   key := os.Getenv("ENCRYPTION_KEY")
   ```

2. **使用环境变量存储密钥**
   ```bash
   # Linux/macOS
   export ENCRYPTION_KEY="your-base64-encoded-key"
   
   # Windows (PowerShell)
   $env:ENCRYPTION_KEY="your-base64-encoded-key"
   ```

3. **使用密钥管理服务（生产环境推荐）**
   - AWS KMS (Key Management Service)
   - Google Cloud KMS
   - Azure Key Vault
   - HashiCorp Vault

#### 密钥传输

1. **使用安全通道传输密钥**
   - HTTPS
   - SSH
   - 加密邮件

2. **不要通过明文渠道传输密钥**
   - 禁止通过普通邮件发送
   - 禁止通过即时通讯软件发送
   - 禁止提交到版本控制系统

#### 密钥轮换

1. **定期更换密钥**
   - 建议每 90 天更换一次密钥
   - 如果密钥可能泄露，立即更换

2. **密钥版本管理**
   - 保留旧密钥版本以解密旧数据
   - 使用新密钥加密新数据

## 常见问题

### Q: 如何验证密钥是否正确？

可以通过加密一段已知数据，然后解密验证：
```go
// 加密
encrypted, _ := encryptAES([]byte("test"), key)

// 解密
decrypted, _ := decryptAES(encrypted, key)

// 验证
if string(decrypted) == "test" {
    fmt.Println("密钥验证成功！")
}
```

### Q: 快速调试工具支持哪些脚本语言？

- ✅ Lua 脚本
- ✅ JavaScript 脚本

### Q: 如何在调试工具中使用字节码？

1. 首先编译 Lua 脚本为字节码
2. 将生成的 `.gluac` 文件放到项目的 `scripts/` 目录
3. 在调试工具中选择运行该字节码文件

### Q: 调试工具如何处理加密脚本？

1. 确保你有正确的加密密钥
2. 在配置文件中设置密钥
3. 调试工具会自动解密并执行加密的脚本

## 更新日志

### v1.0.0 (2026-04-06)

- 初始版本
- 支持项目初始化和管理
- 支持设备连接和管理
- 支持 Lua 和 JavaScript 脚本运行
- 支持实时日志输出
- 支持 AG 工具自动下载和更新

## 贡献

欢迎提交 Issue 和 Pull Request 来改进这个项目！

## 许可证

MIT 许可证
