package main

import (
	"fmt"
	"os"
)

func main() {
	write()
	read()
}

func write() {
	//ファイルオープン
	file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY, 0)

	//エラーチェック
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//ファイルクローズ(遅延実行)
	defer func() {
		file.Close()
	}()

	file.WriteString("test")
}

func read() {
	//ファイルオープン
	file, err := os.OpenFile("test.txt", os.O_RDONLY, 0)

	//エラーチェック
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//ファイルクローズ(遅延実行)
	defer func() {
		file.Close()
	}()

	//ファイルから文字列を読み込む
	var str string
	fmt.Fscanf(file, "%s", &str)

	//読み込んだ文字列の出力
	fmt.Println(str)

}
