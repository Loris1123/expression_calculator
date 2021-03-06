/*
* Author: https://github.com/Loris1123
* Date: 2013-03-1
* Last Update: 2013-03-3
* File: main.go
* Description: Main file for the expression calculator.
*/

package main

import (
	"fmt"
	"flag"
	"expression_calculator/interpreter"
)

var a,b,c,d float64
var expression string

func main(){
	initFlags()
	fmt.Println("Hallo Expression Calculator")

	fmt.Println(interpreter.Interpret_simple(a,b,c,d,expression))
}

func initFlags(){
	flag.Float64Var(&a,"a",0, "Float value for a")
	flag.Float64Var(&b,"b",0, "Float value for b")
	flag.Float64Var(&c,"c",0, "Float value for c")
	flag.Float64Var(&d,"d",0, "Float value for d")
	flag.StringVar(&expression,"exp","","The expression to be calculated")
	flag.Parse()
}
