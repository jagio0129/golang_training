package main

import (
	"io/ioutil"
	"log"

	"fmt"

	simplejson "github.com/bitly/go-simplejson"
)

func UseSimpleJSON() {
	//jsonファイルの読み込み
	bytes, err := ioutil.ReadFile("./simple.json")
	if err != nil {
		log.Fatal(err)
	}

	// simplejsonで操作
	json, err := simplejson.NewJson(bytes)
	if err != nil {
		log.Fatal(err)
	}

	for i, _ := range json.MustArray() {
		fmt.Printf("id: %d\n", json.GetIndex(i).Get("id").MustInt())
		fmt.Printf("name: %s\n", json.GetIndex(i).Get("name").MustString())
		fmt.Printf("birthday: %s\n", json.GetIndex(i).Get("birthday").MustString())
		fmt.Printf("vivid_info:color: %s\n", json.GetIndex(i).Get("vivi_info").Get("color").MustString())
		fmt.Printf("vivid_info:weapon: %s\n", json.GetIndex(i).Get("vivi_info").Get("weapon").MustString())
	}
}
