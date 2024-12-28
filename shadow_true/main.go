package main

import (
	"fmt"

	// "golang.org/x/exp/rand"
)

var pl = fmt.Println

func main(){
	// pl(true)
	// true := 10
	// pl(true)

	// if n := rand.Intn(10); n == 0{
	// 	pl("That's too low")
	// }else if n > 5{
	// 	pl("That's too big", n)
	// }else{
	// 	pl("That's a good number", n)
	// }
	// pl(n) undefined: n

	// i := 0
	// for; i < 10; i++{
	// 	pl(i)
	// }

	// for i := 0; i < 10;{
	// 	pl(i)
	// 	if i % 2 == 0{
	// 		i++
	// 	}else{
	// 		i+=2
	// 	}
	// }
	i := 1
	for i < 100{
		pl(i)
		i = i * 2
	}
}