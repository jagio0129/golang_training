package main

import "os/exec"
import "fmt"
import "os"

func main() {
	myExecute()
}

//外部プログラムを実行する
func myExecute() {
	//「go help」を実行するためのCmdを作成
	cmd := exec.Command("go", "help")

	//実行し、コマンドが出力した結果を取得
	result, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//実行結果を出力
	fmt.Printf("%s\n", result)
}
