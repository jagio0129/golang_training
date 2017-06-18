# ゴルーチン
Go言語において、並行処理を実現するためのしくみ
- [Coroutine](http://qiita.com/fujimisakari/items/811e350cbaeb45b6165e)ではない
- Thread,Processでもない
- 複数のThread上に多重化されて実装されている
- main()自信やScavenger(GC)などランタイムもゴルーチンをつかっている

# 並行(Concurrency)処理とは
- Concurrency => 並行
  - 同時に幾つかの質の**異なることをおこなう**
- Parallelism => 並列
  - 同時に幾つかの質の**同じことを行う**

## Goの並列処理モデル
並列処理プログラミングには2つのアプローチがある

**Shared-memory communication**
複数のプロセスがロックを取りながら共通のメモリにアクセスする。

**Message-passing communication**
各プロセスはメッセージを送り合い、内容は書き変わらない。

Goの場合は後者を採用している。

## ゴルーチンで並行を実現
- 複数のゴルーチンで同時に複数のタスクをこなす
- 各ゴルーチンに役割を与えて分業する

## 軽量なスレッドのようなもの
- LinuxやUnixのスレッドより**コストが低い**
- 一つのスレッドの上で複数のゴルーチンが動く

## ゴルーチンのつくりかた
- `go func()`

# チャネル
- Go言語プログラム内での通信機構で、ゴルーチン間における**通信、同期、値の共有**に使用する。
- チャネルで通信できるのは同一プログラム内だけ
- チャネルを使用するにはまずチャネル型の値を作成し、送信側ではその値に対して何らかの値を送信する。もう一方の受信側では、同じチャネルを使って受信を行う。

```go
chan 要素型
chan <- 要素型   //送信専用チャネル型
<- chan 要素型   //受信専用チャネル型
```

## チャネルの作成
- チャネルはスライスやマップと同じく参照型の一種であり、他の参照型と同様に値を作成するには`make`関数を使用する。
- 第一パラメータには「チャネル型」、第二パラメータには「キャパシティ」を指定する。
  - チャネルのキャパシティとは、チャネル内でバッファリング可能な要素数の上限、バッファ容量。

```go
make(chan 要素数,キャパシティ)
make(chan 要素型)
```

## チャネルへの値の送受信
```go
チャネル <- 送信する値 //送信
<- チャネル
```

## チャネルのクローズ
- 使用しなくなったチャネルはcloseする。
  - ただし、受信専用チャネルはクローズできない
  - クローズ済みのチャネルに値を送信することはできない
- クローズ済みのチャネルから値を受信しようとすると、クローズ直前まで送信されてチャネルでバッファリングされている値がなくなるまで受信できる。
  - バッファ内の値を全て受信し尽くした後は、そのチャネルの要素型のゼロ値が返る。
  - ゼロ値を受信したからといってそのチャネルがクローズされたかは判定できない。


```go
close(チャネル)
```

## 複数のチャネルから同時に受信
`select`構文を使うと複数のチャネルの受信を同時に行うことができる

```go
func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	go func() { ch1 <- 100 }()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "hello"
	}()

	select {
	case v1 := <-ch1:
		fmt.Println(v1)
	case v2 := <-ch2:
		fmt.Println(v2)
	}
}
```

## 参考
- [Goroutineと channelから はじめるgo言語](https://www.slideshare.net/takuyaueda967/goroutine-channel-go)
- [sync.WaitGroupの正しい使い方](http://qiita.com/ruiu/items/dba58f7b03a9a2ffad65)
- [Go の並行処理](http://jxck.hatenablog.com/entry/20130414/1365960707)
- [Go言語のchannelって一体何よ ~基礎編~【golang】](http://otiai10.hatenablog.com/entry/2014/01/22/095902)
- [GoのChannelを使いこなせるようになるための手引](http://qiita.com/awakia/items/f8afa070c96d1c9a04c9)
- [chanの使い方パターンメモ。](http://golang.rdy.jp/2015/03/25/chan_tips/)
- [golang の channel を使ったテクニックあれこれ](http://mattn.kaoriya.net/software/lang/go/20160706165757.htm)