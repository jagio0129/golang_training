# データ
## newによる割り当て
Go言語の基本的なメモリ割り当てには、**newとmake**の2つの組み込み関数が用意されている。

### new
**メモリの割当を行う組み込み関数**だが、多言語のnewの多くと異なり、**メモリの初期化でなく、ゼロ化のみを行う**。すなわち、new(T)は、**型Tの新しいアイテム用にゼロ化した領域を割り当て、そのアドレスである\*T型の値を返す**。Go言語風に言い換えると、、**new(T)が返す値は、新しく割り当てられた型Tのゼロ値のポインタ**。

ゼロ化済みオブジェクトは、さらなる初期化を行わなくても使用できるため、こういったオブジェクトの準備にnewは便利。すなわち、**データ構造体の利用者がnewでそれを作成すると、すぐに使える状態となる。**例えば、byte.Bufferのドキュメントには、「Bufferのゼロ値は、利用準備が整ったからのバッファである」と記述されている。同様にsync.Mutexには明示的なコンストラクタやInitメソッドは用意されていないが、そのかわりにsync.Mutexのゼロ値は、非ロック状態のミューテックスであること定められている。

この便利なゼロ値は連鎖的に働く。
```golang
type SyncedBuffer struct {
    lock   syc.Mutex
    buffer bytes.Buffer
}
```
このSyncedBuffer型の値もまた、割り当てや宣言を行うと同時に準備が整う。下のコードの変数pとvは、このままで正しく機能する
```golang
p := new(SyncedBuffer)  // *SyncedBuffer型
var v SyncedBuffer      // SyncedBuffer型
```