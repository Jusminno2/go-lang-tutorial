package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/*
・(v Vertex) -> Vertex型の値を受け取るレシーバ
・Abs() -> method名
*/
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
