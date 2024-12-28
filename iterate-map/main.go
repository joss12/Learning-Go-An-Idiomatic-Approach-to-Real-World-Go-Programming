package main

import "fmt"

var pl = fmt.Println

func main(){
	m := map[string]int{
		"a":1,
		"c":3,
		"d":2,
	}
	for i := 0; i < 3; i++{
		pl("Loop", i)
		for k, v := range m {
			pl(k, v)
		}
	}
}