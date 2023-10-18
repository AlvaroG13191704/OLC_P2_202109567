package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

// visit idexpr
func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
	id := ctx.GetText() // get the id
	// fmt.Println("Id -> ", id)

	// verify if the id is in the scope or others
	variable, ok := v.VerifyScope(id)
	if ok {
		value := variable.(Symbol)

		if value.Value.(*values.C3DPrimitive).GetType() == values.NilType {
			return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
		}
		// t0 = stack[(int)0 ]
		temp := v.Generator.NewTemp()
		v.Generator.AccessStack(temp, fmt.Sprintf("%d", value.StackDirection))
		// change temp
		return values.NewC3DPrimitive(temp, temp, value.TypeData, true)

	} else {
		fmt.Printf("Scope -> %v\n", v.SymbolStack)
		// add the error to the errors
		log.Printf("Error: Variable '%s' not declared", id)
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Variable '" + id + "' not declared",
			Type:   "VariableError",
		})
	}
	return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
}

// visit parenexpr
func (v *Visitor) VisitParenExpr(ctx *parser.ParenExprContext) interface{} {
	return v.Visit(ctx.Expr()) // visit the expression
}

// visit negativeexpr
func (v *Visitor) VisitNegExpr(ctx *parser.NegExprContext) interface{} {
	// get the value
	value := v.Visit(ctx.Expr()).(*values.C3DPrimitive)
	// verify if the value is an integer or a float
	if value.GetType() == values.IntType {

		// add a temp
		temp := v.Generator.NewTemp()
		// add the code
		v.Generator.GenArithmetic(temp, "0", value.GetValue(), "-")
		return values.NewC3DPrimitive(temp, "", values.IntType, false)

	} else if value.GetType() == values.FloatType {

		temp := v.Generator.NewTemp()
		v.Generator.GenArithmetic(temp, "0", value.GetValue(), "-")
		return values.NewC3DPrimitive(temp, "", values.FloatType, false)

	}

	return values.NewC3DPrimitive(values.NilType, "", values.NilType, false)

}

// visit notexpr
func (v *Visitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
	// get the value
	value := v.Visit(ctx.Expr()).(*values.C3DPrimitive)

	if !v.Generator.GeneratorNativeVariables.LogicalNative.IsNotFunc {
		v.Generator.GenNotFunc()
		v.Generator.GeneratorNativeVariables.LogicalNative.IsNotFunc = true
		temp := v.Generator.GenLogical(value.GetValue(), "", v.Generator.GeneratorNativeVariables.LogicalNative.TempNot, "", v.Generator.GeneratorNativeVariables.LogicalNative.TempNoPointer, "not")
		// return the value
		return values.NewC3DPrimitive(temp, temp, values.BooleanType, true)
	} else {
		temp := v.Generator.GenLogical(value.GetValue(), "", v.Generator.GeneratorNativeVariables.LogicalNative.TempNot, "", v.Generator.GeneratorNativeVariables.LogicalNative.TempNoPointer, "not")
		// return the value
		return values.NewC3DPrimitive(temp, temp, values.BooleanType, true)
	}

	// return values.NewC3DPrimitive(values.NilType, "", values.NilType, false)
}

// VisitCallBackExpr -> Visit a parse tree produced by SFGrammarParser#CallBackExpr.
func (v *Visitor) VisitCallBackExpr(ctx *parser.CallBackExprContext) interface{} {

	return v.Visit(ctx.CallBack())
}
