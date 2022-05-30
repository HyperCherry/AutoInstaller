package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ConfigInfo struct {
	Name     string
	Url      string
	Argument string
}

func ConfigReader() (in []ConfigInfo) {
	filePtr, err := os.Open("./info.json")
	if err != nil {
		fmt.Printf("配置文件info.json丟失 [Err:%s]\n", err.Error())
		return
	}
	defer filePtr.Close()
	var info []ConfigInfo
	decoder := json.NewDecoder(filePtr)
	err = decoder.Decode(&info)
	if err != nil {
		fmt.Println("解码失败", err.Error())
	}
	return info
}
