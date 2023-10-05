package compiler

import (
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

// visit idexpr
// func (v *Visitor) VisitIdExpr(ctx *parser.IdExprContext) interface{} {
// 	id := ctx.GetText() // get the id
// 	// fmt.Println("Id -> ", id)

// 	// verify if the id is in the scope or others
// 	variable, ok := v.VerifyScope(id)
// 	if ok {
// 		value := variable.(Symbol)
// 		// return the value
// 		return value.Value

// 	} else {
// 		fmt.Printf("Scope -> %v\n", v.SymbolStack)
// 		// add the error to the errors
// 		log.Printf("Error: Variable '%s' not declared", id)
// 		v.Errors = append(v.Errors, Error{
// 			Line:   ctx.GetStart().GetLine(),
// 			Column: ctx.GetStart().GetColumn(),
// 			Msg:    "Variable '" + id + "' not declared",
// 			Type:   "VariableError",
// 		})
// 	}
// 	return &values.Nil{
// 		Value: nil,
// 	}
// }

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

// // visit notexpr
// func (v *Visitor) VisitNotExpr(ctx *parser.NotExprContext) interface{} {
// 	// get the value
// 	value := v.Visit(ctx.Expr()).(values.PRIMITIVE)
// 	// verify if the value is a boolean
// 	if value.GetType() == values.BooleanType {
// 		return &values.Boolean{Value: !value.GetValue().(bool)} // return the negative of the value
// 	}
// 	return &values.Nil{
// 		Value: nil,
// 	}

// }
