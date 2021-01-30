package util

import (
	"crypto/sha256"
	"errors"
	"io"
	"log"
	"os"
)

func SHA256File(filePath string) ([32]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return [32]byte{}, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Print(err)
		}
	}()
	return SHA256FileIO(file)
}

func SHA256FileIO(file *os.File) ([32]byte, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return [32]byte{}, err
	}
	hashInBytes := hash.Sum(nil)
	if len(hashInBytes) != 32 {
		return [32]byte{}, errors.New("wrong length")
	}
	res := [32]byte{}
	copy(res[:], hashInBytes)
	return res, nil
}

func SHA256String(str string) [32]byte {
	data := sha256.Sum256([]byte(str))
	return data
}

func SHA256Bytes(str []byte) [32]byte {
	data := sha256.Sum256(str)
	return data
}
