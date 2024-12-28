package main

import (
	"fmt"
	"maps"
)

var pl = fmt.Println

func main() {
	// totalWins := map[string]int{}
	// totalWins["Orcas"] = 1
	// totalWins["lions"] = 2
	// pl(totalWins["Orcas"])
	// pl(totalWins["lions"])
	// totalWins["kittens"]++
	// pl(totalWins["kittens"])
	// totalWins["lions"] = 3
	// pl(totalWins["lions"])

	// m := map[string]int{
	// 	"hello":5,
	// 	"world":0,
	// }
	// v, ok :=m["hello"]
	// pl(v, ok)

	// v, ok = m["world"]
	// pl(v, ok)

	// v, ok = m["goodbey"]
	// pl(v, ok)

	// m :=map[string]int{
	// 	"hello" : 5,
	// 	"world" : 10,
	// }
	// pl(m, len(m))
	// clear(m)
	// pl(m, len(m))

	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	n := map[string]int{
		"world": 10,
		"hello": 5,
	}
	pl(maps.Equal(m, n))
}
