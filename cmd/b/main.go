package main

import (
	"fmt"
	"github.com/lishimeng/app-starter/buildscript"
)

func main() {
	err := buildscript.Generate("home",
		"lishimeng",
		"cmd/home/main.go", true)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
