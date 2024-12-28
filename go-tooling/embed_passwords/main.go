package main

import(
	_"embed"
	"fmt"
	"io/fs"
	"os"
	"strings"
)


//go:embed password
var passwords  string

func mai(){
	pwds := strings.Split(passwords, "\n")
	if len(os.Args) > 1{
		for _, v := range pwds{
			if v == os.Args[1]{
				fmt.Println("true")
				os.Exit(0)
			}
		}
		fmt.Println("false")
	}
}