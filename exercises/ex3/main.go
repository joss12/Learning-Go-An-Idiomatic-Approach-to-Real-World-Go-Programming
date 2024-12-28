package main

import "fmt"

var pl = fmt.Println

func main(){
	var total int
	for i := 0; i <10; i++{
		total := total + 1
		pl(total)
	}
	pl(total)
}