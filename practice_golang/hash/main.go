package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
)

func main() {
	/*
	 * SHA-1
	 */

	//SHA-1のハッシュの準備
	s := sha1.New()

	//テキストデータを書き込む
	io.WriteString(s, "test")

	//書き込んだテキストデータのハッシュ値を得る
	// Sumメソッドにスライスを渡すことで、そのスライスへの追記も可能
	result1 := s.Sum(nil)

	//算出したSHA-1の値をダンプ出力
	fmt.Printf("SHA-1:% x\n", result1)

	/*
	 * MD5
	 */
	m := md5.New()

	//ハッシュ算出元データ
	data := []byte{0x74, 0x5, 0x73, 0x74} //test

	//バイト配列のデータを書き込む
	m.Write(data)

	//書き込んだバイト配列のデータのハッシュ値を得る
	result2 := m.Sum(nil)

	//算出したSHA-1の値をダンプ出力
	fmt.Printf("MD5:% x\n", result2)

}
