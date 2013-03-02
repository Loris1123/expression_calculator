/*
* Author: https://github.com/Loris1123
* Date: 2013-03-1
* Last Update: 2013-03-2
* File: simple_interpreter.go
* Description: This is the simple interpreter which can calculate only simple expressions. (No exponents or numbers in expression)
*/

package interpreter



var a,b,c,d float64
var expression, symbol string
//TODO: Symbol as character


func Interpret_simple(a,b,c,d float64, expression string){

	//check_success(expression != nil && len(s)>0, "Expression is nil or too short")

	println(a)
	println(b)
	println(c)
	println(d)
	println(expression)

	println(float_value(4,"/",5))


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


