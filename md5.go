package util

import (
	"crypto/md5"
	"errors"
	"io"
	"log"
	"os"
)

func MD5File(filePath string) ([16]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return [16]byte{}, err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Print(err)
		}
	}()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return [16]byte{}, err
	}
	hashInBytes := hash.Sum(nil)
	if len(hashInBytes) != 16 {
		return [16]byte{}, errors.New("wrong length")
	}
	res := [16]byte{}
	copy(res[:], hashInBytes)
	return res, nil
}

func MD5String(str string) [16]byte {
	return md5.Sum([]byte(str))
}

func MD5Bytes(b []byte) [16]byte {
	return md5.Sum(b)
}
