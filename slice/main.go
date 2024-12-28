package main

import (
	"fmt"
	"slices"
)

var pl = fmt.Println

func main(){
	x := []int{1, 2, 3, 4, 5}
	y := []int{1, 2, 3, 4, 5}
	z := []int{1, 2, 3, 4, 5, 6}
	s := []string{"a", "b", "c"}
	pl(slices.Equal(x, y))
	pl(slices.Equal(x, z))
	pl(slices.Equal(x, s))
}