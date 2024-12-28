package main

import (
	"fmt"
)

var pl = fmt.Println

func main(){
	// x := 10
	// pointerToX := &x
	// pl(pointerToX)
	// pl(*pointerToX)
	// z := 5 **pointerToX
	// pl(z)

	// var x *int
	// pl(x == nil)
	// pl(*x)

	// var x = new(int)
	// pl( x == nil)
	// pl(*x)

	// people := person{
	// 	FirstName: "Eddy",
	// 	MiddleName: makePointer("Egs"), // line won't compile
	// 	LastName: "Grace",
	// }
	// pl(people)

	// x:= 10
	// failedUpdate(&x)
	// pl(x)
	// update(&x)
	// pl(x)

	// pl(MakeFoo())

	// file, err := os.Open(fileName)
	// if err != nil {
	// 	return 
	// }
	// defer file.Close()
	// data := make([]byte, 100)
	// for {
	// 	count, err := file.Read(data)
	// 	process(data[:count])
	// 	if err != nil{
	// 		if errors.Is(err, io.EOF){
	// 			return nil
	// 		}
	// 		return err
	// 	}
	// }


}





// func MakeFoo()(Foo, error){
// 	f := Foo{
// 		Field1: "val",
// 		Field2: 20,
// 	}
// 	return f, nil
// }

// f := struct {
// 	Name string `json:"name"`
// 	Age int `json:"age"`
// }{}
// err := json.Unmarshal([]byte(`{"name":"Eddy", "age":23}`), &f)

	
// func MakeFoo(f *Foo)error{
// 	f.Field1 = "val",
// 	f.Field2 = 20,
// 	return f, nil
// }

// func  failedUpdate(px *int){
// 	x2 := 20
// 	px = &x2
// }

// func update(px *int){
// 	*px = 20
// }

// func makePointer[T any](t T) *T{
// 	return &t
// }

// type person struct{
// 	FirstName string
// 	MiddleName *string
// 	LastName string
// }
