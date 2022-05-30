package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var Path string = "res"

func main() {
	checkDir()
	ConfigInfo := ConfigReader()
	for _, info := range ConfigInfo {
		fmt.Println("开始安装:" + info.Name)
		Download(info.Url, Path)
		split := strings.Split(info.Argument, " ")
		command := exec.Command(Path+"/"+getProgramName(info.Url), split...)
		command.CombinedOutput()
		fmt.Println(info.Name + "安装完成")
	}
	fmt.Printf("Install Complete..Press any key to exit...")
	b := make([]byte, 1)
	os.Stdin.Read(b)
}
func getProgramName(url string) (name string) {
	split := strings.Split(url, "/")
	return split[len(split)-1]
}
func checkDir() {
	_, err := os.Stat(Path)
	if os.IsNotExist(err) {
		fmt.Println("文件夹不存在，已创建")
		os.MkdirAll(Path, os.ModePerm)
	} else {
		fmt.Println("文件夹存在，跳过创建")
	}
}
