package util

import (
	"errors"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"time"
)

const (
	fileStatNotExist = iota
	fileStatIsDir
	fileStatIsFile
	fileStatDenied
	fileStatError
)

type FileStatModel struct {
	stat int
}

func (f *FileStatModel) IsFile() bool {
	return f.stat == fileStatIsFile
}

func (f *FileStatModel) NotFile() bool {
	return f.stat != fileStatIsFile
}

func (f *FileStatModel) IsDir() bool {
	return f.stat == fileStatIsDir
}

func (f *FileStatModel) NotDir() bool {
	return f.stat != fileStatIsDir
}

func (f *FileStatModel) IsExist() bool {
	return f.stat != fileStatNotExist
}

func (f *FileStatModel) NotExist() bool {
	return f.stat == fileStatNotExist
}

func (f *FileStatModel) IsDenied() bool {
	return f.stat == fileStatDenied
}

func (f *FileStatModel) NotDenied() bool {
	return f.stat != fileStatDenied
}

func (f *FileStatModel) IsError() bool {
	return f.stat == fileStatError
}

func (f *FileStatModel) NotError() bool {
	return f.stat != fileStatError
}

func FileStat(filename string) *FileStatModel {
	info, err := os.Stat(filepath.Clean(filename))
	if errors.Is(err, fs.ErrNotExist) {
		return &FileStatModel{stat: fileStatNotExist}
	}
	if errors.Is(err, fs.ErrPermission) {
		return &FileStatModel{stat: fileStatDenied}
	}
	if err != nil {
		return &FileStatModel{stat: fileStatError}
	}
	if info.IsDir() {
		return &FileStatModel{stat: fileStatIsDir}
	} else {
		return &FileStatModel{stat: fileStatIsFile}
	}
}

type DirListModel struct {
	Files []DirListUnitModel
	Dirs  []DirListUnitModel
	Error error
}

type DirListUnitModel struct {
	Name    string
	Size    int64
	Mode    fs.FileMode
	ModTime time.Time
}

func NewDirListUnit(name string, size int64, mode fs.FileMode, modeTime time.Time) DirListUnitModel {
	return DirListUnitModel{
		Name:    name,
		Size:    size,
		Mode:    mode,
		ModTime: modeTime,
	}
}

func DirList(p string) *DirListModel {
	ls := &DirListModel{Error: nil}
	dir, err := os.Open(filepath.Clean(p))
	if err != nil {
		ls.Error = err
		return ls
	}
	defer func() {
		_ = dir.Close()
	}()
	dirs, err := dir.Readdir(-1)
	if err != nil {
		ls.Error = err
		return ls
	}
	for _, v := range dirs {
		if v.IsDir() {
			ls.Dirs = append(ls.Dirs, NewDirListUnit(v.Name(), v.Size(), v.Mode(), v.ModTime()))
		} else {
			ls.Files = append(ls.Files, NewDirListUnit(v.Name(), v.Size(), v.Mode(), v.ModTime()))
		}
	}
	return ls
}

func MkdirP(p string, perm ...os.FileMode) error {
	p = path.Clean(p)
	if FileStat(p).NotExist() {
		if len(perm) > 0 {
			return os.MkdirAll(p, perm[0])
		} else {
			return os.MkdirAll(p, 0750)
		}
	}
	return nil
}

func DirSize(p string) (int64, error) {
	var size int64
	p = path.Clean(p)
	err := filepath.Walk(p, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	return size, err
}
