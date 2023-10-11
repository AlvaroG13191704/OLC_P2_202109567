package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

func (v *Visitor) VisitRelationalOperationExpr(ctx *parser.RelationalOperationExprContext) interface{} {
	// get the left value
	leftValue := v.Visit(ctx.GetLeft()).(*values.C3DPrimitive)
	// get the right value
	rightValue := v.Visit(ctx.GetRight()).(*values.C3DPrimitive)

	sign := ctx.GetOp().GetText() // Get the operator

	// gen new temp
	newTemp := v.Generator.NewTemp()
	// add comment
	v.Generator.GenComment("Comparation operation")

	if !v.Generator.GeneratorNativeVariables.PrintBoolNative.IsBoolFunc {
		v.Generator.GenPrintBoolFunc()
		v.Generator.GeneratorNativeVariables.PrintBoolNative.IsBoolFunc = true
	}

	switch sign {
	case "<":
		if (leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType) || // Int < Int
			(leftValue.GetType() == values.FloatType && rightValue.GetType() == values.FloatType) || // Float < Float
			(leftValue.GetType() == values.CharType && rightValue.GetType() == values.CharType) { // Char < Char

			if !leftValue.IsTemp || !rightValue.IsTemp {
				// generate c3d arithmetic operation
				temp := v.Generator.GenComparation(leftValue.GetValue(), rightValue.GetValue(), "<")
				// return the value
				return values.NewC3DPrimitive(temp, temp, values.BooleanType, true)

			} else {

				// generate c3d arithmetic operation
				temp := v.Generator.GenComparation(leftValue.GetValue(), rightValue.GetValue(), "<")
				// return the value
				return values.NewC3DPrimitive(temp, temp, values.BooleanType, true)

			}

		} else if leftValue.GetType() == values.StringType && rightValue.GetType() == values.StringType { // String < String

			// pending
			var tempLeft string
			var tempRight string
			var initPointer string
			// generate c3d to concat strings
			if !leftValue.IsTemp {
				// gen left
				tempLeft = v.Generator.GenString(leftValue.GetValue())
			} else {
				tempLeft = leftValue.GetValue()
			}

			if !rightValue.IsTemp {
				// gen right
				tempRight = v.Generator.GenString(rightValue.GetValue())
			} else {
				tempRight = rightValue.GetValue()
			}
			// gen code to concat
			if !v.Generator.GeneratorNativeVariables.CompareStringsNative.IsCompareStringFuncGL {
				initPointer = v.Generator.GenComparationStringFuncGL()
				v.Generator.GeneratorNativeVariables.CompareStringsNative.IsCompareStringFuncGL = true
				// gen code to assign the pointer
				v.Generator.GenComparationStringGL(tempLeft, tempRight)
				temp := v.Generator.GenComparation(initPointer, "1", "==")
				fmt.Println("tempLeft", tempLeft, "tempRight", tempRight)
				return values.NewC3DPrimitive(temp, initPointer, values.BooleanType, true)
			} else {
				// gen code to assign the pointer
				v.Generator.GenComparationStringGL(tempLeft, tempRight)
				temp := v.Generator.GenComparation(initPointer, "1", "==")
				fmt.Println("tempLeft", tempLeft, "tempRight", tempRight)
				return values.NewC3DPrimitive(temp, newTemp, values.BooleanType, true)
			}

		} else {
			// error
			log.Printf("Error: Invalid operation between '<' '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Invalid operation between '<'",
				Type:   "Semantic",
			})

			return values.NewC3DPrimitive(values.NilType, newTemp, values.NilType, false)
		}
	case "<=":
		if (leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType) || // Int <= Int
			(leftValue.GetType() == values.FloatType && rightValue.GetType() == values.FloatType) || // Float <= Float
			(leftValue.GetType() == values.CharType && rightValue.GetType() == values.CharType) { // Char <= Char

			if !leftValue.IsTemp || !rightValue.IsTemp {
				// generate c3d arithmetic operation
				temp := v.Generator.GenComparation(leftValue.GetValue(), rightValue.GetValue(), "<=")
				// return the value
				return values.NewC3DPrimitive(temp, temp, values.BooleanType, true)

			} else {

				// generate c3d arithmetic operation
				temp := v.Generator.GenComparation(leftValue.GetValue(), rightValue.GetValue(), "<=")
				// return the value
				return values.NewC3DPrimitive(temp, temp, values.BooleanType, true)

			}

		} else if leftValue.GetType() == values.StringType && rightValue.GetType() == values.StringType { // String <= String
			// pending
			var tempLeft string
			var tempRight string
			var initPointer string
			// generate c3d to concat strings
			if !leftValue.IsTemp {
				// gen left
				tempLeft = v.Generator.GenString(leftValue.GetValue())
			} else {
				tempLeft = leftValue.GetValue()
			}

			if !rightValue.IsTemp {
				// gen right
				tempRight = v.Generator.GenString(rightValue.GetValue())
			} else {
				tempRight = rightValue.GetValue()
			}
			// gen code to concat
			if !v.Generator.GeneratorNativeVariables.CompareStringsNative.IsCompareStringFuncGLE {
				initPointer = v.Generator.GenComparationStringFuncGLE()
				v.Generator.GeneratorNativeVariables.CompareStringsNative.IsCompareStringFuncGLE = true
				// gen code to assign the pointer
				v.Generator.GenComparationStringGLE(tempLeft, tempRight)
				temp := v.Generator.GenComparation(initPointer, "1", "==")
				fmt.Println("tempLeft", tempLeft, "tempRight", tempRight)
				return values.NewC3DPrimitive(temp, initPointer, values.BooleanType, true)
			} else {
				// gen code to assign the pointer
				v.Generator.GenComparationStringGLE(tempLeft, tempRight)
				temp := v.Generator.GenComparation(initPointer, "1", "==")
				fmt.Println("tempLeft", tempLeft, "tempRight", tempRight)
				return values.NewC3DPrimitive(temp, newTemp, values.BooleanType, true)
			}

		} else {

			// error
			log.Printf("Error: Invalid operation between '<=' '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Invalid operation between '<='",
				Type:   "Semantic",
			})

			return values.NewC3DPrimitive(values.NilType, newTemp, values.NilType, false)

		}
	case ">":
		if (leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType) || // Int > Int
			(leftValue.GetType() == values.FloatType && rightValue.GetType() == values.FloatType) || // Float > Float
			(leftValue.GetType() == values.CharType && rightValue.GetType() == values.CharType) { // Char > Char

			if !leftValue.IsTemp || !rightValue.IsTemp {
				// generate c3d arithmetic operation
				temp := v.Generator.GenComparation(leftValue.GetValue(), rightValue.GetValue(), ">")
				// return the value
				return values.NewC3DPrimitive(temp, temp, values.BooleanType, true)

			} else {

				// generate c3d arithmetic operation
				temp := v.Generator.GenComparation(leftValue.GetValue(), rightValue.GetValue(), ">")
				// return the value
				return values.NewC3DPrimitive(temp, temp, values.BooleanType, true)

			}

		} else if leftValue.GetType() == values.StringType && rightValue.GetType() == values.StringType { // String > String

			// pending
			var tempLeft string
			var tempRight string
			var initPointer string
			// generate c3d to concat strings
			if !leftValue.IsTemp {
				// gen left
				tempLeft = v.Generator.GenString(leftValue.GetValue())
			} else {
				tempLeft = leftValue.GetValue()
			}

			if !rightValue.IsTemp {
				// gen right
				tempRight = v.Generator.GenString(rightValue.GetValue())
			} else {
				tempRight = rightValue.GetValue()
			}
			// gen code to concat
			if !v.Generator.GeneratorNativeVariables.CompareStringsNative.IsCompareStringFuncGL {
				initPointer = v.Generator.GenComparationStringFuncGL()
				v.Generator.GeneratorNativeVariables.CompareStringsNative.IsCompareStringFuncGL = true
				// gen code to assign the pointer
				v.Generator.GenComparationStringGL(tempLeft, tempRight)
				temp := v.Generator.GenComparation(initPointer, "1", "==")
				fmt.Println("tempLeft", tempLeft, "tempRight", tempRight)
				return values.NewC3DPrimitive(temp, initPointer, values.BooleanType, true)
			} else {
				// gen code to assign the pointer
				v.Generator.GenComparationStringGL(tempLeft, tempRight)
				temp := v.Generator.GenComparation(initPointer, "1", "==")
				fmt.Println("tempLeft", tempLeft, "tempRight", tempRight)
				return values.NewC3DPrimitive(temp, newTemp, values.BooleanType, true)
			}

		} else {

			// error
			log.Printf("Error: Invalid operation between '>' '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Invalid operation between '>'",
				Type:   "Semantic",
			})

			return values.NewC3DPrimitive(values.NilType, newTemp, values.NilType, false)
		}

	case ">=":
		if (leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType) || // Int >= Int
			(leftValue.GetType() == values.FloatType && rightValue.GetType() == values.FloatType) || // Float >= Float
			(leftValue.GetType() == values.CharType && rightValue.GetType() == values.CharType) { // Char >= Char

			if !leftValue.IsTemp || !rightValue.IsTemp {
				// generate c3d arithmetic operation
				temp := v.Generator.GenComparation(leftValue.GetValue(), rightValue.GetValue(), ">=")
				// return the value
				return values.NewC3DPrimitive(temp, temp, values.BooleanType, true)

			} else {

				// generate c3d arithmetic operation
				temp := v.Generator.GenComparation(leftValue.GetValue(), rightValue.GetValue(), ">=")
				// return the value
				return values.NewC3DPrimitive(temp, temp, values.BooleanType, true)

			}

		} else if leftValue.GetType() == values.StringType && rightValue.GetType() == values.StringType { // String >= String

			// pending
			var tempLeft string
			var tempRight string
			var initPointer string
			// generate c3d to concat strings
			if !leftValue.IsTemp {
				// gen left
				tempLeft = v.Generator.GenString(leftValue.GetValue())
			} else {
				tempLeft = leftValue.GetValue()
			}

			if !rightValue.IsTemp {
				// gen right
				tempRight = v.Generator.GenString(rightValue.GetValue())
			} else {
				tempRight = rightValue.GetValue()
			}
			// gen code to concat
			if !v.Generator.GeneratorNativeVariables.CompareStringsNative.IsCompareStringFuncGLE {
				initPointer = v.Generator.GenComparationStringFuncGLE()
				v.Generator.GeneratorNativeVariables.CompareStringsNative.IsCompareStringFuncGLE = true
				// gen code to assign the pointer
				v.Generator.GenComparationStringGLE(tempLeft, tempRight)
				temp := v.Generator.GenComparation(initPointer, "1", "==")
				fmt.Println("tempLeft", tempLeft, "tempRight", tempRight)
				return values.NewC3DPrimitive(temp, initPointer, values.BooleanType, true)
			} else {
				// gen code to assign the pointer
				v.Generator.GenComparationStringGLE(tempLeft, tempRight)
				temp := v.Generator.GenComparation(initPointer, "1", "==")
				fmt.Println("tempLeft", tempLeft, "tempRight", tempRight)
				return values.NewC3DPrimitive(temp, newTemp, values.BooleanType, true)
			}
		} else {

			// error
			log.Printf("Error: Invalid operation between '>=' '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Invalid operation between '>='",
				Type:   "Semantic",
			})

			return &values.Nil{
				Value: nil,
			}
		}
	}

	return values.NewC3DPrimitive(values.NilType, newTemp, values.NilType, false)
}
