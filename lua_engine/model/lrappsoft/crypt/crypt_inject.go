//go:build ignore
// +build ignore

package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"

	"github.com/ZingYao/autogo_scriptengine/lua_engine/model"

	lua "github.com/yuin/gopher-lua"
)

// CryptModule cryptLib 模块（懒人精灵兼容）
type CryptModule struct{}

// Name 返回模块名称
func (m *CryptModule) Name() string {
	return "cryptLib"
}

// IsAvailable 检查模块是否可用
func (m *CryptModule) IsAvailable() bool {
	return true
}

// Register 向引擎注册方法
func (m *CryptModule) Register(engine model.Engine) error {
	state := engine.GetState()

	// 创建 cryptLib 表
	cryptLibTable := state.NewTable()
	state.SetGlobal("cryptLib", cryptLibTable)

	// 注册 cryptLib.aes_crypt - AES 加密/解密
	cryptLibTable.RawSetString("aes_crypt", state.NewFunction(func(L *lua.LState) int {
		data := L.CheckString(1)
		key := L.CheckString(2)
		operation := L.CheckString(3)
		mode := L.CheckString(4)

		// 检查参数
		if len(key) != 16 && len(key) != 24 && len(key) != 32 {
			L.Push(lua.LNil)
			L.Push(lua.LString("密钥长度必须为 16/24/32 字节"))
			return 2
		}

		// 获取可选参数
		iv := ""
		if L.GetTop() >= 5 {
			iv = L.CheckString(5)
		}

		padding := true
		if L.GetTop() >= 6 {
			padding = L.CheckBool(6)
		}

		// 检查 iv 参数
		if mode != "ecb" && iv == "" {
			L.Push(lua.LNil)
			L.Push(lua.LString("非 ECB 模式需要 iv 参数"))
			return 2
		}

		if mode != "ecb" && len(iv) != 16 {
			L.Push(lua.LNil)
			L.Push(lua.LString("iv 长度必须为 16 字节"))
			return 2
		}

		// 执行加密/解密
		result, err := aesCrypt(data, key, operation, mode, iv, padding)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LString(result))
		return 1
	}))

	// 注册 cryptLib.aes_keygen - 生成 AES 密钥
	cryptLibTable.RawSetString("aes_keygen", state.NewFunction(func(L *lua.LState) int {
		keyLength := L.CheckInt(1)

		if keyLength != 16 && keyLength != 24 && keyLength != 32 {
			L.Push(lua.LNil)
			L.Push(lua.LString("密钥长度必须为 16/24/32"))
			return 2
		}

		key := make([]byte, keyLength)
		_, err := rand.Read(key)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LString(string(key)))
		return 1
	}))

	// 注册 cryptLib.aes_ivgen - 生成随机 IV
	cryptLibTable.RawSetString("aes_ivgen", state.NewFunction(func(L *lua.LState) int {
		iv := make([]byte, 16)
		_, err := rand.Read(iv)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LString(string(iv)))
		return 1
	}))

	// 注册 cryptLib.rsa_generate_key - 生成 RSA 密钥对
	cryptLibTable.RawSetString("rsa_generate_key", state.NewFunction(func(L *lua.LState) int {
		keyBits := 2048
		if L.GetTop() >= 1 {
			keyBits = L.CheckInt(1)
		}

		// 生成私钥
		privateKey, err := rsa.GenerateKey(rand.Reader, keyBits)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 3
		}

		// 生成公钥 PEM
		pubKeyBytes := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
		pubKeyPEM := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubKeyBytes,
		})

		// 生成私钥 PEM
		privKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
		privKeyPEM := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privKeyBytes,
		})

		L.Push(lua.LString(string(pubKeyPEM)))
		L.Push(lua.LString(string(privKeyPEM)))
		return 2
	}))

	// 注册 cryptLib.rsa_encrypt - RSA 加密
	cryptLibTable.RawSetString("rsa_encrypt", state.NewFunction(func(L *lua.LState) int {
		data := L.CheckString(1)
		key := L.CheckString(2)

		isPublicKey := true
		if L.GetTop() >= 3 {
			isPublicKey = L.CheckBool(3)
		}

		result, err := rsaEncrypt(data, key, isPublicKey)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LString(result))
		return 1
	}))

	// 注册 cryptLib.rsa_decrypt - RSA 解密
	cryptLibTable.RawSetString("rsa_decrypt", state.NewFunction(func(L *lua.LState) int {
		data := L.CheckString(1)
		key := L.CheckString(2)

		isPublicKey := false
		if L.GetTop() >= 3 {
			isPublicKey = L.CheckBool(3)
		}

		result, err := rsaDecrypt(data, key, isPublicKey)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}

		L.Push(lua.LString(result))
		return 1
	}))

	// 注册 cryptLib.base64_encode - Base64 编码
	cryptLibTable.RawSetString("base64_encode", state.NewFunction(func(L *lua.LState) int {
		data := L.CheckString(1)
		encoded := base64.StdEncoding.EncodeToString([]byte(data))
		L.Push(lua.LString(encoded))
		return 1
	}))

	// 注册 cryptLib.base64_decode - Base64 解码
	cryptLibTable.RawSetString("base64_decode", state.NewFunction(func(L *lua.LState) int {
		data := L.CheckString(1)
		decoded, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 2
		}
		L.Push(lua.LString(string(decoded)))
		return 1
	}))

	// 注册 cryptLib.md5 - MD5 哈希
	cryptLibTable.RawSetString("md5", state.NewFunction(func(L *lua.LState) int {
		data := L.CheckString(1)
		hash := md5.Sum([]byte(data))
		L.Push(lua.LString(hex.EncodeToString(hash[:])))
		return 1
	}))

	// 注册 cryptLib.sha256 - SHA256 哈希
	cryptLibTable.RawSetString("sha256", state.NewFunction(func(L *lua.LState) int {
		data := L.CheckString(1)
		hash := sha256.Sum256([]byte(data))
		L.Push(lua.LString(hex.EncodeToString(hash[:])))
		return 1
	}))

	// 注册 cryptLib.sha512 - SHA512 哈希
	cryptLibTable.RawSetString("sha512", state.NewFunction(func(L *lua.LState) int {
		data := L.CheckString(1)
		hash := sha512.Sum512([]byte(data))
		L.Push(lua.LString(hex.EncodeToString(hash[:])))
		return 1
	}))

	// 注册 cryptLib.hmac_sha256 - HMAC-SHA256
	cryptLibTable.RawSetString("hmac_sha256", state.NewFunction(func(L *lua.LState) int {
		data := L.CheckString(1)
		key := L.CheckString(2)
		h := hmac.New(sha256.New, []byte(key))
		h.Write([]byte(data))
		hash := h.Sum(nil)
		L.Push(lua.LString(hex.EncodeToString(hash)))
		return 1
	}))

	// 注册 cryptLib.rsa_keygen - 生成 RSA 密钥对（别名）
	cryptLibTable.RawSetString("rsa_keygen", state.NewFunction(func(L *lua.LState) int {
		keyBits := 2048
		if L.GetTop() >= 1 {
			keyBits = L.CheckInt(1)
		}

		privateKey, err := rsa.GenerateKey(rand.Reader, keyBits)
		if err != nil {
			L.Push(lua.LNil)
			L.Push(lua.LNil)
			L.Push(lua.LString(err.Error()))
			return 3
		}

		pubKeyBytes := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
		pubKeyPEM := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: pubKeyBytes,
		})

		privKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
		privKeyPEM := pem.EncodeToMemory(&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: privKeyBytes,
		})

		L.Push(lua.LString(string(pubKeyPEM)))
		L.Push(lua.LString(string(privKeyPEM)))
		return 2
	}))

	// 注册到方法注册表
	engine.RegisterMethod("cryptLib.aes_crypt", "AES 加密/解密", func(data, key, operation, mode string, options ...interface{}) (string, error) {
		iv := ""
		if len(options) >= 1 {
			optionIV, ok := options[0].(string)
			if !ok {
				return "", errors.New("iv 参数必须为字符串")
			}
			iv = optionIV
		}
		padding := true
		if len(options) >= 2 {
			optionPadding, ok := options[1].(bool)
			if !ok {
				return "", errors.New("padding 参数必须为布尔值")
			}
			padding = optionPadding
		}
		return aesCrypt(data, key, operation, mode, iv, padding)
	}, true)

	engine.RegisterMethod("cryptLib.aes_keygen", "生成 AES 密钥", func(keyLength int) (string, error) {
		if keyLength != 16 && keyLength != 24 && keyLength != 32 {
			return "", errors.New("密钥长度必须为 16/24/32")
		}
		key := make([]byte, keyLength)
		_, err := rand.Read(key)
		if err != nil {
			return "", err
		}
		return string(key), nil
	}, true)

	engine.RegisterMethod("cryptLib.aes_ivgen", "生成随机 IV", func() (string, error) {
		iv := make([]byte, 16)
		_, err := rand.Read(iv)
		if err != nil {
			return "", err
		}
		return string(iv), nil
	}, true)

	engine.RegisterMethod("cryptLib.rsa_generate_key", "生成 RSA 密钥对", func(keyBits ...int) (string, string, error) {
		return generateRSAKeyPair(keyBits...)
	}, true)

	engine.RegisterMethod("cryptLib.rsa_keygen", "生成 RSA 密钥对（兼容别名）", func(keyBits ...int) (string, string, error) {
		return generateRSAKeyPair(keyBits...)
	}, true)

	engine.RegisterMethod("cryptLib.rsa_encrypt", "RSA 加密", func(data, key string, isPublicKey bool) (string, error) {
		return rsaEncrypt(data, key, isPublicKey)
	}, true)

	engine.RegisterMethod("cryptLib.rsa_decrypt", "RSA 解密", func(data, key string, isPublicKey bool) (string, error) {
		return rsaDecrypt(data, key, isPublicKey)
	}, true)
	engine.RegisterMethod("cryptLib.base64_encode", "Base64 编码", func(data string) string {
		return base64.StdEncoding.EncodeToString([]byte(data))
	}, true)
	engine.RegisterMethod("cryptLib.base64_decode", "Base64 解码", func(data string) (string, error) {
		decoded, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			return "", err
		}
		return string(decoded), nil
	}, true)
	engine.RegisterMethod("cryptLib.md5", "MD5 哈希", func(data string) string {
		hash := md5.Sum([]byte(data))
		return hex.EncodeToString(hash[:])
	}, true)
	engine.RegisterMethod("cryptLib.sha256", "SHA256 哈希", func(data string) string {
		hash := sha256.Sum256([]byte(data))
		return hex.EncodeToString(hash[:])
	}, true)
	engine.RegisterMethod("cryptLib.sha512", "SHA512 哈希", func(data string) string {
		hash := sha512.Sum512([]byte(data))
		return hex.EncodeToString(hash[:])
	}, true)
	engine.RegisterMethod("cryptLib.hmac_sha256", "HMAC-SHA256", func(data, key string) string {
		h := hmac.New(sha256.New, []byte(key))
		h.Write([]byte(data))
		return hex.EncodeToString(h.Sum(nil))
	}, true)

	return nil
}

// aesCrypt AES 加密/解密
func aesCrypt(data, key, operation, mode, iv string, padding bool) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	dataBytes := []byte(data)
	var result []byte

	switch mode {
	case "ecb":
		if operation == "encrypt" {
			dataBytes = pkcs7Pad(dataBytes, block.BlockSize())
			result = make([]byte, len(dataBytes))
			for i := 0; i < len(dataBytes); i += block.BlockSize() {
				block.Encrypt(result[i:i+block.BlockSize()], dataBytes[i:i+block.BlockSize()])
			}
		} else {
			result = make([]byte, len(dataBytes))
			for i := 0; i < len(dataBytes); i += block.BlockSize() {
				block.Decrypt(result[i:i+block.BlockSize()], dataBytes[i:i+block.BlockSize()])
			}
			if padding {
				result = pkcs7Unpad(result)
			}
		}
	case "cbc", "cfb", "ofb", "ctr":
		ivBytes := []byte(iv)
		stream := cipher.NewCTR(block, ivBytes)
		result = make([]byte, len(dataBytes))
		stream.XORKeyStream(result, dataBytes)
	default:
		return "", fmt.Errorf("不支持的加密模式: %s", mode)
	}

	return string(result), nil
}

// generateRSAKeyPair 生成 RSA 公钥与私钥 PEM 文本
func generateRSAKeyPair(keyBits ...int) (string, string, error) {
	bits := 2048
	if len(keyBits) > 0 && keyBits[0] != 0 {
		bits = keyBits[0]
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	pubKeyBytes := x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	pubKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKeyBytes,
	})
	privKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privKeyBytes,
	})
	return string(pubKeyPEM), string(privKeyPEM), nil
}

// pkcs7Pad PKCS#7 填充
func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := make([]byte, padding)
	for i := 0; i < padding; i++ {
		padText[i] = byte(padding)
	}
	return append(data, padText...)
}

// pkcs7Unpad PKCS#7 去填充
func pkcs7Unpad(data []byte) []byte {
	if len(data) == 0 {
		return data
	}
	padding := int(data[len(data)-1])
	if padding < 1 || padding > len(data) {
		return data
	}
	for i := len(data) - padding; i < len(data); i++ {
		if int(data[i]) != padding {
			return data
		}
	}
	return data[:len(data)-padding]
}

// rsaEncrypt RSA 加密
func rsaEncrypt(data, key string, isPublicKey bool) (string, error) {
	var rsaKey *rsa.PublicKey
	var err error

	if isPublicKey {
		block, _ := pem.Decode([]byte(key))
		if block == nil {
			return "", errors.New("无效的公钥格式")
		}
		rsaKey, err = x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return "", err
		}
	} else {
		block, _ := pem.Decode([]byte(key))
		if block == nil {
			return "", errors.New("无效的私钥格式")
		}
		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return "", err
		}
		rsaKey = &privateKey.PublicKey
	}

	hash := sha256.New()
	encrypted, err := rsa.EncryptOAEP(hash, rand.Reader, rsaKey, []byte(data), nil)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

// rsaDecrypt RSA 解密
func rsaDecrypt(data, key string, isPublicKey bool) (string, error) {
	var rsaKey *rsa.PrivateKey
	var err error

	if isPublicKey {
		return "", errors.New("使用公钥解密不支持")
	} else {
		block, _ := pem.Decode([]byte(key))
		if block == nil {
			return "", errors.New("无效的私钥格式")
		}
		rsaKey, err = x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return "", err
		}
	}

	// 尝试 base64 解码
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		// 如果 base64 解码失败，直接使用原始数据
		decodedData = []byte(data)
	}

	hash := sha256.New()
	decrypted, err := rsa.DecryptOAEP(hash, rand.Reader, rsaKey, decodedData, nil)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}
