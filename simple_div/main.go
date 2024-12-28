package main

import "fmt"

var pl = fmt.Println


type MyFuncOpts struct{
	FirstName string
	LastName string
	Age int
}

func MyFunc(opts, MyFuncOpts)error{

}

func main(){
	// var result int
	result := div(2, 3)
	pl(result)
	MyFunc(MyFuncOpts{
		LastName: "Eddy",
		Age: 50,
	})
	// MyFunc(MyFuncOpts{
	// 	FirstName: "Joss",
	// 	LastName: "Mouity",
	// })
}


func div(num int, demon int)int{
	if demon == 0{
		return 0
	}
	return num / demon
}


