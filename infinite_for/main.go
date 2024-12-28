package main

import "fmt"

var pl = fmt.Println

func main(){
	// for i := 1; i <= 100; i++{
	// 	if i%3 == 0{
	// 		if i%5 == 0{
	// 			pl("FuzzBuzz")
	// 		}else{
	// 			pl("Fizz")
	// 		}
	// 	}else if i%5 == 0{
	// 		pl("Buzz")
	// 	}else{
	// 		pl(i)
	// 	}
	// }

	for i := 1; i <= 100; i++{
		if i%3 == 0 && i%5 == 0{
			pl("FizzBuzz")
			continue
		}
		if i%3 == 0{
			pl("Fizz")
			continue
		}
		if i%5 == 0{
			pl("Buzz")
			continue
		}
		pl(i)
	}
	
}