package compiler

import (
	"fmt"
	"log"
	"server/compiler/parser"
)

// CallFunctionExpr
func (v *Visitor) VisitCallFunctionExpr(ctx *parser.CallFunctionExprContext) interface{} {
	v.Visit(ctx.CallFunctionStmt())
	// visit the expression
	// // change the return value
	// v.IsReturn = false
	// fmt.Println("VisitCallFunctionExpr", v.ReturnValue)

	// return v.ReturnValue
	return nil
}

// VisitCallFunctionWithoutParams
func (v *Visitor) VisitCallFunctionWithoutParams(ctx *parser.CallFunctionWithoutParamsContext) interface{} {

	functionName := ctx.ID_PRIMITIVE().GetText()

	// get function from the scope
	function, ok := v.GetFunction(functionName)
	if !ok {
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("function %s not found", functionName),
			Type:   "Semantic",
		})
		log.Printf("function %s not found", functionName)
		return nil
	}

	// verifiy if the return value type is the same as the function return type
	if function.ReturnType == "void" {
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("function '%s' is void", functionName),
			Type:   "Semantic",
		})
		log.Printf("function '%s' is void", functionName)
		return nil
	} else {
		// execute the function
		// v.Visit(function.Value.(*parser.BlockContext))

		// // evaluate if the return type is the same as the function return type
		// if function.TypeVariable != v.ReturnValue.(values.PRIMITIVE).GetType() {
		// 	// add error
		// 	v.Errors = append(v.Errors, Error{
		// 		Line:   ctx.GetStart().GetLine(),
		// 		Column: ctx.GetStart().GetColumn(),
		// 		Msg:    fmt.Sprintf("function %s return type is %s, expected %s", functionName, function.TypeVariable, v.ReturnValue.(values.PRIMITIVE).GetType()),
		// 		Type:   "Semantic",
		// 	})
		// 	log.Printf("function %s return type is %s, expected %s", functionName, function.TypeVariable, v.ReturnValue.(values.PRIMITIVE).GetType())
		// 	return nil
		// }

		return nil
	}

}
