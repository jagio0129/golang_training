# JSONの処理
JSONを扱う場合[encoding/json](https://golang.org/pkg/encoding/json/)を使う。Goの場合、JSONデータを定義した構造体に入れるのが一般的。

## JSONの解析

```
{
    "id": 1,
    "name": "nana",
    "birthday": "08-16",
    "vivid_info": {
      "color": "red",
      "weapon": "Rang"
    }
}
```
このようなデータ構造の場合、以下のように構造体を定義する。
```go
type Person struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Birthday string `json:"birthday"`
	VividInfo struct {
		Color string `json:"color"`
		Weapon string `json:"weapon"`
	} `json:"vivid_info"`
}
```

JSONから構造体のコードを生成してくれる[JSON-to-Go](https://mholt.github.io/json-to-go/)を使うと簡単。

JSONをデコードするには`json.Unmarshal`関数を使う。
```go
func Unmarshal(data []byte, v interface{}) error
```

デコードされたデータは以下のようにして扱うことができる。
```go
for _, p := range simples {
  fmt.Printf("%d : %s : %s : %s : %s\n", p.ID, p.Name, p.Birthday, p.VividInfo.Color, p.VividInfo.Weapon)
}

// 表示結果
// 1 : nana : 08-16 : red : Rang
```

## 未知のデータ構造をもつJSONの解析
JSONデータの構造を事前に知らない場合、interfaceを用いて解析する。JSONパッケージでは`map[string]interface{}`と`interface{}`構造を採用して任意のJSONオブジェクトを保存する。Goの型とJSONの型の対応は以下。
- boolはJSON booleansを表す。
- float64はJSON numberを表す。
- stringはJSON stringを表す。
- nilはJSON nilを表す。

以下のようなJSONデータがあるとし、この構造を把握していないと仮定する。
```golang
b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
```

これをinterface{}の中に解析。
```golang
var f interface{}
err := json.Unmarshal(b,&f)
```

このt気fの中にはmap型が保存されている。これらのkeyはstringで、値はからのinterface{}の中に保存される。
```golang
f = map[string]interface{}{
  "Name": "Wednesday",
	"Age":  6,
	"Parents": []interface{}{
		"Gomez",
		"Morticia",
	},
}
```

これらのデータにアクセスするには以下のようにする。
```golang
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
```
(めんどくさいのでJSONを取得してJSON-to-Goを利用して構造体にぶち込んだほうが見通しが良さそう。)

## [go-simplejson](https://github.com/bitly/go-simplejson)
上の例よりも簡単にJSONを扱うことのできるパッケージ
```
go get github.com/bitly/go-simplejson
```

### 使い方
下記のjsonを扱う場合、
```
{
    "hoge": true,
    "piyo":{
        "foo":[1,2],
        "bar":"test"
    }
}
```

```golang
// []byte型からjson型へ変換
json,err := simplejson.NewJson(bytes)

// Get("hoge").Type()またはGet("hoge").MustType()でアクセスできる。
// Type()は値とエラーを返す
b, _ := json.Get("hoge").Bool() // => true
 // MustType()は値のみ
m := json.Get("piyo").MustMap() // => map[bar:test foo:[1 2]]
a, _ := json.Get("piyo").Get("foo").Array()  // => [1,2]
// GetPath()で一括記述
s := js.GetPath("piyo", "bar").MustString() // => test
```

以下のような、配列で構成されるJSONデータの扱いは
```
[
  {
    "id": 1,
    "name": "akane",
    "birthday": "08-16",
    "vivid_info": {
      "color": "red",
      "weapon": "Rang"
    }
  },
  {
    "id": 2,
    "name": "aoi",
    "birthday": "06-17",
    "vivid_info": {
      "color": "blue",
      "weapon": "Impact"
    }
  },
  {
    "id": 3,
    "name": "wakaba",
    "birthday": "05-22",
    "vivid_info": {
      "color": "green",
      "weapon": "Blade"
    }
  },
  {
    "id": 4,
    "name": "himawari",
    "birthday": "07-23",
    "vivid_info": {
      "color": "yellow",
      "weapon": "Collider"
    }
  },
  {
    "id": 0,
    "name": "rei"
  }
]
```
このようにforで回してアクセス。他にもっといい方法あったら教えてください。
```golang
for i, _ := range json.MustArray() {
  fmt.Printf("id: %d\n", json.GetIndex(i).Get("id").MustInt())
  fmt.Printf("name: %s\n", json.GetIndex(i).Get("name").MustString())
  fmt.Printf("birthday: %s\n", json.GetIndex(i).Get("birthday").MustString())
  fmt.Printf("vivid_info:color: %s\n", json.GetIndex(i).Get("vivi_info").Get("color").MustString())
  fmt.Printf("vivid_info:weapon: %s\n", json.GetIndex(i).Get("vivi_info").Get("weapon").MustString())
}
```

rubyに比べると少し面倒な印象。

## 参考
- [Go言語でJSONを扱う](http://qiita.com/nayuneko/items/2ec20ba69804e8bf7ca3)
- [Goで簡単にJSON形式を扱うパッケージ： go-simplejson](http://qiita.com/TalesofFox/items/5c147a19a9ae5c41f41a)
- [astaxie/build-web-application-with-golang](https://github.com/astaxie/build-web-application-with-golang/blob/master/ja/07.2.md)