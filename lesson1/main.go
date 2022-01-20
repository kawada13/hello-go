package main

import (
	"fmt"
)

func main() {
	var i int = 100
	fmt.Println(i)

	var s string = "konnnitiha"
	fmt.Println(s)

	var t, f bool = true, false
	fmt.Println(t, f)

	var (
		i2 int    = 200
		s2 string = "hyde"
	)

	fmt.Println(i2, s2)

	// 暗黙的な定義
	// 最初に定義した型で決まる
	// 関数内でしか定義できない⇨明示的な定義は関数外で定義できる
	// 明示的な定義と暗黙的な定義の違いは、関数内か関数外で定義するかで決める
	// 一般的には明示的な定義を使うようにする
	// 変数は定義したら必ずその変数を使わないとエラーが出る
	i4 := 450
	fmt.Println(i4)
}
