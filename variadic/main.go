package main

import (
	"errors"
	"fmt"
	"os"
)

var pl = fmt.Println

// func addTo(base int, vals ...int)[]int {
// 	out:= make([]int, 0, len(vals))
// 	for _, v := range vals{
// 		out = append(out, base+v)
// 	}
// 	return out
// }



func divAndRemainder(num, denom int)(result int, remainder int, err error){
	if  denom == 0{
		err = errors.New("cannot divide by zero")
		return result, remainder, err
	}
	result, remainder = num/denom, num%denom
	return result, remainder, err
}

func main(){
	// pl(addTo(3))
	// pl(addTo(3, 2))
	// pl(addTo(3,2,4,6,8))
	// a :=[]int{4, 3}
	// pl(addTo(3, a...))
	// pl(addTo(3, []int{1, 2,3,4,5}...))
	// result := divAndRemainder(5,10)

	result, remainder, err := divAndRemainder(5, 2)
	if err != nil{
		pl(err)
		os.Exit(1)
	}
	pl(result, remainder)
}