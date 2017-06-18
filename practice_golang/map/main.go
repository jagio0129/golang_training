package main

import (
	"fmt"
)

func main() {
	map02()

}

func map01() {
	//マップの作成
	capitals := make(map[string]string)

	//インデックスを使用してマップに値を書くの
	capitals["日本"] = "東京"
	capitals["アメリカ"] = "ワシントンD.C."
	capitals["中国"] = "北京"

	//インデックスを使用して値を取り出す
	fmt.Println(capitals["日本"])
	fmt.Println(capitals["アメリカ"])
	fmt.Println(capitals["中国"])

	//rangeでキーも取り出す
	for country, capital := range capitals {
		fmt.Println("国名："+country, "首都："+capital)
	}
}

func map02() {
	//マップの作成
	capitals := make(map[string]string)

	//インデックスを使用してマップに値を書くの
	capitals["日本"] = "東京"
	capitals["アメリカ"] = "ワシントンD.C."
	capitals["中国"] = "北京"

	// キーが存在するか
	capital, ok := capitals["イギリス"]
	if ok {
		fmt.Println("国名：" + capital)
	} else {
		fmt.Println("未登録")
	}
}
