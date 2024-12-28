package main

import (
	"errors"
	"fmt"
	"strconv"
)

var pl = fmt.Println

func add(i int, j int)(int, error){
	return i + j, nil
}

func sub(i int, j int)(int, error){
	return i + j, nil
}

func mul(i int, j int)(int, error){
	return i + j, nil
}

func div(i int, j int)(int, error){
	if j == 0{
		return 0, errors.New("division by zero")
	}
	return i / j, nil
}

var opMap=map[string]func(int, int)(int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main(){
	expressions:= [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
		{"10", "/", "0"},
	}
	for _, expression := range expressions{
		if len(expression) != 3{
			pl("Invalid expression:", expression)
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
			pl("unsupported operator:",op)
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil{
			pl(err)
			continue
		}
		result, err := opFunc(p1, p2)
			if err != nil{
				pl(err)
				continue
			}
			pl(result)
	}
}