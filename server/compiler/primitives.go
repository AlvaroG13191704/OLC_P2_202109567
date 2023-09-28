package compiler

import (
	"server/compiler/compiler/values"
	"server/compiler/parser"
	"strconv"
	"strings"
)

// visit digitexpr
func (v *Visitor) VisitDigitExpr(ctx *parser.DigitExprContext) interface{} {
	digit := ctx.GetText() // get the digit
	// evalue if the digit has a . to know if it is a float or an integer
	if strings.Contains(digit, ".") {
		f, _ := strconv.ParseFloat(digit, 64) // convert to float

		// fmt.Println("Digito primitivo float: ", f)
		return &values.Float{Value: f}
	} else {
		i, _ := strconv.ParseInt(digit, 10, 64) // convert to integer
		// fmt.Println("Digito primitivo int: ", i)
		return &values.Integer{Value: i}
	}
}

// visit stringexpr and char
func (v *Visitor) VisitStringExpr(ctx *parser.StringExprContext) interface{} {
	str := strings.Trim(ctx.GetText(), "\"") // get the string
	// fmt.Println("Primitive String: ", str)
	return &values.String{Value: str}
}

// visit char
func (v *Visitor) VisitCharExpr(ctx *parser.CharExprContext) interface{} {
	str := strings.Trim(ctx.GetText(), "'") // get the char
	// fmt.Println("Primitive Char: ", str)
	return &values.Character{Value: str}
}

// visit boolean
func (v *Visitor) VisitBooleanExpr(ctx *parser.BooleanExprContext) interface{} {
	value := strings.Trim(ctx.GetText(), "\"")
	// fmt.Println("Primitive Boolean: ", value)
	if value == "true" {
		return &values.Boolean{Value: true}
	} else {
		return &values.Boolean{Value: false}
	}
}

// visit nil
func (v *Visitor) VisitNilExpr(ctx *parser.NilExprContext) interface{} {
	// fmt.Println("Primitive Nil")
	return &values.Nil{
		Value: nil,
	}
}
