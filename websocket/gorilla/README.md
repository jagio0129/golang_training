# gorilla websocket
- https://github.com/gorilla/websocket

以下は[Gorilla websocket package](http://www.gorillatoolkit.org/pkg/websocket)の翻訳

## Installation
```
$ go get github.com/gorilla/websocket
```

## Overview
websocketパッケージはRFC6455で定義されているWebsocketプロトコルを実装している。

Conn型はWebsocket接続を表す。サーバアプリケーションは、HTTPリクエストハンドラを持つUpgraderオブジェクトのUpgrade機能を使用して、Connへのポインタを取得する。

```go
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//... Use conn to send and receive messages.
}
```

接続のWriteMessageメソッドとReadMessageメソッドを呼び出して、メッセージをバイトのスライスとして送受信する。このコードのスニペットは、次のメソッドを使用してメッセージをエコーする方法を示す。

```go
for {
    messageType, p, err := conn.ReadMessage()
    if err != nil {
        return
    }
    if err = conn.WriteMessage(messageType, p); err != nil {
        return err
    }
}
```

上記のスニペットコードでは、pは[]byteで、messageTypeはwebsocket.BinaryMessageかwebsocket.TextMessageかを持つintである。

アプリケーションは、io.WriteCloserおよびio.Readerインターフェイスを使用してメッセージを送受信することもできる。**メッセージを送信するには**、NextWriterメソッドを呼び出してio.WriteCloserを取得し、メッセージをwriterに書き込み、完了したらwriterを閉じる。**メッセージを受信するには**、NextReaderメソッドを呼び出してio.Readerを取得し、io.EOFが返されるまで読み取る。以下のスニペットには、NextWriterメソッドとNextReaderメソッドを使用してメッセージをエコーする方法を示す。

```go
for {
    messageType, r, err := conn.NextReader()
    if err != nil {
        return
    }
    w, err := conn.NextWriter(messageType)
    if err != nil {
        return err
    }
    if _, err := io.Copy(w, r); err != nil {
        return err
    }
    if err := w.Close(); err != nil {
        return err
    }
}
```

## Data Messages
WebSocketプロトコルは、テキストメッセージとバイナリメッセージを区別する。テキストメッセージはUTF-8エンコードされたテキストとして解釈される。バイナリメッセージの解釈はアプリケーションに委ねられる。

このパッケージは、TextMessageとBinaryMessageの整数の定数を使用して2つのデータメッセージタイプを識別する。ReadMessageメソッドとNextReaderメソッドは、受信したメッセージのタイプを返す。WriteMessageメソッド及びNextWriterメソッドの引数:messageTypeは、送信されるメッセージタイプを指定する。

テキストメッセージが有効なUTF-8エンコードされたテキストであることを確認するのは、アプリケーションの責任。

## Control Messages
WebSocketプロトコルは、close・ping・pongの3種類の制御メッセージを定義する。WriteControl、WriteMessage、またはNextWriterメソッドを呼び出して、制御メッセージをpeer(相手?)に送信する。

コネクションハンドルは、peerにcloseメッセージを送信し、NextReader、ReadMessageなどのメッセージReadメソッドから*CloseErrorを返すことで、closeメッセージを受け取る。

コネクションハンドルは、SetPingHandlerメソッドやSetPongHandlerメソッドで設定したコールバック関数を呼び出すことにより、pingやpongメッセージを受け取る。コールバック関数はNextReaderメソッドやReadMessageメソッドなどのメッセージReadメソッドから呼び出される。

デフォルトのpingハンドラーは相手にpongを送る。ハンドラがpongデータをコネクションに書き込んでいる間、アプリケーションの読み取り用ゴルーチンは短い時間の間ブロックする。

アプリケーションは、相手から送信されたpingやpongおよびcloseメッセージを処理するためにコネクションを読み取る必要がある。アプリケーションが相手からのメッセージに関心がなければ、相手からのメッセージを読み込んで破棄するためにゴルーチンを開始する。簡単な例は以下の通り。

```go
func readLoop(c *websocket.Conn) {
    for {
        if _, _, err := c.NextReader(); err != nil {
            c.Close()
            break
        }
    }
}
```