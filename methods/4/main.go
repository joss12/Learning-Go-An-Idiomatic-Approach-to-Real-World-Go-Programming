package main


import "fmt"

var pl = fmt.Println

// type BitField int

// const(
// 	Field1 = 0
// 	Field2 = 1 + iota
// 	Field3 = 20
// 	Field4
// 	Field5 = iota
// )

// type Employee struct{
// 	Name string
// 	ID 	string
// }

// func (e Employee)Description() string{
// 	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
// }

// type Manager struct{
// 	Employee
// 	Reports []Employee
// }

// func (m Manager) FindNewEmployee() []Employee{
	// newEmp := make([]Employee, 0)
	// for _, emp := range m.Reports{
	// 	if emp.ID == "12345"{
	// 		newEmp = append(newEmp, emp)
	// 	}
	// }
	// return newEmp
	// return nil
// }


// type Inner struct{
// 	x int
// }

// type Outer struct{
// 	Inner
// 	x int
// }



func main(){
	

	// o := Outer{
	// 	Inner: Inner{
	// 		x: 10,
	// 	},
	// 	x: 20,
	// }
	// pl(o.x)
	// pl(o.Inner.x)
	

	// m := Manager{
	// 	Employee: Employee{
	// 		Name: "Eddy Mouity",
	// 		ID: "123456",
	// 	},
	// 	Reports: []Employee{},
	// }
	// pl(m.ID)
	// pl(m.Description())

	// pl(Field1, Field2, Field3, Field4, Field5)
}