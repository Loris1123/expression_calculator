/*
* Author: https://github.com/Loris1123
* Date: 2013-03-1
* Last Update: 2013-03-3
* File: simple_interpreter.go
* Description: This is the simple interpreter which can calculate only simple expressions. (No exponents or numbers in expression)
*/

package interpreter

import "strings"

var a,b,c,d float64
//TODO: Symbol as character
var symbol string = ""
var index int = 0
//TODO: Input as char array
var input = make([]string, 0)  //initial capacity 0 and no values. Overwrite later
var end_of_input bool = false

func Interpret_simple(locA,locB,locC,locD float64, expression string) float64{
	result := 0.0
	//check_success(expression != nil && len(s)>0, "Expression is nil or too short")
	a = locA
	b = locB
	c = locC
	d = locD

	input = strings.Split(expression,"")  //Overwrite input with the real input.

	read_next_symbol()

	if !end_of_input{
		result = simple_expression()
	}
	return result

}

/*
* Reads an expression from input
* EBNF Rule:
* SimpleExpression = [AddOp] Term { AddOb Term}
* AddOp = "+" | "-"
*/
func simple_expression() float64{
	check_success(!end_of_input, "Unexpected end of input")
	result := 1.0
	if symbol == "+"{
		read_next_symbol()
	}else if symbol == "-"{
		result = -1
		read_next_symbol()
	}
	result = result * term()
	for !end_of_input && (symbol == "+" || symbol == "-"){
		addOp := symbol
		read_next_symbol()
		result = float_value(result, addOp, term())
	}
	return result
}

/*
* Reads a term from input
* EBNF Rule:
* Term = Factor {MulOp Factor}
* MulOp = "*"|"/"
*/
func term() float64{
	result := factor()
	for !end_of_input && (symbol == "*" || symbol == "/"){
		mulOp := symbol
		read_next_symbol()
		result = float_value(result, mulOp, factor())
	}	
	return result
}

/*
* Reads a factor from input
* EBNF Rule:
* Factor = Ident | "("SimpleExpression")"
*/
func factor() float64{
	result := 0.0
	check_success(!end_of_input,"Unexpected end of input")
	if symbol == "a" || symbol == "b" || symbol == "c" || symbol == "d"{
		result = ident()
	} else if symbol == "("{
		read_next_symbol()
		result = simple_expression()
		check_success(!end_of_input,"Unexpected end of input")
		if symbol == ")"{
			read_next_symbol()
		}else{
			panic("Unexpected "+symbol+" -- Expecting )")
		}
	}else{
		panic("Unexpected "+symbol+" -- Expecting ( or Ident")
	}
	return result
}

/*
* Reads an ident from input
* EBNF Rule: 
* ident = "a" | "b" | "c" | "d"
*/
func ident() float64{
	result := 0.0
	switch symbol{
		case "a": result = a
		case "b": result = b
		case "c": result = c
		case "d": result = d
		default: panic("Unknown variable: "+ symbol)
	}
	read_next_symbol()
	return result
}

func read_next_symbol(){
	if index >= 0 && index < len(input){
		symbol = input[index]
		index ++
	}else if index == len(input){
		end_of_input = true
	}
}


//TODO: operatator as character
func float_value (left float64, operator string, right float64) float64{
	switch operator {
		case "+":
			return left+right
		case "-":
			return left-right
		case "*":
			return left*right
		case "/":
			return left/right
		default:
			//Unreachable
			panic("Unknown Operator: "+ operator)
	}
	//Unreachable
	return 0
}

/*
*Error Handling. Creates a panic when condition is wrong
*/
func check_success(condition bool, message string){
	if(!condition){
		panic(message)
	}

}


