package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

// VisitEmbeddedFunctionExpr
func (v *Visitor) VisitEmbeddedFunctionExpr(ctx *parser.EmbeddedFunctionExprContext) interface{} {
	// return the visit embbded function
	return v.Visit(ctx.EmbbededFunc())
}

// VisitIntstmt
func (v *Visitor) VisitIntstmt(ctx *parser.IntstmtContext) interface{} {
	var tempReturn string
	// get expresion
	expr := v.Visit(ctx.Expr()).(*values.C3DPrimitive)

	fmt.Println("expr", expr)

	// evaluate the type
	if (expr.GetType() == values.StringType || expr.GetType() == values.CharType) && expr.GetIsTemp() {

		// verify if it's first time
		if !v.Generator.GeneratorNativeVariables.EmbebbedNative.IsEmbebbedFunc1 {
			// set true
			v.Generator.GeneratorNativeVariables.EmbebbedNative.IsEmbebbedFunc1 = true
			// generate the function
			tempReturn = v.Generator.GenNativeIntFloat()
			v.Generator.GeneratorNativeVariables.EmbebbedNative.TempReturnValue1 = tempReturn
			v.Generator.CallFunctionIntFloat(expr.GetValue())

			return values.NewC3DPrimitive(tempReturn, tempReturn, values.IntType, true)
		}

		// call the function
		v.Generator.CallFunctionIntFloat(expr.GetValue())

		return values.NewC3DPrimitive(v.Generator.GeneratorNativeVariables.EmbebbedNative.TempReturnValue1, tempReturn, values.IntType, true)

	} else if (expr.GetType() == values.StringType || expr.GetType() == values.CharType) && !expr.GetIsTemp() {
		// generate sring
		tempString := v.Generator.GenString(expr.GetValue())

		// verify if it's first time
		if !v.Generator.GeneratorNativeVariables.EmbebbedNative.IsEmbebbedFunc1 {
			// set true
			v.Generator.GeneratorNativeVariables.EmbebbedNative.IsEmbebbedFunc1 = true
			// generate the function
			tempReturn = v.Generator.GenNativeIntFloat()
			v.Generator.GeneratorNativeVariables.EmbebbedNative.TempReturnValue1 = tempReturn

			v.Generator.CallFunctionIntFloat(tempString)

			return values.NewC3DPrimitive(tempReturn, tempReturn, values.IntType, true)
		}

		// call the function
		v.Generator.CallFunctionIntFloat(tempString)

		return values.NewC3DPrimitive(v.Generator.GeneratorNativeVariables.EmbebbedNative.TempReturnValue1, tempReturn, values.IntType, false)

	} else if expr.GetType() == values.FloatType || expr.GetType() == values.IntType {

		// gen temp
		temp := v.Generator.NewTemp()
		// assign temp
		v.Generator.AssignTemp(temp, "(int)"+expr.GetValue())

		return values.NewC3DPrimitive(temp, temp, values.IntType, true)

	} else {
		log.Printf("Error: Invalid type '%s' to convert to integer \n", expr.GetType())
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Invalid type '%s' to convert to integer", expr.GetType()),
			Type:   "Semantic",
		})
		values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
	}

	return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)

}

// VisitFloatstmt
func (v *Visitor) VisitFloatstmt(ctx *parser.FloatstmtContext) interface{} {
	var tempReturn string
	// get expresion
	expr := v.Visit(ctx.Expr()).(*values.C3DPrimitive)

	fmt.Println("expr", expr)

	// evaluate the type
	if expr.GetType() == values.StringType && expr.GetIsTemp() {

		// verify if it's first time
		if !v.Generator.GeneratorNativeVariables.EmbebbedNative.IsEmbebbedFunc1 {
			// set true
			v.Generator.GeneratorNativeVariables.EmbebbedNative.IsEmbebbedFunc1 = true
			// generate the function
			tempReturn = v.Generator.GenNativeIntFloat()
			v.Generator.GeneratorNativeVariables.EmbebbedNative.TempReturnValue1 = tempReturn
			v.Generator.CallFunctionIntFloat(expr.GetValue())

			return values.NewC3DPrimitive(tempReturn, tempReturn, values.FloatType, true)
		}

		// call the function
		v.Generator.CallFunctionIntFloat(expr.GetValue())

		return values.NewC3DPrimitive(v.Generator.GeneratorNativeVariables.EmbebbedNative.TempReturnValue1, tempReturn, values.FloatType, true)

	} else if expr.GetType() == values.StringType && !expr.GetIsTemp() {
		// generate sring
		tempString := v.Generator.GenString(expr.GetValue())

		// verify if it's first time
		if !v.Generator.GeneratorNativeVariables.EmbebbedNative.IsEmbebbedFunc1 {
			// set true
			v.Generator.GeneratorNativeVariables.EmbebbedNative.IsEmbebbedFunc1 = true
			// generate the function
			tempReturn = v.Generator.GenNativeIntFloat()
			v.Generator.GeneratorNativeVariables.EmbebbedNative.TempReturnValue1 = tempReturn

			v.Generator.CallFunctionIntFloat(tempString)

			return values.NewC3DPrimitive(tempReturn, tempReturn, values.FloatType, true)
		}

		// call the function
		v.Generator.CallFunctionIntFloat(tempString)

		return values.NewC3DPrimitive(v.Generator.GeneratorNativeVariables.EmbebbedNative.TempReturnValue1, tempReturn, values.FloatType, false)

	} else if expr.GetType() == values.FloatType || expr.GetType() == values.IntType {

		// gen temp
		temp := v.Generator.NewTemp()
		// assign temp
		v.Generator.AssignTemp(temp, "(float)"+expr.GetValue())

		return values.NewC3DPrimitive(temp, temp, values.FloatType, true)

	} else {
		log.Printf("Error: Invalid type '%s' to convert to integer \n", expr.GetType())
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Invalid type '%s' to convert to integer", expr.GetType()),
			Type:   "Semantic",
		})
		values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
	}

	return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
}

// VisitStringstmt
func (v *Visitor) VisitStringstmt(ctx *parser.StringstmtContext) interface{} {
	var tempReturn string
	// get expresion
	expr := v.Visit(ctx.Expr()).(*values.C3DPrimitive)

	// evaluate the type
	if expr.GetType() == values.StringType {

		return &values.String{Value: expr.GetValue()}

	} else if expr.GetType() == values.IntType || expr.GetType() == values.FloatType {
		// convert the int to string
		// verify if it's first time
		if !v.Generator.GeneratorNativeVariables.EmbebbedNative.IsEmbebbedFunc2 {
			// set true
			v.Generator.GeneratorNativeVariables.EmbebbedNative.IsEmbebbedFunc2 = true
			// generate the function
			tempReturn = v.Generator.GenNativeDigitsToString()
			v.Generator.GeneratorNativeVariables.EmbebbedNative.TempReturnValue2 = tempReturn
		}

		tempReturn = v.Generator.CallFunctionString(expr.GetValue())

		return values.NewC3DPrimitive(tempReturn, tempReturn, values.StringType, true)

	} else if expr.GetType() == values.BooleanType {
		// TODO: only work for primitives
		var tempHeap string
		fmt.Println("expr", expr)
		// convert the boolean to string
		if expr.GetValue() == "1" {
			tempHeap = v.Generator.GenString("true")
		} else {
			tempHeap = v.Generator.GenString("false")
		}

		return values.NewC3DPrimitive(tempHeap, tempHeap, values.StringType, true)

	} else {
		log.Printf("Error: Invalid type '%s' to convert to string \n", expr.GetType())
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Invalid type '%s' to convert to string", expr.GetType()),
			Type:   "Semantic",
		})
		return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
	}

}
