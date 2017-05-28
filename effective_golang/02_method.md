# 関数
## 複数の戻り値
- Goの新しい特徴の一つは、関数及びメソッドの値が返せること。
    - C言語で扱いにくかった、受信エラー(EOFを表す01など)や引数の値の変更などを改善する。
- Go言語では、Writeは書き込みデータ数とエラーを別々に返すことができる。
    - 「何バイト書き込めたが、デバイスがいっぱいになったので一部書き込めなかった」というような情報を得ることができる。

osパッケージの*File.Writeのシグネチャは次のように定義される。
```golang
func(file *File) Write(b []byte)(n int,err Error)
```
このメソッドは、**書き込んだバイト数を返す**とともに、**n != len(b)のとき非nilのErrorを返す**。

同様のアプローチにより、戻り値に参照パラメータを模してポインタを返す必要がなくなる。次の関数は、バイト列の指定位置から数値を取り出し、その値と次の値を返す単純な関数。
```golang
func netInt(b []byte,i int)(int,int){
    for ; i< len(b) && !isDigit(b[i]); i++{
    }
    x := 0
    for ; i < len(b) && isDigit(b[i]); i++ {
        x = x*10 + int(b[i])-'0'
    }
    return x,i
}
```
この関数は次のようにして、入力配列から数値を取り出すために利用できる。
```golang
for i := 0; i < len(a); {
    x, i = nextInt(a,i)
    fmt.Println(x)
}
```

## 名前付き結果パラメータ
戻り/結果「パラメータ」には、名前をつけることができ、引数パラメータの用に通常変数として扱うことができる。名前がつけられていると、関数が呼び出されたときにその時点で結果パラメータに格納されている値が、戻り値として使われる。名前をつけることでコードをより簡潔にすることができる。
```golang
func nextInt(b []byte, pos int)(value, nextPos int)
```

名前付きの結果パラメータは、初期化が行われた上で、パラメータなしのreturnと結び付けられるので、明確になるだけでなくシンプルになる。下は、この仕組をうまく使ったio.ReadFullの例。
```golang
func ReadFull(r Reader, buf []byte)(n int,err os.Error){
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:len(buf)]
    }
    return
}
```

## Defer
Go言語のdeferステートメントは、**deferを実行した関数がリターンする直前に、指定した関数の呼び出しが行われるようにスケジューリングする**。これは一般的な方式ではないが、**関数内のどの経路を通ってリターンするかにかかわらず、開放しなければならないリソースなどを処理するのに効果的な方式**。よくある例としては、ミューテックスのアンロックや、ファイルのクローズに使われている。
```golang
// Contentsは、ファイルの内容を文字列として返す。
func Contents(filename string)(string,os.Error){
    f,err := os.Open(filename)
    if err != nil {
        return "",err
    }
    defer f.Close() // f.Closeは完了時に実行される。

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        rresult = append(result,buf[0:n]...)
        if err != nil {
            if err == os.EOF {
                break
            }
            return "", err //ここでリターンしたときには、fはクローズされる
        }
    }
    return string(result), nil // ここでリターンしたときに、fはクローズされる。
}
```
Closeのような関数の呼び出しを遅延させることには、2つの利点がある。一つ目は**ファイルの閉じ忘れがないことを保証する**。後から関数を修正し、新しいリターン経路を付け加えたときに起こしやすいミスに対処できる。2つ目に**closeがopenの近くに記述されるので、closeを関数の最後に記述するよりコードが見やすくなる。**

遅延指定された関数の引数(関数がメソッドであれば、レシーバも含む)は、**関数の実行時ではなく、deferを実行した時に評価される**。それ以外にも、**関数の実行によって変数の値が変更されることを気にする必要がなくなる**。以下は、一箇所のdefer呼び出しで、複数の関数の実行を遅延させられることを意味する。
```golang
for i := 0;i < 5; i++{
    defer fmt.Printf("%d ",i)
}
```
遅延実行された関数は**LIFO順に実行される**ので、このコードでは、関数からリターンするときに`4 3 2 1 0`と出力される。

もう少し実用的な例は、プログラムの関数の実行を簡単にトレースする方法。次のように、シンプルなトレースルーチンを書くことができる。
```golang
func trace(s string) { fmt.Println("entering:",s)}
func untrace(s string) { fmt.Println("entering:",s)}

//これらの関数を、次のように使います
func a(){
    trace("a")
    defer untrace("a")
    //何か処理をする
}
```
遅延指定された関数の引数が、defer実行時に評価されることを利用した、もっと良い手もある。
```golang
func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string){
    fmt.Println("leaving:",s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b(){
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main(){
    b()
}
```