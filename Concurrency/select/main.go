package main

import "fmt"


func main(){
	ch1 := make(chan int)
	ch2 := make(chan int)

	go  func(){
		inGroroutine :=1
		ch1 <- inGroroutine
		fromMain := <- ch2
		fmt.Println("goroutine:", inGroroutine, fromMain)
	}()
	inMain := 2
	var fromGoroutine int
	select {
	case ch2 <- inMain:
	case fromGoroutine = <-ch1:
	}
	fmt.Println("main:", inMain, fromGoroutine)
}