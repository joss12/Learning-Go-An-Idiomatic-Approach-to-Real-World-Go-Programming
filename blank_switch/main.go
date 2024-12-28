package main

import "fmt"

var pl = fmt.Println

func main(){
	// words := []string{"hi", "salutations", "hello"}
	// for _, word := range words{
	// 	switch wordLen := len(word);{
	// 	case wordLen < 5:
	// 		pl(word, "is a short word!")
	// 	case wordLen > 10:
	// 		pl(word, "is a long word")
	// 	default:
	// 		pl(word, "is exactly the right length.")
	// 	}
	// }

	for i := 1; i <= 100; i++{
		switch{
		case i%3 == 0 && i%5 == 0:
			pl("FuzzBuzz")
		case i%3 == 0:
			pl("Fuzz")
		case i%5 == 0:
			pl("Buzz")
		default:
			pl(i)
		}
	}
}