# 密钥生成指南

本文档详细说明如何生成和管理加密密钥，包括对称加密（AES）和非对称加密（RSA）。

## 目录

1. [对称加密密钥（AES）生成指南](#对称加密密钥aes生成指南)
2. [非对称加密密钥对（RSA）生成指南](#非对称加密密钥对rsa生成指南)
3. [安全最佳实践](#安全最佳实践)
4. [密钥使用示例](#密钥使用示例)

---

## 对称加密密钥（AES）生成指南

### 什么是 AES？

AES（Advanced Encryption Standard）是一种对称加密算法，加密和解密使用相同的密钥。本系统使用 AES-GCM 模式，提供加密和完整性验证。

### 支持的密钥长度

- **AES-128**: 16 字节密钥（128 位）
- **AES-192**: 24 字节密钥（192 位）
- **AES-256**: 32 字节密钥（256 位）- **推荐**

### 方法一：使用 OpenSSL 生成（推荐）

```bash
# 生成 256 位（32 字节）AES 密钥，并转换为 Base64
openssl rand -base64 32

# 输出示例：
# kYp3s6v9y$B&E)H@McQfThWmZq4t7w!z
```

### 方法二：使用 Python 生成

```python
import base64
import os

# 生成 32 字节（256 位）随机密钥
key = os.urandom(32)

# 转换为 Base64 编码
key_base64 = base64.b64encode(key).decode('utf-8')
print(f"AES 密钥 (Base64): {key_base64}")
```

### 方法三：使用 Go 生成

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

### 方法四：使用在线工具

可以使用以下在线工具生成随机字节：
- https://www.uuidgenerator.net/dev-tools/random-string-generator
- 设置长度为 32，字符集选择 "Hex" 或 "Base64"

---

## 非对称加密密钥对（RSA）生成指南

### 什么是 RSA？

RSA 是一种非对称加密算法，使用一对密钥：
- **公钥**：用于加密数据
- **私钥**：用于解密数据

### 支持的密钥长度

- **RSA-2048**: 2048 位密钥 - **推荐**
- **RSA-3072**: 3072 位密钥
- **RSA-4096**: 4096 位密钥（更安全，但性能较低）

### 方法一：使用 OpenSSL 生成（推荐）

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

### 方法二：生成 PKCS#8 格式的私钥

```bash
# 生成 PKCS#8 格式的私钥（推荐用于 Go）
openssl genpkey -algorithm RSA -out private_key_pkcs8.pem -pkeyopt rsa_keygen_bits:2048

# 从私钥提取公钥
openssl rsa -in private_key_pkcs8.pem -pubout -out public_key.pem
```

### 方法三：使用 Python 生成

```python
from cryptography.hazmat.primitives import serialization
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.backends import default_backend

# 生成 RSA 密钥对
private_key = rsa.generate_private_key(
    public_exponent=65537,
    key_size=2048,
    backend=default_backend()
)

# 导出私钥（PEM 格式）
private_pem = private_key.private_bytes(
    encoding=serialization.Encoding.PEM,
    format=serialization.PrivateFormat.PKCS8,
    encryption_algorithm=serialization.NoEncryption()
)

# 导出公钥（PEM 格式）
public_key = private_key.public_key()
public_pem = public_key.public_bytes(
    encoding=serialization.Encoding.PEM,
    format=serialization.PublicFormat.SubjectPublicKeyInfo
)

print("私钥:")
print(private_pem.decode())
print("\n公钥:")
print(public_pem.decode())
```

### 方法四：使用 Go 生成

```go
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func generateRSAKeyPair() (string, string, error) {
	// 生成 2048 位 RSA 密钥对
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}

	// 导出私钥（PKCS#8 PEM 格式）
	privateKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
	if err != nil {
		return "", "", err
	}
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	// 导出公钥（PEM 格式）
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", err
	}
	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	return string(privateKeyPEM), string(publicKeyPEM), nil
}

func main() {
	privateKey, publicKey, err := generateRSAKeyPair()
	if err != nil {
		panic(err)
	}
	fmt.Println("私钥:")
	fmt.Println(privateKey)
	fmt.Println("公钥:")
	fmt.Println(publicKey)
}
```

---

## 安全最佳实践

### 密钥存储

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

### 密钥传输

1. **使用安全通道传输密钥**
   - HTTPS
   - SSH
   - 加密邮件

2. **不要通过明文渠道传输密钥**
   - 禁止通过普通邮件发送
   - 禁止通过即时通讯软件发送
   - 禁止提交到版本控制系统

### 密钥轮换

1. **定期更换密钥**
   - 建议每 90 天更换一次密钥
   - 如果密钥可能泄露，立即更换

2. **密钥版本管理**
   - 保留旧密钥版本以解密旧数据
   - 使用新密钥加密新数据

### 访问控制

1. **最小权限原则**
   - 只给需要的人/系统访问密钥的权限
   - 使用角色基础访问控制（RBAC）

2. **审计日志**
   - 记录所有密钥访问和使用
   - 定期审查访问日志

---

## 密钥使用示例

### AES 加密示例（Go）

```go
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

// encryptAES 使用 AES-GCM 加密数据
func encryptAES(plaintext []byte, keyBase64 string) (string, error) {
	// 解码密钥
	key, err := base64.StdEncoding.DecodeString(keyBase64)
	if err != nil {
		return "", err
	}

	// 创建 AES 块
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// 创建 GCM 模式
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// 生成随机 nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// 加密
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	// 返回 Base64 编码的密文
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func main() {
	// 你的 AES 密钥（Base64 编码）
	key := "kYp3s6v9y$B&E)H@McQfThWmZq4t7w!z" // 替换为你的密钥

	// 要加密的数据
	plaintext := []byte("Hello, World!")

	// 加密
	encrypted, err := encryptAES(plaintext, key)
	if err != nil {
		panic(err)
	}

	fmt.Printf("加密结果: %s\n", encrypted)
}
```

### RSA 加密示例（Go）

```go
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

// encryptRSA 使用 RSA-OAEP 加密数据
func encryptRSA(plaintext []byte, publicKeyPEM string) (string, error) {
	// 解析公钥
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return "", fmt.Errorf("无法解析公钥")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		return "", fmt.Errorf("不是 RSA 公钥")
	}

	// 使用 RSA-OAEP 加密
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, rsaPublicKey, plaintext, nil)
	if err != nil {
		return "", err
	}

	// 返回 Base64 编码的密文
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func main() {
	// 你的 RSA 公钥（PEM 格式）
	publicKey := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA...（你的公钥内容）
-----END PUBLIC KEY-----`

	// 要加密的数据
	plaintext := []byte("Hello, World!")

	// 加密
	encrypted, err := encryptRSA(plaintext, publicKey)
	if err != nil {
		panic(err)
	}

	fmt.Printf("加密结果: %s\n", encrypted)
}
```

---

## 常见问题

### Q: AES 和 RSA 应该选择哪个？

**AES（对称加密）：**
- ✅ 加密速度快
- ✅ 适合加密大量数据
- ❌ 密钥需要安全传输
- **推荐用于：脚本代码加密**

**RSA（非对称加密）：**
- ✅ 公钥可以公开，私钥保密
- ✅ 不需要安全传输密钥
- ❌ 加密速度慢
- ❌ 只能加密少量数据（密钥长度限制）
- **推荐用于：密钥交换、数字签名**

### Q: 密钥丢失了怎么办？

如果密钥丢失，加密的数据将无法解密。建议：
1. 定期备份密钥到安全位置
2. 使用密钥管理服务
3. 实施密钥恢复流程

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

---

## 相关文档

- [构建配置说明](./BUILD_CONFIG.md)
- [脚本加载模式](./LOAD_MODES.md)
- [安全配置指南](./SECURITY.md)
