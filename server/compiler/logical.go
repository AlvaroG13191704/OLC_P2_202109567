package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

func (v *Visitor) VisitLogicalOperationExpr(ctx *parser.LogicalOperationExprContext) interface{} {
	// get the left value
	leftValue := v.Visit(ctx.GetLeft()).(*values.C3DPrimitive)
	// get the right value
	rightValue := v.Visit(ctx.GetRight()).(*values.C3DPrimitive)

	sign := ctx.GetOp().GetText() // Get the operator

	// add comment
	v.Generator.GenComment("Logical operation")

	if !v.Generator.GeneratorNativeVariables.PrintBoolNative.IsBoolFunc {
		v.Generator.GenPrintBoolFunc()
		v.Generator.GeneratorNativeVariables.PrintBoolNative.IsBoolFunc = true
	}

	// verify the type of the values
	switch sign {
	case "&&":
		if leftValue.GetType() == values.BooleanType && rightValue.GetType() == values.BooleanType { // Boolean && Boolean

			fmt.Println("AND")
			fmt.Println("Left: ", leftValue.GetValue(), "Right: ", rightValue.GetValue())

			if !v.Generator.GeneratorNativeVariables.LogicalNative.IsAndFunc {
				v.Generator.GenAndFunc()
				v.Generator.GeneratorNativeVariables.LogicalNative.IsAndFunc = true
			}
			temp := v.Generator.GenLogical(
				leftValue.GetValue(),
				rightValue.GetValue(),
				v.Generator.GeneratorNativeVariables.LogicalNative.Temp1And,
				v.Generator.GeneratorNativeVariables.LogicalNative.Temp2And,
				v.Generator.GeneratorNativeVariables.LogicalNative.TempAndPointer, "and")
			// return the value
			return values.NewC3DPrimitive(temp, temp, values.BooleanType, true)

			// if !v.Generator.GeneratorNativeVariables.LogicalNative.IsAndFunc {
			// 	v.Generator.GenAndFunc()
			// 	v.Generator.GeneratorNativeVariables.LogicalNative.IsAndFunc = true
			// 	temp := v.Generator.GenLogical(
			// 		leftValue.GetValue(),
			// 		rightValue.GetValue(),
			// 		v.Generator.GeneratorNativeVariables.LogicalNative.Temp1And,
			// 		v.Generator.GeneratorNativeVariables.LogicalNative.Temp2And,
			// 		v.Generator.GeneratorNativeVariables.LogicalNative.TempAndPointer, "and")
			// 	// return the value
			// 	return values.NewC3DPrimitive(temp, temp, values.BooleanType, true)
			// } else {
			// 	temp := v.Generator.GenLogical(
			// 		leftValue.GetValue(),
			// 		rightValue.GetValue(),
			// 		v.Generator.GeneratorNativeVariables.LogicalNative.Temp1And,
			// 		v.Generator.GeneratorNativeVariables.LogicalNative.Temp2And,
			// 		v.Generator.GeneratorNativeVariables.LogicalNative.TempAndPointer, "and")
			// 	// return the value
			// 	return values.NewC3DPrimitive(temp, temp, values.BooleanType, true)
			// }

		} else {

			// error
			log.Printf("Error: Invalid operation between '&&' '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Invalid operation between '&&'",
				Type:   "Semantic",
			})

			return &values.Nil{
				Value: nil,
			}

		}
	case "||":
		if leftValue.GetType() == values.BooleanType && rightValue.GetType() == values.BooleanType { // Boolean || Boolean
			if !v.Generator.GeneratorNativeVariables.LogicalNative.IsOrFunc {
				v.Generator.GenOrFunc()
				v.Generator.GeneratorNativeVariables.LogicalNative.IsOrFunc = true
			}
			temp := v.Generator.GenLogical(
				leftValue.GetValue(),
				rightValue.GetValue(),
				v.Generator.GeneratorNativeVariables.LogicalNative.Temp1Or,
				v.Generator.GeneratorNativeVariables.LogicalNative.Temp2Or,
				v.Generator.GeneratorNativeVariables.LogicalNative.TempOrPointer, "or")
			// return the value
			return values.NewC3DPrimitive(temp, temp, values.BooleanType, true)

			// if !v.Generator.GeneratorNativeVariables.LogicalNative.IsOrFunc {
			// 	v.Generator.GenOrFunc()
			// 	v.Generator.GeneratorNativeVariables.LogicalNative.IsOrFunc = true
			// 	temp := v.Generator.GenLogical(
			// 		leftValue.GetValue(),
			// 		rightValue.GetValue(),
			// 		v.Generator.GeneratorNativeVariables.LogicalNative.Temp1Or,
			// 		v.Generator.GeneratorNativeVariables.LogicalNative.Temp2Or,
			// 		v.Generator.GeneratorNativeVariables.LogicalNative.TempOrPointer, "or")
			// 	// return the value
			// 	return values.NewC3DPrimitive(temp, temp, values.BooleanType, true)
			// } else {
			// 	temp := v.Generator.GenLogical(
			// 		leftValue.GetValue(),
			// 		rightValue.GetValue(),
			// 		v.Generator.GeneratorNativeVariables.LogicalNative.Temp1Or,
			// 		v.Generator.GeneratorNativeVariables.LogicalNative.Temp2Or,
			// 		v.Generator.GeneratorNativeVariables.LogicalNative.TempOrPointer, "or")
			// 	// return the value
			// 	return values.NewC3DPrimitive(temp, temp, values.BooleanType, true)
			// }
		} else {

			// error
			log.Printf("Error: Invalid operation between '||' '%s' and '%s' \n", leftValue.GetType(), rightValue.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Invalid operation between '||'",
				Type:   "Semantic",
			})

			return &values.Nil{
				Value: nil,
			}

		}

	}

	return &values.Nil{
		Value: nil,
	}
}
