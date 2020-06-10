package main

import (
	"fmt"
	"time"
)

/*
	Printlnはカンマ区切りで出力可能．
	現在時刻はtime.Now().
*/

func main() {
	fmt.Println("Welcome to the playground!!")
	fmt.Println("The time is ", time.Now())
	// Welcome to the playground!!
	// The time is  2020-06-10 22:51:34.454787 +0900 JST m=+0.000149066
}
