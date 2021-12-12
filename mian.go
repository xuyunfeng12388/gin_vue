package main

import (
	"github.com/xuyunfeng12388/gin_vue/cmd"
	"os"
)

func main(){
	if err := cmd.Execute(); err != nil {
		println("start fail: ", err.Error())
		os.Exit(-1)
	}
}
