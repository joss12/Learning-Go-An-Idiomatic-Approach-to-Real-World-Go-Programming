package main


import "fmt"

var pl = fmt.Println


type Adder struct{
	start int
}

func (a Adder)AddTo(val int) int{
	return a.start + val
}


func main(){
	myAdder := Adder{start: 10}
	pl(myAdder.AddTo(5))
	f1 := myAdder.AddTo
	pl(f1(10))
	f2 := Adder.AddTo
	pl(f2(myAdder, 15))
}