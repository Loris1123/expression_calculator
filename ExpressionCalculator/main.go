package main

import (
	"fmt"
	"flag"
)

var a,b,c,d float32
var expression string

func main(){
	initFlags()
	fmt.Println("Hallo Expression Calculator")

	println(a)
	println(b)
	println(c)
	println(d)
	println(expression)
}

func initFlags(){
	flag.IntVar(&a,"a",0, "Value for a")
	flag.IntVar(&b,"b",0, "Value for a")
	flag.IntVar(&c,"c",0, "Value for a")
	flag.IntVar(&d,"d",0, "Value for a")
	flag.StringVar(&expression,"exp","","The expression to be calculated")
	flag.Parse()
}
