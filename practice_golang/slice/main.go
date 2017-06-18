package main

import (
	"fmt"
)

func main() {
	slice03()
}

// 各要素を2倍に
func double(values []int) {
	for i := 0; i < len(values); i++ {
		values[i] *= 2
	}
}

func slice01() {
	x := [...]int{0, 1, 2, 3, 4, 5}
	var s []int

	//配列をスライスにして関数へ渡す
	s = x[:]
	double(s)

	fmt.Println(s)

}

func slice02() {
	array := [...]int{0, 1, 2, 3, 4, 5}

	// 配列をスライス
	s1 := array[1:3]
	fmt.Println(s1)
	fmt.Println("len:", len(s1))
	fmt.Println("cap:", cap(s1))

	//スライスをスライス
	s2 := s1[1:4]
	fmt.Println(s2)
	fmt.Println("len:", len(s2))
	fmt.Println("cap:", cap(s2))

	s3 := s1[1:3:4]
	fmt.Println(s3)
	fmt.Println("len:", len(s3))
	fmt.Println("cap:", cap(s3))
}

func slice03() {
	//文字列のスライス
	var x string = "abcdefg"[1:4]

	//ひらがなはUTF-8で3バイトなので「いう」が取り出せる
	var y string = "あいうえお"[3:9]

	//UTF-8の文字の境界として不適切な値を指定すると文字化けする
	var z string = "あいうえお"[1:4]

	//出力
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
