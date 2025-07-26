package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	person    Person
	EmployeID int
}

func (e Employee) PrintInfo() {
	fmt.Printf("Name: %s, Age: %d, EmployeeID: %d\n",
		e.person.Name, e.person.Age, e.EmployeID)
}

func main() {
	employee := Employee{Person{"AAA", 30}, 12345}

	employee.PrintInfo()
}
