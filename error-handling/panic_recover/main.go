package main

import ("fmt")

func div06(i int){
	defer func(){
		if v := recover(); v != nil{
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}

func main(){
	for _, val := range []int{1, 2, 0, 6}{
		div06(val)
	}
}