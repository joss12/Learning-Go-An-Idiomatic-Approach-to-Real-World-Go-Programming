package main

import "fmt"

var pl = fmt.Println

func main() {
// 	samples := []string{"hello", "apple_π!"}
// outer:
// 	for _, sample := range samples {
// 		for i, r := range sample {
// 			pl(i, r, string(r))
// 			if r == 'l' {
// 				continue outer
// 			}
// 		}
// 		pl()
// 	}

	// evenVals := []int{2,4,6,8,10}
	// for i, v := range evenVals{
	// 	if i == 0{
	// 		continue
	// 	}
	// 	if i == len(evenVals)-1{
	// 		break
	// 	}
	// 	pl(i, v)
	// }

	evenVals := []int{2,4,6,8,10}
	for i := 1; i <len(evenVals)-1; i++{
		pl(i, evenVals[i])
	}
}
