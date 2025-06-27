package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

type Human struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

func (h Human) M() {
	fmt.Println(h.S)
}

func main() {
	var i I = T{"John"}
	var j I = Human{"Johnasan"}
	i.M()
	j.M()
}
