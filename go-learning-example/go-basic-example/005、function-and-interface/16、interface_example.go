package main

import (
	"fmt"
)

type Programmer interface {

	PrintHelloWorld() string
}

type GoProgrammer struct {
}

func (p *GoProgrammer) PrintHelloWorld() string {
	return "fmt.Println(\"hello world\")"
}


func main()  {

	var p Programmer
	p = new(GoProgrammer)
	fmt.Println(p.PrintHelloWorld())

}