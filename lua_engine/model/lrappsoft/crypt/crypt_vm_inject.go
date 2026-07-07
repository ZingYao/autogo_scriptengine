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
)

// CryptModule 是 go-lua-vm 迁移后的模块壳。
type CryptModule struct{}

func New() *CryptModule { return &CryptModule{} }

func (m *CryptModule) Name() string { return "cryptLib" }

func (m *CryptModule) IsAvailable() bool { return true }

func (m *CryptModule) Register(engine model.Engine) error {
	engine.RegisterMethod("cryptLib.aes_crypt", "AES 加密/解密", func(data, key, operation, mode string, options ...interface{}) (string, error) {
		iv := ""
		if len(options) > 0 {
			optionIV, ok := options[0].(string)
			if !ok {
				return "", errors.New("iv 参数必须为字符串")
			}
			iv = optionIV
		}
		padding := true
		if len(options) > 1 {
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
		if _, err := rand.Read(key); err != nil {
			return "", err
		}
		return string(key), nil
	}, true)
	engine.RegisterMethod("cryptLib.aes_ivgen", "生成随机 IV", func() (string, error) {
		iv := make([]byte, aes.BlockSize)
		if _, err := rand.Read(iv); err != nil {
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
		hash := hmac.New(sha256.New, []byte(key))
		_, _ = hash.Write([]byte(data))
		return hex.EncodeToString(hash.Sum(nil))
	}, true)
	return nil
}

func GetModule() model.Module { return &CryptModule{} }

func aesCrypt(data, key, operation, mode, iv string, padding bool) (string, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("密钥长度必须为 16/24/32 字节")
	}
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	if mode != "ecb" && len(iv) != aes.BlockSize {
		return "", errors.New("非 ECB 模式需要 16 字节 iv 参数")
	}
	dataBytes := []byte(data)
	switch mode {
	case "ecb":
		return aesCryptECB(block, dataBytes, operation, padding)
	case "cbc", "cfb", "ofb", "ctr":
		stream := cipher.NewCTR(block, []byte(iv))
		result := make([]byte, len(dataBytes))
		stream.XORKeyStream(result, dataBytes)
		return string(result), nil
	default:
		return "", fmt.Errorf("不支持的加密模式: %s", mode)
	}
}

func aesCryptECB(block cipher.Block, data []byte, operation string, padding bool) (string, error) {
	if operation == "encrypt" {
		if padding {
			data = pkcs7Pad(data, block.BlockSize())
		}
		if len(data)%block.BlockSize() != 0 {
			return "", errors.New("明文长度必须为块大小的整数倍")
		}
		result := make([]byte, len(data))
		for index := 0; index < len(data); index += block.BlockSize() {
			block.Encrypt(result[index:index+block.BlockSize()], data[index:index+block.BlockSize()])
		}
		return string(result), nil
	}
	if len(data)%block.BlockSize() != 0 {
		return "", errors.New("密文长度必须为块大小的整数倍")
	}
	result := make([]byte, len(data))
	for index := 0; index < len(data); index += block.BlockSize() {
		block.Decrypt(result[index:index+block.BlockSize()], data[index:index+block.BlockSize()])
	}
	if padding {
		result = pkcs7Unpad(result)
	}
	return string(result), nil
}

func generateRSAKeyPair(keyBits ...int) (string, string, error) {
	bits := 2048
	if len(keyBits) > 0 && keyBits[0] > 0 {
		bits = keyBits[0]
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", err
	}
	publicPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
	})
	privatePEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	return string(publicPEM), string(privatePEM), nil
}

func pkcs7Pad(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := make([]byte, padding)
	for index := range padText {
		padText[index] = byte(padding)
	}
	return append(data, padText...)
}

func pkcs7Unpad(data []byte) []byte {
	if len(data) == 0 {
		return data
	}
	padding := int(data[len(data)-1])
	if padding < 1 || padding > len(data) {
		return data
	}
	for index := len(data) - padding; index < len(data); index++ {
		if int(data[index]) != padding {
			return data
		}
	}
	return data[:len(data)-padding]
}

func rsaEncrypt(data, key string, isPublicKey bool) (string, error) {
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return "", errors.New("无效的密钥格式")
	}
	var publicKey *rsa.PublicKey
	if isPublicKey {
		parsedKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
		if err != nil {
			return "", err
		}
		publicKey = parsedKey
	} else {
		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return "", err
		}
		publicKey = &privateKey.PublicKey
	}
	encrypted, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, []byte(data), nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func rsaDecrypt(data, key string, isPublicKey bool) (string, error) {
	if isPublicKey {
		return "", errors.New("使用公钥解密不支持")
	}
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return "", errors.New("无效的私钥格式")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		decodedData = []byte(data)
	}
	decrypted, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, decodedData, nil)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}
