package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	myRand()
}

// 16進数ダンプ出力
func hexDump() {
	//256バイトのスライスを用意
	data := make([]byte, 256)

	//値を設定
	for i := 0; i < len(data); i++ {
		data[i] = byte(i)
	}

	//ダンプ出力
	fmt.Println(hex.Dump(data))

}

// 乱数生成
func myRand() {
	//乱数の初期値(現在時刻を乱数の元として使用)
	rand.Seed(time.Now().UnixNano())

	//int型の乱数
	fmt.Println(rand.Int())

	//float32型の乱数
	fmt.Println(rand.Float32())

	//float64型の乱数
	fmt.Println(rand.Float64())
}
