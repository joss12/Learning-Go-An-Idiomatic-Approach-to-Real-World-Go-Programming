package main

import "fmt"

var pl = fmt.Println

func main() {
	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for _, v := range evenVals {
	// 	pl(v)
	// }

	// uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Grace": true}
	// for k := range uniqueNames{
	// 	pl(k)
	// }

	evenVals  := []int{2,4,6,8,10,12}
	for _, v := range evenVals{
		v *= 2
	}
	pl(evenVals)
}
