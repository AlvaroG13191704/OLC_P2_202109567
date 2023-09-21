package compiler

import (
	"fmt"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

// visit printstmt
func (v *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) interface{} {

	exprList := ctx.ExprList().(*parser.ExprListContext)
	expressiongs := exprList.AllExpr()

	for _, expression := range expressiongs {

		value := v.Visit(expression).(values.PRIMITIVE) // visit the expression

		//evalue the type of the value
		if value.GetType() == values.IntType || value.GetType() == values.FloatType {
			// generate c3d
			v.Generator.GenPrint("d", "(int)"+fmt.Sprintf("%v", value.GetValue()))
			v.Generator.GenPrint("c", "10")
			v.Generator.AddNewLine()
		} else {

			// TODO: add error
			return nil
		}

	}

	return nil

}
