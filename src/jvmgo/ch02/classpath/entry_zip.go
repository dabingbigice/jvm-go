package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	//文件路径
	absDir string
}

func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		//发生异常了，直接panic
		panic(err)
	}
	return &ZipEntry{absDir}
}

// readClass（）方法负责寻找和加载class文件
func (self *ZipEntry) readClass(classname string) ([]byte, Entry, error) {
	//打开jar文件
	r, err := zip.OpenReader(self.absDir)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	//遍历jar文件中的所有文件。其中是否包含classname
	for _, f := range r.File {
		//判断文件名是不是需要的。不带扩展名
		//object类在rt.jar中
		if f.Name == classname {
			//打开当前找到文件，然后读入他作为二进制流。最后返回
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found:" + classname)
}

// String（）方法的作用相当于Java中的toString（），用于返回变量的字符串表示
func (self *ZipEntry) String() string {

	return self.absDir
}
