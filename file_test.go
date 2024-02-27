package util

import (
	"fmt"
	"testing"
)

func TestFileStat(t *testing.T) {
	// dd if=/dev/zero of=testfile/testFile1 bs=4KB count=4
	if FileStat("testfile").NotDir() {
		t.Errorf("testfile not dir")
	}
	if FileStat("testfile/testFile1").NotFile() {
		t.Errorf("testfile/testFile1 not file")
	}
	if FileStat("testfile/testFile2").IsExist() {
		t.Errorf("testfile/testFile2 is exist")
	}
}

func ExampleFileStat() {
	if FileStat("testfile").IsDir() {
		fmt.Println("testfile is dir")
	}
	if FileStat("testfile/testFile1").IsFile() {
		fmt.Println("testfile/testFile1 is file")
	}
	if FileStat("testfile/testFile2").NotExist() {
		fmt.Println("testfile/testFile2 not exist")
	}

	// Output:
	// testfile is dir
	// testfile/testFile1 is file
	// testfile/testFile2 not exist
}

func TestDirList(t *testing.T) {
	// dd if=/dev/zero of=testfile/testFile1 bs=4KB count=4
	ls := DirList("testfile")
	if ls.Error != nil {
		t.Errorf("DirList with error %v", ls.Error)
	}
	if len(ls.Files) != 1 {
		t.Errorf("failed get files")
	}
	if len(ls.Dirs) != 0 {
		t.Errorf("failed get dirs")
	}
	// for _, v := range ls.Dirs {
	// 	fmt.Printf("%s %s %s %s\n", v.Mode, v.ModTime.Format(time.DateTime), Size(v.Size).HighestUnit(true).Format(7, "", false), v.Name)
	// }
	// for _, v := range ls.Files {
	// 	fmt.Printf("%s %s %s %s\n", v.Mode, v.ModTime.Format(time.DateTime), Size(v.Size).HighestUnit(true).Format(7, "", false), v.Name)
	// }
}

func ExampleMkdirP() {
	err := MkdirP("testfile/testfile")
	if err != nil {
		fmt.Println(err)
	}
}

func ExampleDirSize() {
	p, err := DirSize("./testfile")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
}
