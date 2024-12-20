package main

import (
	"fmt"
	"jvmgo/ch02/classpath"
	"strings"
)

func main() {
	//解析命令
	cmd := parseCmd()
	if cmd.versionFlag {
		//如果带了-version命令
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		//如果带了help命令。或者class为空
		printUsage()
	} else {
		//开启jvm
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	//拿到那些加载器后
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n", cp, cmd.class, cmd.args)
	//替换.为/  方便搜索文件
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Cloud not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
