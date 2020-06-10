package main

import (
	"fmt"
	"math"
)

// フォーマット文字列については，
// https://oohira.github.io/gobyexample-jp/string-formatting.html
// を参照のこと．
// 基本的にはコンマ繋ぎで書けば良い．

func main() {
	fmt.Println("Now you have %g problems", math.Sqrt(7))
}
