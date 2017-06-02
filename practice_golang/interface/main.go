package main

import (
	"fmt"
)

/*
 * インターフェースの基本的な使い方
 */
//演算インターフェース型
type Caluculator interface {
	//関数の定義
	Caluculate(a int, b int) int
}

//足し算型
type Add struct {
	//フィールドは持たない
}

//Add型にCaluculatorインターフェースのCaluculate関数を実装
func (x Add) Caluculate(a int, b int) int {
	//足し算
	return a + b
}

//引き算
type Sub struct {
	//フィールドは持たない
}

//Sub型にCaluculatorインターフェースのCaluculate関数を実装
func (x Sub) Caluculate(a, b int) int {
	//引き算
	return a - b
}

func main() {
	//Caluculatorインターフェースを実装した型の変数を宣言
	var add Add //足し算
	var sub Sub //引き算

	//Caluculatorインターフェース型の変数を宣言
	var cal Caluculator

	//Add型の値を代入
	cal = add

	//インターフェース経由でメソッドを呼び出す
	fmt.Println("和：", cal.Caluculate(1, 2))

	//Sub型の値を代入
	cal = sub

	fmt.Println("差：", cal.Caluculate(1, 2))
}

/*
 *インターフェースの埋め込み
 */
type Encoder interface {
	Encode(string) string
}

type Decoder interface {
	Decode(string) string
}
