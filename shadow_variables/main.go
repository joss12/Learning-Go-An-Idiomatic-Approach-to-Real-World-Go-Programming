package main


import "fmt"

var pl = fmt.Println

func main(){
	x := 10
	if x > 5{
		pl(x)
		 x :=5
		 pl(x)
	}
	pl(x)
}