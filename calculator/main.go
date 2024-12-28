package main

import (
	"fmt"
	"strconv"
)

var pl = fmt.Println


func add(i int, j int) int{
	return i + j
}

func sub(i int, j int)int{
	return i - j
}

func mul(i int, j int)int{
	return i * j
}
func div(i int, j int)int{
	return i / j
}

var opMap = map[string]func(int, int)int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main(){
	expressions :=[][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
	}
	for _, expression := range expressions{
		if len(expressions) != 3{
			pl("invalid expression:", expressions)
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil{
			pl(err)
			continue
		}
		op := expression[1]
		opFunc, ok :=opMap[op]
		if !ok{
		pl("unsupported operator:", op)
		continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil{
			pl(err)
			continue
		}
		result := opFunc(p1, p2)
		pl(result)
	}
}