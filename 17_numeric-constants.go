package main

import "fmt"

/*
  数値の定数は、高精度な 値 ( values )です。
  型のない定数は、その状況によって必要な型を取ることになります。
  例で needInt(Big) を出力してみてください。
	( int は64-bitの整数を保持できますが、それでは足りないことが時々あります。そういったときにconstを活用しましょう)
	Constは無限長精度の値を保持できる．
*/

const (
	Big   = 1 << 256
	Small = Big >> 255
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
