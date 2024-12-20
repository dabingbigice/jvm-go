package classpath

import (
	"os"
	"strings"
)

const pathListSeparator string = string(os.PathSeparator)

type Entry interface {
	//readClass（）方法负责寻找和加载class文件
	readClass(classname string) ([]byte, Entry, error)
	//String（）方法的作用相当于Java中的toString（），用于返回变量的字符串表示
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		//普通path
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		//带*的path
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		//压缩包的path
		return newZipEntry(path)
	}
	//如果都不是上面的，则加载当前文件夹下的
	return newDirEntry(path)
}
