package main

import "fmt"
import "go-x-example/map_reduce/example"
//"go-basic-example/math_example"

func main() {

	// EmployeeConutIf 和 EmployeeSumIf 分别用于统满足某个条件的个数或总数。它们都是Filter + Reduce的语义。
	//EmployeeFilterIn 就是按某种条件过虑。就是Fitler的语义。


	old := obj.EmployeeCountIf(obj.List, func(e *obj.Employee) bool {
		return e.Age > 40
	})
	fmt.Printf("old people: %d\n", old)


	high_pay := obj.EmployeeCountIf(obj.List, func(e *obj.Employee) bool {
		return e.Salary >= 6000
	})
	fmt.Printf("High Salary people: %d\n", high_pay)
	//High Salary people: 4
}
