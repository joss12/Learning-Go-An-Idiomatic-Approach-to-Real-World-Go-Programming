package main

import (
	"fmt"
	"log"
	"os"
)

var pl = fmt.Println

func getFile(name string)(*os.File, func(), error){
	file, err := os.Open(name)
	if err != nil{
		return nil, nil, err
	}
	return file, func(){
		file.Close()
	}, nil
}

func main(){
	f, closer, err := getFile(os.Args[1])
	if err != nil{
		log.Fatal(err)
	}
	defer closer()
	pl(f)
}