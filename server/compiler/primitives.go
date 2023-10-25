package compiler

import (
	"fmt"
	"server/compiler/compiler/values"
	"server/compiler/parser"
	"strings"
)

// visit digitexpr
func (v *Visitor) VisitDigitExpr(ctx *parser.DigitExprContext) interface{} {
	digit := ctx.GetText() // get the digit
	// evalue if the digit has a . to know if it is a float or an integer
	if strings.Contains(digit, ".") {
		// f, _ := strconv.ParseFloat(digit, 64) // convert to float
		return values.NewC3DPrimitive(digit, "", values.FloatType, false)

	} else {
		// i, _ := strconv.ParseInt(digit, 10, 64) // convert to integer
		// fmt.Println("Digito primitivo int: ", i)
		return values.NewC3DPrimitive(digit, "", values.IntType, false)

	}
}

// visit stringexpr and char
func (v *Visitor) VisitStringExpr(ctx *parser.StringExprContext) interface{} {
	str := strings.Trim(ctx.GetText(), "\"") // get the string

	// fmt.Println("Primitive String: ", str)
	return values.NewC3DPrimitive(str, "", values.StringType, false)
}

// visit char
func (v *Visitor) VisitCharExpr(ctx *parser.CharExprContext) interface{} {
	str := strings.Trim(ctx.GetText(), "'") // get the char
	// conver to ascii
	char := fmt.Sprintf("%v", int(str[0]))
	return values.NewC3DPrimitive(fmt.Sprintf("%s", char), "", values.CharType, false)

}

// visit boolean
func (v *Visitor) VisitBooleanExpr(ctx *parser.BooleanExprContext) interface{} {
	value := strings.Trim(ctx.GetText(), "\"")
	// fmt.Println("Primitive Boolean: ", value)
	if value == "true" {
		// return &values.Boolean{Value: true}
		return values.NewC3DPrimitive("1", "", values.BooleanType, false)
	} else {
		return values.NewC3DPrimitive("0", "", values.BooleanType, false)
	}
}

// visit nil
func (v *Visitor) VisitNilExpr(ctx *parser.NilExprContext) interface{} {
	// fmt.Println("Primitive Nil")
	return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
}
