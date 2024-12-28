package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
)

func main(){
	const data = `
	{"name": "Eddy", "age": 40}
	{"name": "Grace", "age": 9}
	{"name": "Stan", "age": 5}
	`

	var t struct{
		Name string `json:"name"`
		Age int `json:"age"`
	}

	dec := json.NewDecoder(strings.NewReader(data))
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	for{
		err := dec.Decode(&t)
		if err != nil{
			if errors.Is(err, io.EOF){
				break
			}
			panic(err)
		}
		fmt.Println(t)
		err = enc.Encode(t)
		if err != nil{
			panic(err)
		}
	}
	out := b.String()
	fmt.Println(out)
}
