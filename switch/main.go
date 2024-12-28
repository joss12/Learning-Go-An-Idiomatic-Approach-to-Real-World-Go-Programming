package main

import "fmt"

var pl = fmt.Println

func main(){
	// words :=[]string{"a", "cow", "smile", "gopher", "octopus", "anthropologist"}
	// for _, word := range words{
	// 	switch size := len(word); size{
	// 	case 1, 2, 3, 4:
	// 		pl(word, "is a short word")
	// 	case 5:
	// 		wordLen := len(word)
	// 		pl(word, "is exactly the right length:", wordLen)
	// 	case 6, 7, 8, 9:
	// 	default:
	// 		pl(word, "is a long word")
	// 	}
	// }

loop:
	for i := 0; i < 10; i++{
		switch i {
		case 0, 2, 4, 6:
			pl(i, "is even")
		case 3:
			pl(i, "is divisible by 3 but not 2")
		case 7:
			pl("exit the loop")
			break loop
		default:
			pl(i, "is boring")
		}
	}
	
}