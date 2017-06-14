package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// JSONから構造体定義を行ってくれるJSON-to-Goが便利
// https://mholt.github.io/json-to-go/
type Simple []struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Birthday  string `json:"birthday,omitempty"`
	VividInfo struct {
		Color  string `json:"color"`
		Weapon string `json:"weapon"`
	} `json:"vivid_info,omitempty"`
}

func simple() {
	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile("simple.json")
	if err != nil {
		log.Fatal(err)
	}
	// JSONデコード
	var simples Simple
	if err := json.Unmarshal(bytes, &simples); err != nil {
		log.Fatal(err)
	}

	// デコードしたデータを表示
	for _, p := range simples {
		fmt.Printf("%d : %s : %s : %s : %s\n", p.ID, p.Name, p.Birthday, p.VividInfo.Color, p.VividInfo.Weapon)
	}

}
