//package main

import "fmt"

func main() {
	var a [2]string // aという名前の長さ2の文字列型配列を定義
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{1, 2, 3, 4, 5, 6} // primesという長さ6のint型配列を定義し、初期値を与えて初期化
	fmt.Println(primes)
}
