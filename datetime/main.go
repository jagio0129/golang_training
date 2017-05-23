package main

import (
	"fmt"
	"time"
)

func main() {
	//現在時刻の取得
	//time := time.Now()
	//t := reflect.TypeOf(time)
	//fmt.Println(t)
	//fmt.Println(time)

	myProcceedTime()
}

//時刻をフォーマット
func format(time *time.Time) {
	//時刻を文字列にして出力
	fmt.Printf("%04d/%02d/%02d %02d/:%02d:%02d\n",
		time.Year(),
		time.Month(),
		time.Day(),
		time.Hour(),
		time.Minute(),
		time.Second())
}

//日時を指定してTime型の値を作成する
func makeTime() {
	//ローカルのタイムゾーン情報を取得
	loc, _ := time.LoadLocation("Local")

	// Time型の値を作成(2017/3/3 11:22:33.44)
	time := time.Date(2017, 3, 3, 11, 22, 33, 44, loc)

	//時刻の出力
	fmt.Println(time)
}

//TIme型に時間を加算・減算する
func changeTime() {

	myTime := time.Now()

	//時刻の表示
	fmt.Println(myTime)

	//日時を加算
	add := myTime.Add(time.Hour*24 + time.Minute*30 + time.Second*5)

	//時刻の減算
	dec := myTime.Add(time.Hour*24 - time.Minute*30 - time.Second*5)

	fmt.Println(add)
	fmt.Println(dec)
}

//曜日を取得する
func getWeek() {
	t := time.Now()
	fmt.Println(t)

	switch t.Weekday() {
	case time.Sunday:
		fmt.Println("日曜")
	case time.Monday:
		fmt.Println("月曜")
	case time.Tuesday:
		fmt.Println("火曜")
	case time.Wednesday:
		fmt.Println("水曜")
	case time.Thursday:
		fmt.Println("木曜")
	case time.Friday:
		fmt.Println("金曜")
	case time.Saturday:
		fmt.Println("土曜")
	}
}

//Sleep
func mySleep() {
	fmt.Println(time.Now())

	//10秒スリープ
	time.Sleep(10 * time.Second)

	fmt.Println(time.Now())

	//1分スリープ
	time.Sleep(1 * time.Minute)

	fmt.Println(time.Now())
}

//経過時間を算出する
func myProcceedTime() {
	start := time.Now()

	//15秒待つ
	time.Sleep(time.Second * 15)

	//スリープ後の現在日時を取得
	end := time.Now()

	//経過時間の算出(Duration型が返る)
	sub := end.Sub(start)

	fmt.Println(sub)
	fmt.Println(sub/time.Second, "秒")
}
