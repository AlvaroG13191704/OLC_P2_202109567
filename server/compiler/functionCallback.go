package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

// CallFunctionExpr
func (v *Visitor) VisitCallFunctionExpr(ctx *parser.CallFunctionExprContext) interface{} {
	valueExpr := v.Visit(ctx.CallFunctionStmt()).(*values.C3DPrimitive)

	return valueExpr

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

	// create temp
	temp := v.Generator.NewTemp()
	// assign the P to the temp
	v.Generator.AssignTemp(temp, "P")

	// add the function
	v.Generator.Code = append(v.Generator.Code, fmt.Sprintf("%s();\n", function.Id))

	var tempReturn string
	// verifiy if there is a return value
	if function.ReturnType != "void" {
		// create temp
		tempReturn = v.Generator.NewTemp()
		// assign the return value to the temp
		v.Generator.AssignTemp(tempReturn, v.ReturnTemp)

	}

	return values.NewC3DPrimitive(tempReturn, function.ReturnTemp, function.ReturnType, true)

}
