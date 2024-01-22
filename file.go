package util

import "os"

const (
  FileStatNotExist = 0
  FileStatIsDir    = 1
  FileStatIsFile   = 2
)

func FileStat(filename string) int {
  info, err := os.Stat(filename)
  if os.IsNotExist(err) {
    return FileStatNotExist
  }
  if info.IsDir() {
    return FileStatIsDir
  } else {
    return FileStatIsFile
  }
}
