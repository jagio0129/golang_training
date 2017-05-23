package main

import (
	"fmt"
	"strings"
)

func main() {
	splitStrings()
}

func stringsComaprison() {
	s1 := "alpha"
	s2 := "beta"

	//文字列を比較する
	if s1 == s2 {
		fmt.Println("同じ")
	}

	if s1 != s2 {
		fmt.Println("異なる")
	}

	if s1 < s2 {
		fmt.Println("s1のほうが小さい")
	}

	if s2 < s1 {
		fmt.Println("s2のほうが小さい")
	}

	if s1 <= s2 {
		fmt.Println("s1のほうが小さいか同じ")
	}

	if s2 <= s1 {
		fmt.Println("s2のほうが小さいか同じ")
	}
}

func stringsIntegrate() {
	//演算子による結合
	s := "alpha"
	s = s + "beta"
	s += "gamma"

	//結合結果の出力
	fmt.Println(s)

	//文字列スライスの用意
	arr := []string{"alpha", "beta", "gamma"}

	//文字列スライスの結合(カンマで結合)
	fmt.Println(strings.Join(arr, ","))
}

func getStringsPart() {
	//文字列の準備
	s := "abcdefghijklmn"

	//スライス式をつかて部分文字列を取り戻す
	fmt.Println(s[:10])
	fmt.Println(s[3:10])
	fmt.Println(s[3:])
	fmt.Println(s[:])
}

func stringsTrim() {
	//文字列の準備
	s := "xyz"

	//前後トリム
	fmt.Printf("[%s]\n", strings.TrimSpace(s))

	//左右どちらかだけトリムする
	fmt.Printf("[%s]\n", strings.TrimLeft(s, " "))
	fmt.Printf("[%s]\n", strings.TrimRight(s, " "))
}

func changeStrings() {
	//文字列の準備
	s := "This is a pen."

	//大文字に変換
	fmt.Println(strings.ToUpper(s))

	//小文字に変換
	fmt.Println(strings.ToLower(s))
}

func splitStrings() {
	//分割対象の文字列
	s := "one,two,three,four,five"

	//カンマで分割
	result := strings.Split(s, ",")

	fmt.Printf("%q\n", result)
}
