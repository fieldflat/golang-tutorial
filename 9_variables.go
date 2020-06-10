package main

import "fmt"

/*
  var ステートメントは変数( variable )を宣言します。 関数の引数リストと同様に、複数の変数の最後に型を書くことで、変数のリストを宣言できます。
  var ステートメントはパッケージ、または、関数で利用できます。例のコードで示します。
*/

var c, python, java bool // 初期値はfalse

func main() {
	var i int                       // 初期値は0
	fmt.Println(i, c, python, java) // 0 false false false
}
