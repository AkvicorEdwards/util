package util

import (
	"hash/crc32"
	"io"
	"os"
)

func CRC32File(path string) (uint32, error) {
	val := uint32(0)
	fs, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer func() {_=fs.Close()}()

	buf := make([]byte, 1024*10)
	var n int
	for {
		n, err = fs.Read(buf)
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
		val = crc32.Update(val, crc32.IEEETable, buf[:n])
	}
	return val, nil
}

func CRC32String(data string) uint32 {
	return crc32.ChecksumIEEE([]byte(data))
}

func CRC32Bytes(data []byte) uint32 {
	return crc32.ChecksumIEEE(data)
}
