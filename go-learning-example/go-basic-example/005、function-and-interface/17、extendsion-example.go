package main

import "fmt"

/* 扩展和 复用 */

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" " + host)
}

type Dog struct {
	p Pet
}

func main()  {
	dog := new(Dog)
	dog.p.SpeakTo("wang wang")
}

