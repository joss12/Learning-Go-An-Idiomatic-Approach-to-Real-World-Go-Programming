package main

import "fmt"

var pl = fmt.Println

func makeMult(base int) func(int)int{
	return func(factor int )int {
		return base * factor
	}
}


func main(){
	twoBase := makeMult(2)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++{
		pl(twoBase(i), threeBase(i))
	}
}