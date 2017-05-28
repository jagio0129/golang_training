package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

func main() {
	ls()
}

//ディレクトリの作成
func mkdir() {
	//カレントディレクトリ直下にディレクトリ作成
	os.Mkdir("newdir", 0777)

	//os.MkdirAllを使うと途中のディレクトリも作成される
	os.MkdirAll("a/b/c", 0777)
}

//ディレクトリの削除
func rmdir() {
	//rmdir(実行後、a/bは残る)
	os.Remove("a/b/c")

	//RemoveAllを使うと指定したディレクトリ以下が全て削除される
	os.RemoveAll("a")
}

//カレントディレクトリを変更する
func changeCd() {
	//カレントディレクトリを出力
	current, _ := os.Getwd()
	fmt.Println(current)

	//カレントディレクトリを親ディレクトリに変更
	os.Chdir("..")

	//カレントディレクトリを出力
	current, _ = os.Getwd()
	fmt.Println(current)

}

func rename(dir string, newName string) {
	err := os.Rename(dir, newName)

	if err != nil {
		fmt.Println(err)
	}
}

//テンポラリファイルの作成
func mkTempFile() {
	//get temporary dir name
	dir := os.TempDir()

	//テンポラリディレクトリ内にテンポラリファイルを作成
	file, _ := ioutil.TempFile(dir, "test")

	//作成したテンポラリファイル名の出力
	fmt.Println(file.Name())
}

//ディレクトリ内のファイル一覧を取得
func ls() {
	//一覧取得先のディレクトリとして$GOROOTを使用する
	goroot := runtime.GOROOT()
	fmt.Println("$GORROT:" + goroot)

	//ディレクトリないのファイル、ディレクトリ一覧を取得
	fileinfos, _ := ioutil.ReadDir(goroot)

	for _, fileinfo := range fileinfos {
		//今回はディレクトリは除外
		if !fileinfo.IsDir() {
			fmt.Println(fileinfo.Name())
		}
	}
}
