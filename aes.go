package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"strings"
)

type AESResult struct {
	data      []byte
	encrypted bool
	err       error
}

func (a *AESResult) Bytes() []byte {
	if a.err != nil {
		return []byte{}
	}
	return a.data
}

func (a *AESResult) Base64() *Base64Result {
	if a.err != nil {
		return NewBase64().EncodeBytes([]byte{})
	}
	return NewBase64().EncodeBytes(a.data)
}

func (a *AESResult) String() string {
	if a.err != nil {
		return ""
	}
	return string(a.data)
}

func (a *AESResult) Upper() string {
	if a.err != nil {
		return ""
	}
	return strings.ToUpper(hex.EncodeToString(a.data))
}

func (a *AESResult) Lower() string {
	if a.err != nil {
		return ""
	}
	return strings.ToLower(hex.EncodeToString(a.data))
}

func (a *AESResult) Decrypt(key []byte, iv ...[]byte) *AESResult {
	if a.err != nil {
		return a
	}
	if a.encrypted {
		return NewAES().DecryptCBC(a.data, key, iv...)
	} else {
		return a
	}
}

func (a *AESResult) Encrypt(key []byte, iv ...[]byte) *AESResult {
	if a.err != nil {
		return a
	}
	if a.encrypted {
		return a
	} else {
		return NewAES().DecryptCBC(a.data, key, iv...)
	}
}

func (a *AESResult) Encrypted() bool {
	return a.encrypted
}

func (a *AESResult) Error() error {
	return a.err
}

func NewAESResult(data []byte, encrypted bool, err error) *AESResult {
	res := &AESResult{
		data:      make([]byte, len(data)),
		encrypted: encrypted,
		err:       err,
	}
	copy(res.data, data)
	return res
}

type AES struct{}

func NewAES() *AES {
	return &AES{}
}

func (*AES) EncryptCBC(origData, key []byte, iv ...[]byte) *AESResult {
	block, err := aes.NewCipher(key)
	if err != nil {
		return NewAESResult(nil, false, err)
	}

	blockSize := block.BlockSize()
	var ivi []byte
	ivb := BytesCombine(iv...)
	if len(ivb) >= blockSize {
		ivi = ivb[:blockSize]
	} else {
		ivi = key[:blockSize]
	}
	origData = pkcs5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, ivi)
	encrypted := make([]byte, len(origData))
	blockMode.CryptBlocks(encrypted, origData)
	return NewAESResult(encrypted, true, nil)
}

func (*AES) DecryptCBC(encrypted, key []byte, iv ...[]byte) *AESResult {
	block, err := aes.NewCipher(key)
	if err != nil {
		return NewAESResult(nil, false, err)
	}

	blockSize := block.BlockSize()
	if len(encrypted)%blockSize != 0 {
		return NewAESResult(nil, false, errors.New("encrypted data size error"))
	}
	var ivi []byte
	ivb := BytesCombine(iv...)
	if len(ivb) >= blockSize {
		ivi = ivb[:blockSize]
	} else {
		ivi = key[:blockSize]
	}
	blockMode := cipher.NewCBCDecrypter(block, ivi)
	origData := make([]byte, len(encrypted))
	blockMode.CryptBlocks(origData, encrypted)
	origData = pkcs5UnPadding(origData)
	return NewAESResult(origData, false, nil)
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	if length == 0 {
		return origData
	}
	unPadding := int(origData[length-1])
	offset := length - unPadding
	if offset < 0 {
		offset = 0
	}
	return origData[:offset]
}
