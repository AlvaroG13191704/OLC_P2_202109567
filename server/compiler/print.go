package compiler

import (
	"fmt"
	"server/compiler/compiler/values"
	"server/compiler/parser"
	"strconv"
)

// visit printstmt
func (v *Visitor) VisitPrintstmt(ctx *parser.PrintstmtContext) interface{} {

	exprList := ctx.ExprList().(*parser.ExprListContext)
	expressiongs := exprList.AllExpr()

	for _, expression := range expressiongs {
		fmt.Println("Status functionCode: ", v.Generator.FunctionCode, v.Generator.MainCode)
		// v.Generator.GenPrint("c", "10")

		primitive := v.Visit(expression).(*values.C3DPrimitive) // visit the expression

		//evalue the type of the value
		if primitive.GetType() == values.IntType {
			// generate c3d
			v.Generator.GenPrint("d", "(int)"+primitive.Value)
			v.Generator.GenPrint("c", "10")
			v.Generator.AddNewLine()

		} else if primitive.GetType() == values.FloatType {
			if primitive.IsTemp {
				v.Generator.GenPrint("f", primitive.Value)
			} else {

				digit := primitive.GetValue()
				// conver to float
				floatValue, _ := strconv.ParseFloat(digit, 64) // convert to float
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
			}
			v.Generator.GenPrint("c", "10")
			v.Generator.AddNewLine()

		} else if primitive.GetType() == values.StringType {
			// iterate over the string to conver to ascii code
			if !v.Generator.GeneratorNativeVariables.ConcatNative.IsFirstTimePrintString {
				v.Generator.GenerateFuncToPrint()
				v.Generator.GeneratorNativeVariables.ConcatNative.IsFirstTimePrintString = true
			}
			if primitive.IsTemp {
				v.Generator.GenPrintString(primitive.GetValue())

			} else {

				temp := v.Generator.GenString(primitive.GetValue())
				v.Generator.GenPrintString(temp)
			}
			// v.Generator.GenPrint("c", "10")
			// v.Generator.AddNewLine()

		} else if primitive.GetType() == values.CharType {
			// generate c3d
			char := primitive.GetValue()
			v.Generator.GenPrint("c", fmt.Sprintf("%v", int(char[0])))

		} else if primitive.GetType() == values.BooleanType {
			if !v.Generator.GeneratorNativeVariables.PrintBoolNative.IsBoolFunc {
				v.Generator.GenPrintBoolFunc()
				v.Generator.GeneratorNativeVariables.PrintBoolNative.IsBoolFunc = true
			}
			// generate c3d
			// assign temp to the tempBoolFunc
			v.Generator.AssignTemp(v.Generator.GeneratorNativeVariables.PrintBoolNative.TempBoolFunc, primitive.GetValue())
			// call the function
			v.Generator.Code = append(v.Generator.Code, "_print_bool();\n")
			// add newl ine
			v.Generator.GenPrint("c", "10")
			v.Generator.AddNewLine()

		} else if primitive.GetType() == values.NilType {
			// generate c3d
			v.Generator.GenPrint("c", "110")
			v.Generator.GenPrint("c", "105")
			v.Generator.GenPrint("c", "108")
			v.Generator.GenPrint("c", "10")
			v.Generator.AddNewLine()
		} else {

			// TODO: add error
			return nil
		}
		// v.Generator.GenPrint("c", "10")

	}
	v.Generator.GenPrint("c", "10")
	// v.Generator.AddNewLine()

	return nil

}
