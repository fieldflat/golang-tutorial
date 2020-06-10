package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is ", rand.Intn(10)) // シードを与えていないので，常に1が出力される．
}
