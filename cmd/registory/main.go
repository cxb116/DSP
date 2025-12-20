package main

import (
	"fmt"
	EmployeePB "github.com/cxb116/DSP/src"
)

func main() {
	employee := EmployeePB.Employee{
		Id: 2,
	}

	fmt.Printf("registory", employee)
}
