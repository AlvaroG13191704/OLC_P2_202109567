package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

// CallFunctionExpr
func (v *Visitor) VisitCallFunctionExpr(ctx *parser.CallFunctionExprContext) interface{} {
	// valueExpr := v.Visit(ctx.CallFunctionStmt()).(*values.C3DPrimitive)
	fmt.Println("CallFunctionExpr")

	return v.Visit(ctx.CallFunctionStmt())
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
	if v.Generator.FunctionCode {
		v.Generator.FuncionUser = append(v.Generator.FuncionUser, fmt.Sprintf("%s();\n", function.Id))
	} else {
		v.Generator.Code = append(v.Generator.Code, fmt.Sprintf("%s();\n", function.Id))
	}

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

// VisitCallFunctionWithParamsEI
func (v *Visitor) VisitCallFunctionWithParams(ctx *parser.CallFunctionWithParamsContext) interface{} {
	fmt.Println("CallFunctionWithParams")
	// push a new function context
	v.PushFunctionContext("function")
	defer v.PopFunctionContext()

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
	// get list params
	listParams := function.ListParams
	// get list of arguments
	listArguments := v.Visit(ctx.ListCallFunctionStmt())

	// create temp
	tempStackPointer := v.Generator.NewTemp()
	// assign the P to the temp
	v.Generator.AssignTemp(tempStackPointer, "P") // create temp

	// INTERNAL - EXTERNAL
	if listParams["external"] != nil && listArguments.(map[string][]Symbol)["external"] != nil && listParams["internal"] != nil {
		fmt.Println("params with external and internal --> Value value")
		fmt.Println("listParams ->", listParams)
		fmt.Println("listArguments ->", listArguments)
		// validate external params with external arguments
		if len(listParams["external"]) != len(listArguments.(map[string][]Symbol)["external"]) {
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("function %s expected %d arguments, got %d", functionName, len(listParams["external"]), len(listArguments.(map[string][]Symbol)["external"])),
				Type:   "Semantic",
			})
			log.Printf("function %s expected %d arguments, got %d", functionName, len(listParams["external"]), len(listArguments.(map[string][]Symbol)["external"]))
			return nil
		}

		// evaluate if the values are the same type in external params and external arguments
		for i, param := range listParams["external"] {
			// get the value of the expression
			value := listArguments.(map[string][]Symbol)["external"][i].Value.(*values.C3DPrimitive)

			// evaluate if the values are the same type
			if param.TypeData != value.GetType() {
				// add error
				v.Errors = append(v.Errors, Error{
					Line:   ctx.GetStart().GetLine(),
					Column: ctx.GetStart().GetColumn(),
					Msg:    fmt.Sprintf("function %s expected %s type, got %s", functionName, param.TypeData, value.GetType()),
					Type:   "Semantic",
				})
				log.Printf("function %s expected %s type, got %s", functionName, param.TypeData, value.GetType())
				return nil
			}

		}

		// asign the values to the internal params
		for i, param := range listParams["internal"] {
			// get the value of the expression
			value := listArguments.(map[string][]Symbol)["external"][i].Value.(*values.C3DPrimitive)

			// create the the temp
			temp := v.Generator.NewTemp()
			// get the position of the param
			if listArguments.(map[string][]Symbol)["external"][i].Reference {
				v.Generator.GenArithmetic(temp, "P", fmt.Sprintf("%d", param.StackDirection+len(listArguments.(map[string][]Symbol)["external"])), "+")
			} else {
				v.Generator.GenArithmetic(temp, "P", fmt.Sprintf("%d", param.StackDirection), "+")
			}
			// save in the stack
			v.Generator.SaveStack(temp, value.GetValue())

		}

	} else if listParams["internal"] != nil && listArguments.(map[string][]Symbol)["internal"] != nil {
		fmt.Println("params with only internal, not external --> _ value")
		fmt.Println("listParams ->", listParams)
		fmt.Println("listArguments ->", listArguments)
		// validate internal params with internal arguments
		if len(listParams["internal"]) != len(listArguments.(map[string][]Symbol)["internal"]) {
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("function %s expected %d arguments, got %d", functionName, len(listParams["internal"]), len(listArguments.(map[string][]Symbol)["internal"])),
				Type:   "Semantic",
			})
			log.Printf("function %s expected %d arguments, got %d", functionName, len(listParams["internal"]), len(listArguments.(map[string][]Symbol)["internal"]))
			return nil
		}

		// evaluate if the values are the same type in internal params and internal arguments
		for i, param := range listParams["internal"] {
			// get the value of the expression
			value := listArguments.(map[string][]Symbol)["internal"][i].Value.(*values.C3DPrimitive)

			// evaluate if the values are the same type
			if param.TypeData != value.GetType() {
				// add error
				v.Errors = append(v.Errors, Error{
					Line:   ctx.GetStart().GetLine(),
					Column: ctx.GetStart().GetColumn(),
					Msg:    fmt.Sprintf("function %s expected %s type, got %s", functionName, param.TypeData, value.GetType()),
					Type:   "Semantic",
				})
				log.Printf("function %s expected %s type, got %s", functionName, param.TypeData, value.GetType())
				return nil

			}
			// create the the temp
			temp := v.Generator.NewTemp()
			// get the position of the param
			if listArguments.(map[string][]Symbol)["internal"][i].Reference {
				v.Generator.GenArithmetic(temp, "P", fmt.Sprintf("%d", param.StackDirection+len(listArguments.(map[string][]Symbol)["internal"])), "+")
			} else {
				v.Generator.GenArithmetic(temp, "P", fmt.Sprintf("%d", param.StackDirection), "+")
			}
			// save in the stack
			v.Generator.SaveStack(temp, value.GetValue())
		}

	} else if listParams["internal"] != nil && listArguments.(map[string][]Symbol)["external"] != nil {
		fmt.Println("params with both internal and external, same name")
		fmt.Println("listParams ->", listParams)
		fmt.Println("listArguments ->", listArguments)
		// validate external params with external arguments
		if len(listParams["internal"]) != len(listArguments.(map[string][]Symbol)["external"]) {
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("function %s expected %d arguments, got %d", functionName, len(listParams["internal"]), len(listArguments.(map[string][]Symbol)["external"])),
				Type:   "Semantic",
			})
			log.Printf("function %s expected %d arguments, got %d", functionName, len(listParams["internal"]), len(listArguments.(map[string][]Symbol)["external"]))
			return nil
		}

		// evaluate if the values are the same type in external params and external arguments
		for i, param := range listParams["internal"] {
			// get the value of the expression
			value := listArguments.(map[string][]Symbol)["external"][i].Value.(*values.C3DPrimitive)

			// evaluate if the values are the same type
			if param.TypeData != value.GetType() {
				// add error
				v.Errors = append(v.Errors, Error{
					Line:   ctx.GetStart().GetLine(),
					Column: ctx.GetStart().GetColumn(),
					Msg:    fmt.Sprintf("function %s expected %s type, got %s", functionName, param.TypeData, value.GetType()),
					Type:   "Semantic",
				})
				log.Printf("function %s expected %s type, got %s", functionName, param.TypeData, value.GetType())
				return nil
			}

			// create the the temp
			temp := v.Generator.NewTemp()
			// get the position of the param
			if listArguments.(map[string][]Symbol)["external"][i].Reference {
				v.Generator.GenArithmetic(temp, "P", fmt.Sprintf("%d", param.StackDirection+len(listArguments.(map[string][]Symbol)["external"])), "+")
			} else {
				v.Generator.GenArithmetic(temp, "P", fmt.Sprintf("%d", param.StackDirection), "+")
			}
			// save in the stack
			v.Generator.SaveStack(temp, value.GetValue())

		}

	}
	// reasign the stack pointer
	v.Generator.AssignTemp("P", tempStackPointer)

	// call the function

	if v.Generator.FunctionCode {
		v.Generator.FuncionUser = append(v.Generator.FuncionUser, fmt.Sprintf("%s();\n", function.Id))
	} else {
		v.Generator.Code = append(v.Generator.Code, fmt.Sprintf("%s();\n", function.Id))
	}

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

// listCallFunctionStmtEI
func (v *Visitor) VisitListCallFunctionStmtEI(ctx *parser.ListCallFunctionStmtEIContext) interface{} {
	// create a map to save the params
	params := make(map[string][]Symbol)

	// create two keys, external and internal
	params["external"] = []Symbol{}

	listIds := ctx.AllID_PRIMITIVE()
	listExpr := ctx.AllExpr()

	var reference bool

	for i, id := range listIds {
		if ctx.REFERENCE(i) != nil {
			reference = true
		} else {
			reference = false
		}
		value := v.Visit(listExpr[i])

		// get the id
		idValue := id.GetText()

		primitiveValue := value.(*values.C3DPrimitive)
		// create a symbol table
		symbolTable := Symbol{
			Id:          idValue,
			TypeSymbol:  values.Variable_Type,
			TypeMutable: values.LetMutable,
			TypeData:    primitiveValue.GetType(),
			Value:       primitiveValue,
			Reference:   reference,
			ListParams:  nil,
			Line:        ctx.GetStart().GetLine(),
			Column:      ctx.GetStart().GetColumn(),
		}

		fmt.Println("Reference ->", reference)

		// append the symbol table to the params
		params["external"] = append(params["external"], symbolTable)

	}

	return params
}

// VisitListCallFunctionStmtNEI
func (v *Visitor) VisitListCallFunctionStmtNEI(ctx *parser.ListCallFunctionStmtNEIContext) interface{} {
	// create a map to save the params
	params := make(map[string][]Symbol)

	// create two keys, external and internal
	params["internal"] = []Symbol{}

	listExpr := ctx.AllExpr()
	var reference bool
	// iterate over the list of ids and save the values
	for i, expr := range listExpr {
		if ctx.REFERENCE(i) != nil {
			reference = true
		} else {
			reference = false
		}
		// get the value of the expression
		value := v.Visit(expr)

		// create a symbol table
		primitiveValue := value.(*values.C3DPrimitive)
		// create a symbol table
		symbolTable := Symbol{
			Id:          expr.GetText(),
			TypeSymbol:  values.Variable_Type,
			TypeMutable: values.LetMutable,
			TypeData:    primitiveValue.GetType(),
			Value:       primitiveValue,
			Reference:   reference,
			ListParams:  nil,
			Line:        ctx.GetStart().GetLine(),
			Column:      ctx.GetStart().GetColumn(),
		}
		// append the symbol table to the params
		params["internal"] = append(params["internal"], symbolTable)
		fmt.Println("Reference ->", reference)
	}

	return params
}
