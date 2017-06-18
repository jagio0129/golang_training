# マップ
- **キー**と**値**の組み合わせの集合
- `map [キーの型] 要素の型`

## マップの作成
- スライスと同様**参照型**の一種。
- プログラムからマップの参照先のデータに直接アクセスすることはできない。
- マップに限らず、参照型の値を作成するには`make`関数を使用する。
  - `make(map [キーの型] 要素型, キャパシティの初期値)`
  - `make(map [キーの型] 要素型)`

```go
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
```

## マップのキー存在確認
- マップ内にキーが格納されているかを調べるには、要素の値と一緒に論理値型の値を受け取る
```go
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
```