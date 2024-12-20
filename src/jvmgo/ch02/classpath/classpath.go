package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOptin string) *Classpath {
	cp := &Classpath{}
	//boot和ext加载器的入口有很多jar包
	cp.parseBootAndExtClasspath(jreOption)
	//user加载器就是当前文件夹
	cp.parseUserClasspath(cpOptin)
	return cp
}
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className += ".class"
	//判断当前类是否由boot引导器加载
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		//先使用boot引导器加载
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		//再使用ext引导器加载
		return data, entry, err
	}
	//这两个都没加载成功，此时才在当前文件夹下搜索加载
	return self.userClasspath.readClass(className)
}
func (self *Classpath) String() string {
	return self.userClasspath.String()
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	//jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exist(jreOption) {
		return jreOption
	}
	if exist("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
