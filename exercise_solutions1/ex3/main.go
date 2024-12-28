package main

import "fmt"

var pl = fmt.Println

func prefixer(prefix string)func (string) string{
	return func(body string) string{
		return prefix + " " + body
	}
}

func main(){
	helloPrefix := prefixer("Hello")
	pl(helloPrefix("Eddy"))
	pl(helloPrefix("Grace"))
}