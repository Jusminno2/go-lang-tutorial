package main

import "fmt"

func main() {
	i, j := 42, 2701 // :=はGoの短縮代入構文｜Pythonなら i = 42; j = 2701 に相当

	p := &i         // pはiのアドレス（メモリ位置）を指すポインタ｜&iは「変数iのアドレスを取得する」という意味
	fmt.Println(*p) // pが指すアドレスにある値（＝iの値）
	*p = 21         // *p = 21 は「pが指している変数（i）に21を代入せよ」という意味
	fmt.Println(i)  // *p -> i -> 42 を *p -> i -> 21

	p = &j       // pの住所変更 i -> j
	*p = *p / 37 // jの番号が変更： 2701 / 37 = 73
	fmt.Println(j)
}
