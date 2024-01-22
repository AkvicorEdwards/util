package util

import (
  "fmt"
  "testing"
)

func TestFileStat(t *testing.T) {
  // dd if=/dev/random of=testfile/testFile1 bs=4KB count=4
  if FileStat("testfile") != FileStatIsDir {
    t.Errorf("FileStatIsDir")
  }
  if FileStat("testfile/testFile1") != FileStatIsFile {
    t.Errorf("FileStatIsFile")
  }
  if FileStat("testfile/testFile2") != FileStatNotExist {
    t.Errorf("FileStatNotExist")
  }
}

func ExampleFileStat() {
  if FileStat("testfile") == FileStatIsDir {
    fmt.Println("FileStatIsDir")
  }
  if FileStat("testfile/testFile1") == FileStatIsFile {
    fmt.Println("FileStatIsFile")
  }
  if FileStat("testfile/testFile2") == FileStatNotExist {
    fmt.Println("FileStatNotExist")
  }
  
  // Output:
  // FileStatIsDir
  // FileStatIsFile
  // FileStatNotExist
}
