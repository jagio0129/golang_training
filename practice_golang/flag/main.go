package main

import (
	"flag"
	"fmt"
)

var (
	intOpt  = flag.Int("i", 1234, "help message for \"i\" option")
	boolOpt = flag.Bool("b", false, "help message for \"b\" option")
	strOpt  = flag.String("s", "default", "help message for \"s\" option")
)

func main() {
	flag.Parse()

	fmt.Println(*intOpt)
	fmt.Println(*boolOpt)
	fmt.Println(*strOpt)
}
