package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// 扱うデータ構造がわからない場合
func unknown() {
	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile("unknown.json")
	if err != nil {
		log.Fatal(err)
	}

	//構造がわからないのでinterface{}の中に解析する。
	var f interface{}
	if err := json.Unmarshal(bytes, &f); err != nil {
		log.Fatal(err)
	}

	// アサーションによるデータアクセス
	m := f.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k, "is float64", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I dont know how to handle")
		}
	}
}
