package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	//单个类路径
	absDir string
}

// 实例化结构体,拼接加载路径
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		//发生异常了，直接panic,如果是nil代表没有发生异常
		panic(err)
	}
	return &DirEntry{absDir}
}

// 加载类，返回二进制流
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	//去指定文件下去加载某个类
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

// 返回当前对象存储的加载路径
func (self *DirEntry) String() string {

	return self.absDir
}
