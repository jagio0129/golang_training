package main

import (
	"fmt"
)

func main() {
	//run1()
	//run2()
	//run3()
	run4()
}

/*
 *基本的な使い方
 */
func run1() {
	//構造体型の変数を宣言
	var x struct {
		i1, i2 int
		f      float32
		s      string
	}

	//構造体の書くフィールドに値をセット
	x.i1 = 1
	x.i2 = 2
	x.f = 3.14
	x.s = "Go"

	//出力
	fmt.Println(x)
}

/*
 *埋め込み。クラスの継承と似たようなことを実現できる。
 */
//埋め込まれる側の構造体
type embedded struct {
	i int
}

//embedded型のメソッド
func (x embedded) doSomething() {
	fmt.Println("test.doSomething()")
}

//埋め込み先の構造体
type test struct {
	embedded
}

func run2() {
	var x test

	// embedded型に実装されているメソッドに、test型の値でアクセス
	x.doSomething()

	//embedded型のフィールドにtest型の値でアクセス
	fmt.Println(x.i)
}

/*
 *構造体の初期化
 */
//構造体の宣言
type Person struct {
	name string
	age  int
}

func run3() {
	//構造体リテラルを使用せず、フィールドを個別に初期化
	var p1 Person
	p1.name = "Jhon"
	p1.age = 21

	//構造体リテラルで初期化(フィールド名と値のペアを記述)
	p2 := Person{age: 31, name: "Tom"}

	//構造体リテラルで初期化(フィールドの宣言順に値を記述)
	p3 := Person{"Jane", 42}

	//ポインタも構造体リテラルで作成可能
	p4 := &Person{"Mike", 36}

	// 出力
	fmt.Println(p1, p2, p3, p4)
}

/*
 *構造体リテラルにおける埋込
 */
//構造体型の宣言
type Person2 struct {
	name string
	age  int
}

type Employee struct {
	id int
	Person2
}

func run4() {
	//構造体リテラうで初期化
	e := Employee{1, Person2{"Jack", 28}}

	//出力
	fmt.Println(e)
}
