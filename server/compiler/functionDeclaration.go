package compiler

import (
	"fmt"
	"log"
	"server/compiler/parser"
)

func (v *Visitor) VisitFunctionWithoutParams(ctx *parser.FunctionWithoutParamsContext) interface{} {
	// get the id
	idFunction := ctx.ID_PRIMITIVE().GetText()
	// verify if the function is already declared
	if v.VerifyFunctionScope(idFunction) {
		log.Printf("Function %s already declared \n", idFunction)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Function '%s' already declared", idFunction),
			Type:   "Semantic",
		})
		return nil
	}

	// work with functions -> without params and return
	// set false
	v.Generator.FunctionCode = true
	// gen c3d
	v.Generator.FuncionUser = append(v.Generator.FuncionUser, fmt.Sprintf("void %s(){\n", idFunction))
	// visit block
	v.Visit(ctx.Block())
	// set false
	v.Generator.FunctionCode = false
	// gen c3d
	v.Generator.FuncionUser = append(v.Generator.FuncionUser, "return;\n")
	// close function
	v.Generator.FuncionUser = append(v.Generator.FuncionUser, "}\n")

	// // get the return type if it exists
	// var returnType string
	// if ctx.ARROW_FUNCTION() != nil {
	// 	returnType = ctx.Type_().GetText()
	// 	// save the function in the symbol table
	// 	symbol := Symbol{
	// 		Id:           idFunction,
	// 		TypeSymbol:   values.Function_Type,
	// 		TypeVariable: returnType,
	// 		TypeData:     values.Function_Type,
	// 		Value:        block,
	// 		Line:         ctx.GetStart().GetLine(),
	// 		Column:       ctx.GetStart().GetColumn(),
	// 	}

	// 	//
	// 	v.getCurrentScope()[idFunction] = symbol

	// 	v.TableSymbol = append(v.TableSymbol, symbol)

	// } else {
	// 	returnType = "void"
	// 	// save the function in the symbol table
	// 	symbol := SymbolTable{
	// 		Id:           idFunction,
	// 		TypeSymbol:   values.Type_Function,
	// 		TypeVariable: returnType,
	// 		TypeData:     values.Type_Function,
	// 		Value:        block,
	// 		Line:         ctx.GetStart().GetLine(),
	// 		Column:       ctx.GetStart().GetColumn(),
	// 	}

	// 	v.getCurrentScope()[idFunction] = symbol

	// 	v.TableSymbol = append(v.TableSymbol, symbol)
	// }

	// fmt.Println("FunctionWithoutParams: ", idFunction, returnType)

	return nil
}
