package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	write()
	read()
}

func write() {
	//file open
	file, err := os.OpenFile("test.csv", os.O_CREATE|os.O_WRONLY, 0666)

	// error check
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// file close(defer)
	defer func() {
		file.Close()
	}()

	//csv.Writerの作成
	w := csv.NewWriter(file)

	//データ出力
	w.Write([]string{"No", "商品", "価格", "備考"})
	w.Write([]string{"1", "キャベツ", "150", "とれたて新鮮"})
	w.Write([]string{"2", "にんじん", "100", ""})
	w.Write([]string{"3", "サンマ", "99", "今日の特売品"})

	//フラッシュ
	w.Flush()
}

func read() {
	//file open
	file, err := os.OpenFile("test.csv", os.O_RDONLY, 0)

	//error check
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//file close(defer)
	defer func() {
		file.Close()
	}()

	//csv.Readerを作成
	r := csv.NewReader(file)

	for {
		//一行ずつ読み込み
		record, err := r.Read()
		if err != nil {
			break
		}

		//読み込んだデータを出力
		fmt.Println(record)
	}
}
