package main

import "fmt"


func main(){

	in1 := make(chan int)
	in2 := make(chan int)
	go func(){
		for i := 10; i <100; i  += 10{
			in1 <- i
		}
		close(in2)
	}()
	result := readFromTwochannels(in1, in2)
	fmt.Println(result)
}

func readFromTwochannels(in, in2 chan int)[]int{
	var out []int

	for count := 0; count < 2;{
		select{
		case v, ok := <-in:
			if !ok{
				in = nil // the case will never succeed again!
				count++
				continue
			}
			//process the v that was read from in
			out = append(out, v)
		case v, ok := <-in2:
		if !ok{
			in2 = nil // th case will never succeed again!
			count++
			continue
		}
		//process the v that was read from in2
		out = append(out, v)

		}
	}
	return out
}