package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID string
}

func (e Employee) PrintInfo() string {
	return fmt.Sprintf("员工ID：%s，姓名：%s，年龄：%d", e.EmployeeID, e.Name, e.Age)
}
func main() {
	employee := Employee{Person{"张三", 20}, "1"}
	fmt.Print(employee.PrintInfo())
}
