package main

import "fmt"

var pl = fmt.Println

func main() {
	// intSet :=map[int]bool{}
	// vals :=[]int{5,10,2,8,7,9,1,2,10}
	// for _, v := range vals{
	// 	intSet[v] = true
	// }
	// pl(len(vals), len(intSet))
	// pl(intSet[5])
	// pl(intSet[500])
	// if intSet[100]{
	// 	pl("100 is in the set")
	// }


	// intSet := map[int]struct{}{}
	// vals := []int{5,10,2,8,7,9,1,2,10}
	// for _, v:= range vals {
	// 	intSet[v] = struct{}{}
	// }
	// if _, ok := intSet[5]; ok {
	// 	pl("5 is in the set")
	// }

	type firstName struct {
		name string
		age int
	}
	f:= firstName{
		name: "Eddy",
		age : 50,
	}
	var g struct{
		name string
		age int
	}
	g = f
	pl(f == g)

}
