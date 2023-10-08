package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

// visit expr
func (v *Visitor) VisitArithmeticOperationExpr(ctx *parser.ArithmeticOperationExprContext) interface{} {
	// get the left value
	leftValue := v.Visit(ctx.GetLeft()).(*values.C3DPrimitive)
	// get the right value
	rightValue := v.Visit(ctx.GetRight()).(*values.C3DPrimitive)

	sign := ctx.GetOp().GetText() // Get the operator

	// gen new temp
	newTemp := v.Generator.NewTemp()
	// add comment
	v.Generator.GenComment("Arithmetic operation")
	// for each sing evaluate the operation and verify if the operation is valid
	switch sign {
	case "+":
		if leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType { // Int + Int
			// fmt.Println("Int + Int")
			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "+")
			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.IntType, true)

		} else if leftValue.GetType() == values.IntType && rightValue.GetType() == values.FloatType { // Int + Float

			// fmt.Println("Int + Float")
			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "+")

			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.FloatType, true)

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.FloatType { // Float + Float

			// fmt.Println("Float + Float")

			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "+")

			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.FloatType, true)

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.IntType { // Float + Int

			// fmt.Println("Float + Int")

			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "+")

			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.FloatType, true)

		} else if leftValue.GetType() == values.StringType && rightValue.GetType() == values.StringType { // String + String

			fmt.Println("leftValue", leftValue, "rightValue", rightValue)
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
			if !v.Generator.GeneratorNativeVariables.ConcatNative.IsFirstTimeConcat {
				initPointer = v.Generator.ConcantStrings()
				v.Generator.GeneratorNativeVariables.ConcatNative.IsFirstTimeConcat = true
				// gen code to assign the pointer
				v.Generator.GenConcatString(tempLeft, tempRight)
				return values.NewC3DPrimitive(initPointer, newTemp, values.StringType, true)
			} else {
				// gen code to assign the pointer
				v.Generator.GenConcatString(tempLeft, tempRight)
				return values.NewC3DPrimitive(v.Generator.GeneratorNativeVariables.ConcatNative.TempInitConcat, newTemp, values.StringType, true)
			}

		} else {
			// throw an error
			log.Printf("Error: Invalid operation between '+' '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Invalid operation between '%s' and '%s'", leftValue.GetType(), rightValue.GetType()),
				Type:   "Semantic",
			})
		}

		// rest
	case "-":
		if leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType { // Int - Int

			// fmt.Println("Int - Int")
			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "-")

			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.IntType, true)

		} else if leftValue.GetType() == values.IntType && rightValue.GetType() == values.FloatType { // Int - Float

			// fmt.Println("Int - Float")
			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "-")

			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.FloatType, true)

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.FloatType { // Float - Float

			// fmt.Println("Float - Float")
			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "-")

			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.FloatType, true)

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.IntType { // Float - Int

			// fmt.Println("Float - Int")
			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "-")

			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.FloatType, true)

		} else {
			// throw an error
			log.Printf("Error: Invalid operation '-'  between '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Invalid operation between '%s' and '%s'", leftValue.GetType(), rightValue.GetType()),
				Type:   "Semantic",
			})
		}
		// mul
	case "*":
		if leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType { // Int * Int

			// fmt.Println("Int * Int")
			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "*")

			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.IntType, true)

		} else if leftValue.GetType() == values.IntType && rightValue.GetType() == values.FloatType { // Int * Float

			// fmt.Println("Int * Float")
			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "*")

			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.FloatType, true)

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.FloatType { // Float * Float

			// fmt.Println("Float * Float")
			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "*")

			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.FloatType, true)

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.IntType { // Float * Int

			// fmt.Println("Float * Int")
			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "*")

			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.FloatType, true)

		} else {
			// throw an error
			log.Printf("Error: Invalid operation '*' between '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Invalid operation between '%s' and '%s'", leftValue.GetType(), rightValue.GetType()),
				Type:   "Semantic",
			})
		}
		// div
	case "/":

		if leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType { // Int / Int

			// fmt.Println("Int / Int")
			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "/")

			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.IntType, true)

		} else if leftValue.GetType() == values.IntType && rightValue.GetType() == values.FloatType { // Int / Float

			// fmt.Println("Int / Float")
			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "/")

			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.FloatType, true)

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.FloatType { // Float / Float

			// fmt.Println("Float / Float")
			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "/")

			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.FloatType, true)

		} else if leftValue.GetType() == values.FloatType && rightValue.GetType() == values.IntType { // Float / Int

			// fmt.Println("Float / Int")
			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "/")

			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.FloatType, true)

		} else {
			// throw an error
			log.Printf("Error: Invalid operation '/' between '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Invalid operation between '%s' and '%s'", leftValue.GetType(), rightValue.GetType()),
				Type:   "Semantic",
			})
		}

	case "%":

		if leftValue.GetType() == values.IntType && rightValue.GetType() == values.IntType { // Int % Int

			// fmt.Println("Int % Int")
			// generate c3d arithmetic operation
			v.Generator.GenArithmetic(newTemp, leftValue.GetValue(), rightValue.GetValue(), "%")

			// return the value
			return values.NewC3DPrimitive(newTemp, newTemp, values.IntType, true)

		} else {
			// throw an error
			log.Printf("Error: Invalid operation '%%' between '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Invalid operation between '%s' and '%s'", leftValue.GetType(), rightValue.GetType()),
				Type:   "Semantic",
			})
		}

	default:
		// throw an error
		log.Printf("Error: Unknown operator: %s \n", sign)
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Unknown operator: %s", sign),
			Type:   "Semantic",
		})

	}

	return values.NewC3DPrimitive(values.NilType, newTemp, values.NilType, false)

}
