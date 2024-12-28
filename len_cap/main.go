package main

import "fmt"

var pl = fmt.Println

func main (){
	s:=[]string{"a", "b", "c"}
	pl(s, len(s))
	clear(s)
	pl(s, len(s))
}

func clear(s []string) {
	panic("unimplemented")
}


