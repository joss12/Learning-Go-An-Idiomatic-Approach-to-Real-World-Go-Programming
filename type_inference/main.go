package main

import (
	"fmt"
	"reflect"
)

type Integer interface{
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

// func Convert[T1, T2 Integer](int T1)T2{
// 	return T2(in)
// }

func PlusOneHundred[T Integer](in T) T {
	return in + 100
	}

func main(){
	var a int = 10
	b := Convert[int, int64](a)
	fmt.Println(b)
}