package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
)

func main() {
	httpPost()
}

//ホスト名からIPアドレスを取得する
func getIPAddress() {
	//ホスト名をIPアドレスに変換
	addrs, err := net.LookupHost("www.google.com")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//取得したIPアドレスを出力
	fmt.Printf("%q\n", addrs)
}

//TCP/IP通信を行う
func connectionTcpIp() {
	//ホスト(HTTPサーバ)に接続
	conn, err := net.Dial("tcp", "golang.jp:80")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//通信のクローズ
	defer func() {
		conn.Close()
	}()

	//HTTPリクエストの送信
	fmt.Fprintf(conn, "GET /hello.html HTTP/1.1\r\n")
	fmt.Fprintf(conn, "HOST: golang.jp\r\n")
	fmt.Fprintf(conn, "\r\n")

	//HTTPレスポンスの受信
	response, err := ioutil.ReadAll(conn)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//レスポンスをバイトスライスから文字列に変換し出力
	fmt.Println(string(response))
}

//HTTP通信
func httpGet() {
	//HTTP接続を開始
	res, err := http.Get("http://golang.jp/hello.html")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//レスポンスのステータスを出力
	fmt.Println("status", res.Status)

	//レスポンのボディをすべて読み込む
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//ボディを出力
	fmt.Println(string(body))
}

func httpPost() {
	//HTTP接続を開始
	res, err := http.PostForm("http://golang.jp/hello.html",
		url.Values{
			"key1": {"Value1", "Value2"},
			"key2": {"Value"}})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//レスポンスのステータスを出力
	fmt.Println("status:", res.Status)

	//レスポンスのボディをすべて読み込む
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//ボディを出力
	fmt.Println(string(body))
}
