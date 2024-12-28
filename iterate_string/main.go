package main

import "fmt"

var pl = fmt.Println

func main() {
	samples := []string{"hello", "apple_π!"}
	for _, sample := range samples {
		for i, r := range sample {
			pl(i, r, string(r))
		}
		pl()
	}
}
