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
