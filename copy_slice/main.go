package main

import "fmt"

var pl = fmt.Println

func main() {
	// x := []int {1,2,3,4}
	// y := make([]int, 4)
	// num := copy(y, x)
	// pl(y, num)

	// x := []int{1,2,3,4}
	// y := make([]int, 2)
	// num := copy(y, x)
	// pl(y, num)

	// x:= []int {1,2,3,4}
	// d:= [4]int{5,6,7,8}
	// y := make([]int, 2)
	// copy(y, d[:])
	// pl(y)
	// copy(d[:], x)
	// pl(d)

	totalWins := map[string]int{}
	totalWins["Orcas"] = 1
	totalWins["Lions"] = 2
	fmt.Println(totalWins["Orcas"])
	fmt.Println(totalWins["Kittens"])
	totalWins["Kittens"]++

	pl(totalWins["Kittens"])
	totalWins["Lions"] = 3
	pl(totalWins["Lions"])
}
