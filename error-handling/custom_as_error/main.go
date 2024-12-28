package main


import (
	"errors"
	"fmt"
)

type MyErr struct{
	Codes []int
}

func (me MyErr) CodeVals() []int{
	return me.Codes
}

func (me MyErr) Error() string{
	return fmt.Sprintf("Codes: %v", me.Codes)
}

func AFunctionThatReturnAnError() error{
	return MyErr{Codes: []int{1, 2, 3, 4, 5, 8}}
}

func main(){
	err := AFunctionThatReturnAnError()
	var myErr MyErr
	if errors.As(err, &myErr){
		fmt.Println(myErr.Codes)
	}

	var coder interface{
		CodeVals()[]int
	}
	if errors.As(err, &coder){
		fmt.Println(coder.CodeVals())
	}
}