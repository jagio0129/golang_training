package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"golang.org/x/net/websocket"
)

var origin = "http://localhost/"
var url = "http://localhost/"

type Weather struct {
	Datetime time.Time `json:"datetime"`
	Term     int       `json:"term"`
	Value    struct {
		Thermal     float64 `json:"thermal"`
		Humidity    float64 `json:"humidity"`
		Illuminance float64 `json:"illuminance"`
		CO2         float64 `json:"CO2"`
	} `json:"value"`
}

func myJson() {
	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile("./test.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var weathers []Weather
	json.Unmarshal(bytes, &weathers)

	// デコードしたデータを表示
	for _, p := range weathers {
		fmt.Println(p.Datetime)
	}
}

func main() {
	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile("./test.json")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var weathers Weather
	json.Unmarshal(bytes, &weathers)

	// デコードしたデータを表示
	fmt.Print(weathers.Value)
}

func sendData() {
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	message := []byte("hello, world!")
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", message)

	var msg = make([]byte, 512)
	_, err = ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg)
}
