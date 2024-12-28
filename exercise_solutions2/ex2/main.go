package main

import "fmt"

var pl = fmt.Println

func UpdateSlice(s []string, val string){
	s[len(s)-1] = val
	pl("in UpdateSlice")
}

func GlowSlice(s []string, val string){
	s = append(s, val)
	pl("in GrowSlice:", s)
}

func main(){
	s := [] string{"a", "b", "c"}
	UpdateSlice(s, "d")
	pl("in main after UpdateSlice:", s)
	GlowSlice(s, "e")
	pl("in main, after GrowSlice:", s)
}