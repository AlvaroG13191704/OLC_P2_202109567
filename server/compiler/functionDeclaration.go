package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

func (v *Visitor) VisitFunctionWithoutParams(ctx *parser.FunctionWithoutParamsContext) interface{} {
	// push a new function context
	v.PushFunctionContext("function")
	defer v.PopFunctionContext()
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

	// create temps to return and save the label
	if ctx.ARROW_FUNCTION() != nil {
		v.ReturnLabel = v.Generator.NewLabel()
		v.ReturnTemp = v.Generator.NewTemp()
	}
	// work with functions -> without params and return
	v.SizeFunction = 0 // restart the size of the function
	// set false
	v.Generator.FunctionCode = true
	// gen c3d
	v.Generator.FuncionUser = append(v.Generator.FuncionUser, fmt.Sprintf("void %s(){\n", idFunction))
	// visit block
	v.Visit(ctx.Block())

	// add return if there is return value
	if ctx.ARROW_FUNCTION() != nil {
		v.Generator.AddLabel(v.ReturnLabel)
	}
	// set false
	v.Generator.FunctionCode = false

	// rest
	if ctx.ARROW_FUNCTION() != nil {
		v.Generator.StackCounter = v.SizeFunction
	} else {
		v.Generator.StackCounter = v.SizeFunction - 2
	}
	// gen c3d
	v.Generator.FuncionUser = append(v.Generator.FuncionUser, "return;\n")
	// close function
	v.Generator.FuncionUser = append(v.Generator.FuncionUser, "}\n")

	// get the return type if it exists
	var returnType string
	if ctx.ARROW_FUNCTION() != nil {

		returnType = ctx.Type_().GetText()
		// save the function in the symbol table
		symbol := Symbol{
			Id:           idFunction,
			TypeSymbol:   values.Function_Type,
			TypeMutable:  returnType,
			TypeData:     values.Function_Type,
			ReturnType:   returnType,
			ReturnTemp:   v.ReturnTemp,
			ReturnLabel:  v.ReturnLabel,
			sizeFunction: v.SizeFunction,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}

		//
		v.getCurrentScope()[idFunction] = symbol

		v.TableSymbol = append(v.TableSymbol, symbol)

	} else {
		returnType = "void"
		// save the function in the symbol table
		symbol := Symbol{
			Id:           idFunction,
			TypeSymbol:   values.Function_Type,
			TypeMutable:  returnType,
			TypeData:     values.Function_Type,
			ReturnType:   returnType,
			sizeFunction: v.SizeFunction,
			ReturnTemp:   v.ReturnTemp,
			ReturnLabel:  v.ReturnLabel,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}

		v.getCurrentScope()[idFunction] = symbol

		v.TableSymbol = append(v.TableSymbol, symbol)
	}

	// fmt.Println("FunctionWithoutParams: ", idFunction, returnType)

	return nil
}

// VisitFunctionWithParams
func (v *Visitor) VisitFunctionWithParams(ctx *parser.FunctionWithParamsContext) interface{} {
	v.PushFunctionContext("function")
	defer v.PopFunctionContext()
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
	v.SizeFunction = 0 // restart the size of the function
	// create temps to return and save the label
	if ctx.ARROW_FUNCTION() != nil {
		v.ReturnLabel = v.Generator.NewLabel()
		v.ReturnTemp = v.Generator.NewTemp()
	}
	// set false
	v.Generator.FunctionCode = true
	// gen c3d
	v.Generator.FuncionUser = append(v.Generator.FuncionUser, fmt.Sprintf("void %s(){\n", idFunction))

	// get params
	// get list of params
	params := v.Visit(ctx.ListFunctionParams()).(map[string][]Symbol)

	// visit block
	v.Visit(ctx.Block())

	// delete from the scope
	for _, param := range params["internal"] {
		v.DeleteVariableGlobalScope(param.Id)
	}

	// add return if there is return value
	if ctx.ARROW_FUNCTION() != nil {
		v.Generator.AddLabel(v.ReturnLabel)
	}

	fmt.Println("SizeFunction: ", v.SizeFunction)
	fmt.Println("StackCounter: ", v.Generator.StackCounter)

	// rest
	if ctx.ARROW_FUNCTION() != nil {
		v.Generator.StackCounter = v.SizeFunction
	} else {
		v.Generator.StackCounter = v.SizeFunction - len(params["internal"])
	}

	// set false
	v.Generator.FunctionCode = false
	// gen c3d
	v.Generator.FuncionUser = append(v.Generator.FuncionUser, "return;\n")
	// close function
	v.Generator.FuncionUser = append(v.Generator.FuncionUser, "}\n")

	// get the return type if it exists
	var returnType string
	if ctx.ARROW_FUNCTION() != nil {

		returnType = ctx.Type_().GetText()
		// save the function in the symbol table
		symbol := Symbol{
			Id:           idFunction,
			TypeSymbol:   values.Function_Type,
			TypeMutable:  returnType,
			TypeData:     values.Function_Type,
			ReturnType:   returnType,
			ReturnTemp:   v.ReturnTemp,
			ReturnLabel:  v.ReturnLabel,
			sizeFunction: v.SizeFunction,
			ListParams:   params,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}

		//
		v.getCurrentScope()[idFunction] = symbol

		v.TableSymbol = append(v.TableSymbol, symbol)

	} else {
		returnType = "void"
		// save the function in the symbol table
		symbol := Symbol{
			Id:           idFunction,
			TypeSymbol:   values.Function_Type,
			TypeMutable:  returnType,
			TypeData:     values.Function_Type,
			ReturnType:   returnType,
			sizeFunction: v.SizeFunction,
			ReturnTemp:   v.ReturnTemp,
			ReturnLabel:  v.ReturnLabel,
			ListParams:   params,
			Line:         ctx.GetStart().GetLine(),
			Column:       ctx.GetStart().GetColumn(),
		}

		v.getCurrentScope()[idFunction] = symbol

		v.TableSymbol = append(v.TableSymbol, symbol)
	}

	// fmt.Println("FunctionWithoutParams: ", idFunction, returnType)

	return nil
}

// VisitListFunctionParamsEI
func (v *Visitor) VisitListFunctionParamsEI(ctx *parser.ListFunctionParamsEIContext) interface{} {

	// create a map to save the params
	params := make(map[string][]Symbol)

	// create two keys, external and internal
	params["external"] = []Symbol{}
	params["internal"] = []Symbol{}

	listIds := ctx.AllID_PRIMITIVE()

	listTypes := ctx.AllType_()

	var mutable string

	if ctx.INOUT() != nil {
		mutable = values.VarMutable
	} else {
		mutable = values.LetMutable
	}

	// iterate over the list of ids
	for i, id := range listIds {
		// get the type

		typeParam := listTypes[i/len(listTypes)].GetText()

		// create a new symbol table
		symbol := Symbol{
			Id:             id.GetText(),
			TypeSymbol:     values.Variable_Type,
			TypeMutable:    mutable,
			TypeData:       typeParam,
			StackDirection: v.SizeFunction,
			Value:          values.NewC3DPrimitive("9999999827968.00", "nil", typeParam, false),
			Line:           ctx.GetStart().GetLine(),
			Column:         ctx.GetStart().GetColumn(),
		}
		// comes external then internal
		if i%2 == 0 {
			params["external"] = append(params["external"], symbol)
		} else {
			params["internal"] = append(params["internal"], symbol)
			// add the symbol to the current scope
			v.getCurrentScope()[id.GetText()] = symbol
			// increase the size of the function
			v.SizeFunction += 1
		}

	}

	// fmt.Println("ListFunctionParamsEI:\n ", params)

	return params
}

// VisitListFunctionParamsNEI
func (v *Visitor) VisitListFunctionParamsNEI(ctx *parser.ListFunctionParamsNEIContext) interface{} {

	// create a map to save the params
	params := make(map[string][]Symbol)

	// create two keys, external and internal
	params["internal"] = []Symbol{}

	listIds := ctx.AllID_PRIMITIVE()
	listTypes := ctx.AllType_()
	var mutable string

	if ctx.INOUT() != nil {
		mutable = values.VarMutable
	} else {
		mutable = values.LetMutable
	}

	// iterate over the list of ids
	for i, id := range listIds {
		// get the type
		typeParam := listTypes[i].GetText()

		// create a new Symbol table
		symbol := Symbol{
			Id:             id.GetText(),
			TypeSymbol:     values.Variable_Type,
			TypeMutable:    mutable,
			TypeData:       typeParam,
			StackDirection: v.SizeFunction,
			Value:          values.NewC3DPrimitive("9999999827968.00", "nil", typeParam, false),
			Line:           ctx.GetStart().GetLine(),
			Column:         ctx.GetStart().GetColumn(),
		}
		// comes external then internal
		params["internal"] = append(params["internal"], symbol)
		v.getCurrentScope()[id.GetText()] = symbol
		// increase the size of the function
		v.SizeFunction += 1
	}

	// fmt.Println("ListFunctionParamsNEI:\n", params)

	return params
}

// VisitListFunctionParamsBEI
func (v *Visitor) VisitListFunctionParamsBEI(ctx *parser.ListFunctionParamsBEIContext) interface{} {

	// create a map to save the params
	params := make(map[string][]Symbol)

	// create two keys, external and internal
	params["internal"] = []Symbol{}

	listIds := ctx.AllID_PRIMITIVE()
	listTypes := ctx.AllType_()

	// iterate over the list of ids
	for i, id := range listIds {
		// get the type
		typeParam := listTypes[i].GetText()

		// create a new symbol table
		symbol := Symbol{
			Id:             id.GetText(),
			TypeSymbol:     values.Variable_Type,
			TypeMutable:    values.LetMutable,
			TypeData:       typeParam,
			StackDirection: v.SizeFunction,
			Value:          values.NewC3DPrimitive("9999999827968.00", "nil", typeParam, false),
			Line:           ctx.GetStart().GetLine(),
			Column:         ctx.GetStart().GetColumn(),
		}
		// comes external then internal
		params["internal"] = append(params["internal"], symbol)
		v.getCurrentScope()[id.GetText()] = symbol
		// increase the size of the function
		v.SizeFunction += 1

	}

	// fmt.Println("ListFunctionParamsBEI:\n ", params)

	return params
}
