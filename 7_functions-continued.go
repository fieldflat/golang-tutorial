package main

import "fmt"

/*
	関数の２つ以上の引数が同じ型である場合には、最後の型を残して省略して記述できます。
*/

// 型をまとめて定義することができる．
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
