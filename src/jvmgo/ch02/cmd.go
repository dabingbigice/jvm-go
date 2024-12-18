package main

import (
	"flag"
	"fmt"
	"os"
)

// go build -o ch02
type Cmd struct {
	helpFlag    bool
	versionFlag bool
	XjreOption  string
	cpOption    string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	//java [-options] class [args] 解析命令
	cmd := &Cmd{}
	flag.Usage = printUsage
	//-help
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	//-?
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	//-version
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	//-classpath
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	//-cp
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	//-Xjre
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()
	//获取没有带-参数的所有参数
	args := flag.Args()
	if len(args) > 0 {
		//第一个参数为class
		cmd.class = args[0]
		//后面的都是参数
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	//解析help或者失败会调用这个
	fmt.Printf("Usage: %s [-option] class [args...]\n", os.Args[0])
}
