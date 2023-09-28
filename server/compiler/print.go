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
		if value.GetType() == values.IntType {
			// generate c3d
			v.Generator.GenPrint("d", "(int)"+fmt.Sprintf("%v", value.GetValue()))
			v.Generator.GenPrint("c", "10")
			v.Generator.AddNewLine()

		} else if value.GetType() == values.FloatType {
			floatValue := value.GetValue().(float64)
			format := "f"

			// Use %e for large float values
			if floatValue > 1000 {
				format = "e"
			}

			// Use %g for very large float values
			if floatValue > 1000000 {
				format = "g"
			}

			// Generate C3D
			v.Generator.GenPrint(format, fmt.Sprintf("(float)%v", floatValue))
			v.Generator.GenPrint("c", "10")
			v.Generator.AddNewLine()

		} else if value.GetType() == values.StringType {
			// iterate over the string to conver to ascii code
			for _, char := range value.GetValue().(string) {
				// generate c3d
				v.Generator.GenPrint("c", fmt.Sprintf("%v", int(char)))
			}
			v.Generator.GenPrint("c", "10")
			v.Generator.AddNewLine()

		} else if value.GetType() == values.CharType {
			// generate c3d
			char := value.GetValue().(string)
			v.Generator.GenPrint("c", fmt.Sprintf("%v", int(char[0])))
			v.Generator.GenPrint("c", "10")
			v.Generator.AddNewLine()

		} else if value.GetType() == values.BooleanType {
			// generate c3d
			if value.GetValue().(bool) {
				v.Generator.GenPrint("d", "1")
			} else {
				v.Generator.GenPrint("d", "0")
			}
			v.Generator.GenPrint("c", "10")
			v.Generator.AddNewLine()

		} else if value.GetType() == values.NilType {
			// generate c3d
			v.Generator.GenPrint("f", "9999999827968.00")
			v.Generator.GenPrint("c", "10")
			v.Generator.AddNewLine()
		} else {

			// TODO: add error
			return nil
		}

	}

	return nil

}
