package main

import (
	"fmt"
	"sync"
)

type employee struct {
	Name string
}

var lock = &sync.Mutex{}
var singleInstance *employee

func NewEmployee(name string) *employee {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			singleInstance = &employee{name}
			fmt.Println("Created an employee now: ", singleInstance.Name)
		} else {
			fmt.Println("An employee is already created: ", singleInstance.Name)
		}

	} else {
		fmt.Println("An employee is already created: ", singleInstance.Name)
	}

	return singleInstance
}

func (e *employee) ManageCloud() {
	fmt.Printf("%v manages the cloud\n", e.Name)
}
func (e *employee) ManageInfra() {
	fmt.Printf("%v manages the infra\n", e.Name)
}

func main() {
	emp := NewEmployee("Dave")
	emp.ManageInfra()
	emp1 := NewEmployee("Rak")
	emp1.ManageCloud()
}
