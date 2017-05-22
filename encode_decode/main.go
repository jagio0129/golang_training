package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
)

func main() {
	myUrl()
}

func myBase64() {
	//データ
	data := []byte{0x00, 0x01, 0x02, 0x03, 0x04}

	// BASE64エンコード
	enc := base64.StdEncoding.EncodeToString(data)

	fmt.Println(enc)

	//エンコード結果をBASE64デコード
	dec, _ := base64.StdEncoding.DecodeString(enc)

	fmt.Printf("% x\n", dec)
}

func myUrl() {
	//データ
	data := "abcあいう/:"

	//URLエンコード
	enc := url.QueryEscape(data)

	fmt.Println(enc)

	//エンコード結果をURLデコード
	dec, _ := url.QueryUnescape(enc)

	fmt.Println(dec)
}
