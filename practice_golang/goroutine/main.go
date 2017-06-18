package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	channel04()
}

func channel01() {
	//チャネルの作成
	c := make(chan int)

	// ゴルーチン。この関数のパラメータは送信専用チャネル
	go func(s chan<- int) {
		//チャネルへ0～4の値を順番に送信
		for i := 0; i < 5; i++ {
			s <- i
		}
		//チャネルのクローズ
		close(s)
	}(c)

	//受信ループ
	for {
		//チャネルからの受信
		// クローズされたか知る必要がなければ
		// 2番めの値を受け取る必要はない
		value, ok := <-c

		//チャネルがクローズされるとokにfalseが返される
		if !ok {
			break
		}

		//受信した値を出力
		fmt.Println(value)

	}
}

func channel02() {
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

func channel03() {
	done := make(chan bool)
	go func() {
		time.Sleep(time.Second * 3)
		done <- true
	}()
	<-done //受け取るまでブロック
	fmt.Println("done")
}

func channel04() {
	receiver, fin := worker("job")
	for i := 0; i < 3; i++ {
		select {
		case receive := <-receiver:
			fmt.Printf(receive)
		case <-fin:
			return
		}
	}
}

func worker(msg string) (<-chan string, <-chan bool) {
	var wg sync.WaitGroup
	receiver := make(chan string)
	fin := make(chan bool)

	go func() {
		for i := 0; i < 3; i++ {
			wg.Add(1)
			go func(i int) {
				msg := fmt.Sprintf("%d %s done", i, msg)
				receiver <- msg
				wg.Done()
			}(i)
		}
		wg.Wait()
		fin <- false //終了を伝える
	}()
	return receiver, fin
}
