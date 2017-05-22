package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	/*
		// リスト作成
		list := mkList()

		//作成したリストの要素一覧表示
		showElementList(list)

		//指定したindexの値を確認
		fmt.Println(getElement(list, 3).Value)

		// 指定したindexの値を変更
		changeElement(list, 3, "test")
		fmt.Println(getElement(list, 3).Value)
	*/

	list1 := mkList()
	list2 := mkList()
	mixList(list1, list2)
	showElementList(list1)

	// リストを空にする
	list1.Init()
	showElementList(list1)
}

func mkList() *list.List {
	//からのリストを作成
	l := list.New()

	//リストに要素を追加
	l.PushBack("hyde")
	l.PushBack("tetsu")
	l.PushBack("ken")
	l.PushBack("yukihiro")

	return l
}

func getListLen(l *list.List) {
	fmt.Println(reflect.TypeOf(l))
	//リストの要素数を取得する
	fmt.Println("要素数:", l.Len())
}

func changeElement(list *list.List, index int, newVal string) {
	//リストから4番目(index:3)の要素を取り出す
	target := getElement(list, index)
	//要素の入れ替え
	target.Value = newVal
}

//リストの要素(element)を取り出す関数
func getElement(l *list.List, index int) *list.Element {
	//リストの要素をイテレート
	for e, i := l.Front(), 0; e != nil; e = e.Next() {
		//してしたインデックスと一致したら、要素を返す
		if i == index {
			return e
		}
		i++
	}
	panic("インデックス不正")
}

func showElementList(list *list.List) {
	fmt.Println("リストの要素一覧")
	for e := list.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

//2つのリストを結合する
func mixList(list1, list2 *list.List) {
	//list1にlist2の内容を追加
	list1.PushBackList(list2)
}
