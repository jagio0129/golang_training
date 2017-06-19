# コマンドラインオプションの扱い
Goでコマンドラインオプションを扱うときには**flagsパッケージ**を使用する。

## 使い方
```go
var (
	intOpt  = flag.Int("i", 1234, "help message for \"i\" option")
	boolOpt = flag.Bool("b", false, "help message for \"b\" option")
	strOpt  = flag.String("s", "default", "help message for \"s\" option")
)

func main(){
  flag.Parse()

	fmt.Println(*intOpt)
	fmt.Println(*boolOpt)
	fmt.Println(*strOpt)
}
```

```[実行結果]
$ go run main.go //何も指定しない場合はデフォルト値
1234
false
default

$ go run main.go -i 666 -b -s hello // オプションを指定した場合
666
true
hello

$ go run main.go -h  // helpも利用できる。
Usage of /tmp/go-build005022585/command-line-arguments/_obj/exe/main:
  -b    help message for "b" option
  -i int
        help message for "i" option (default 1234)
  -s string
        help message for "s" option (default "default")
exit status 2
```

よくある`-v, --verbose`みたいに省略形も指定できないか試してみた。

```go
var (
	boolFlag   bool
	intFlag    int
	stringFlag string
)

func register() {
	flag.BoolVar(&boolFlag, "bool", false, "help message for \"b\" option")
	flag.BoolVar(&boolFlag, "b", false, "help message for \"b\" option")
	flag.IntVar(&intFlag, "int", 1234, "help message for \"i\" option (default 1234)")
	flag.IntVar(&intFlag, "i", 1234, "help message for \"i\" option (default 1234)")
	flag.StringVar(&stringFlag, "string", "defalut", "help message for \"s\" option (default \"default\")")
	flag.StringVar(&stringFlag, "s", "defalut", "help message for \"s\" option (default \"default\")")

	flag.Parse()
}

func main() {
	register()

	fmt.Println(boolFlag)
	fmt.Println(intFlag)
	fmt.Println(stringFlag)
}
```

```
$ go run main.go -h
Usage of /tmp/go-build315847528/command-line-arguments/_obj/exe/main:
  -b    help message for "b" option
  -bool
        help message for "b" option
  -i int
        help message for "i" option (default 1234) (default 1234)
  -int int
        help message for "i" option (default 1234) (default 1234)
  -s string
        help message for "s" option (default "default") (default "defalut")
  -string string
        help message for "s" option (default "default") (default "defalut")
exit status 2
```

う～ん...

```go
func register() {
	flag.BoolVar(&boolFlag, "bool", false, "help message for \"b\" option")
	flag.BoolVar(&boolFlag, "b", false, "")
	flag.IntVar(&intFlag, "int", 1234, "help message for \"i\" option (default 1234)")
	flag.IntVar(&intFlag, "i", 1234, "")
	flag.StringVar(&stringFlag, "string", "defalut", "help message for \"s\" option (default \"default\")")
	flag.StringVar(&stringFlag, "s", "defalut", "")

	flag.Parse()
}
```

```
$ go run main.go -h
Usage of /tmp/go-build637651983/command-line-arguments/_obj/exe/main:
  -b
  -bool
        help message for "b" option
  -i int
         (default 1234)
  -int int
        help message for "i" option (default 1234) (default 1234)
  -s string
         (default "defalut")
  -string string
        help message for "s" option (default "default") (default "defalut")
exit status 2
```

やっぱりヘルプがきれいにならない。

## go-flags
- https://github.com/jessevdk/go-flags
- http://godoc.org/github.com/jessevdk/go-flags

よさげなパッケージがあったのでメモ

オプションを定義した構造体を用意
```go
type Options struct{
  Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
}
```
- short:`-v`
- long:`--verbose`
- description:説明
が指定できる。他にも`required:true`にすることでオプション指定を強制できる。その他の値は[こちら](http://godoc.org/github.com/jessevdk/go-flags)を参照。

次にオプションの解析
```go
import (
	flags "github.com/jessevdk/go-flags"
)

var opts Options

func main() {
	args, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	// go run main.go -vv arg1 のように起動された場合、以下のように値がバインドされる
	// opts.Verbose = [true, true]
	// args[0] = "arg1"
}
```

helpを見てみる

```
$ go run main.go -h
Usage:
  main [OPTIONS]

Application Options:
  -v, --verbose  Show verbose debug information

Help Options:
  -h, --help     Show this help message

exit status 1
```
ショートもロングも使え、表示もきれいになっている。

引数の数が足りない場合などに明示的に表示したい、Usageの内容を変更したい場合は
```go
func main() {
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "pt"
	parser.Usage = "[OPTIONS] PATTERN [PATH]"

	args, _ := parser.Parse()

	// 引数がひとつもなければヘルプを表示する
	if len(args) == 0 {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
}
```
のようにすれば良い。

以下はexample
```go
var opts struct {
	// Slice of bool will append 'true' each time the option
	// is encountered (can be set multiple times, like -vvv)
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`

	// Example of automatic marshalling to desired type (uint)
	Offset uint `long:"offset" description:"Offset"`

	// Example of a callback, called each time the option is found.
	Call func(string) `short:"c" description:"Call phone number"`

	// Example of a required flag
	Name string `short:"n" long:"name" description:"A name" required:"true"`

	// Example of a value name
	File string `short:"f" long:"file" description:"A file" value-name:"FILE"`

	// Example of a pointer
	Ptr *int `short:"p" description:"A pointer to an integer"`

	// Example of a slice of strings
	StringSlice []string `short:"s" description:"A slice of strings"`

	// Example of a slice of pointers
	PtrSlice []*string `long:"ptrslice" description:"A slice of pointers to string"`

	// Example of a map
	IntMap map[string]int `long:"intmap" description:"A map from string to int"`

	// Example of positional arguments
	Args struct {
		ID   string
		Num  int
		Rest []string
	} `positional-args:"yes" required:"yes"`
}

func main(){
  _, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Verbosity: %v\n", opts.Verbose)
	fmt.Printf("Offset: %d\n", opts.Offset)
	fmt.Printf("Name: %s\n", opts.Name)
	fmt.Printf("File: %s\n", opts.File)
	fmt.Printf("Ptr: %d\n", *opts.Ptr)
	fmt.Printf("StringSlice: %v\n", opts.StringSlice)
	fmt.Printf("PtrSlice: [%v %v]\n", *opts.PtrSlice[0], *opts.PtrSlice[1])
	fmt.Printf("IntMap: [a:%v b:%v]\n", opts.IntMap["a"], opts.IntMap["b"])
	fmt.Printf("Args.ID: %s\n", opts.Args.ID)
	fmt.Printf("Args.Num: %d\n", opts.Args.Num)
	fmt.Printf("Args.Rest: %v\n", opts.Args.Rest)
}
```

```
$ go run main.go -h
Usage:
  main [OPTIONS] ID Num Rest...

Application Options:
  -v, --verbose      Show verbose debug information
      --offset=      Offset
  -c=                Call phone number
  -n, --name=        A name
  -f, --file=FILE    A file
  -p=                A pointer to an integer
  -s=                A slice of strings
      --ptrslice=    A slice of pointers to string
      --intmap=      A map from string to int

Help Options:
  -h, --help         Show this help message
```

```go
Call func(string) `short:"c" description:"Call phone number"`
```
Callの使い方がわからなかった。

- https://github.com/jessevdk/go-flags
- http://godoc.org/github.com/jessevdk/go-flags
- http://blog.monochromegane.com/blog/2014/01/23/go-flags/
- http://ryochack.hatenablog.com/entry/2013/04/17/232753