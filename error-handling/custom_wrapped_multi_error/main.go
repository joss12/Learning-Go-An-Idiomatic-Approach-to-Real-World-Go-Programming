package main

import(
	"errors"
	"fmt"
)

type Status int

type StatusErr struct{
	Status Status
	Message string
}

func (se StatusErr)Error() string{
	return se.Message
}

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type MyErr struct{
	Code int
	Errors []error
}

func (m MyErr)Error() string{
	return errors.Join(m.Errors...).Error()
}

func (m MyErr)UnWrap() []error{
	return m.Errors
}

func funcThatReturnAnError() error{
	return MyErr{
		Code: 12,
		Errors: []error{
			StatusErr{
				Status: NotFound,
				Message: "file Not Found",
			},
			errors.New("a simple string error"),
		},
	}
}

func main(){
	var err error
	err = funcThatReturnAnError()
	switch err := err.(type){
	case interface{UnWrap() error}:
		// handle single error
		innerErr := err.UnWrap()
		// process innerErr
		fmt.Println(innerErr)
	case interface{UnWrap() []error}:
		// handle multiple errors
		innerErrs := err.UnWrap()
		for _, innerErr := range innerErrs{
			fmt.Println(innerErr)
		}
	default:
		// handle no wrapped error
	}
}