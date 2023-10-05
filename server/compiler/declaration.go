package compiler

// import (
// 	"fmt"
// 	"log"
// 	"server/compiler/compiler/values"
// 	"server/compiler/parser"
// )

// // Visit type declaration with value or not
// func (v *Visitor) VisitTypeValueDeclaration(ctx *parser.TypeValueDeclarationContext) interface{} {
// 	// get the id of the variable
// 	varId := ctx.ID_PRIMITIVE().GetText()
// 	// verify if the variable is in the scope
// 	if v.VerifyVariableCurrentScope(varId) {
// 		log.Printf("Error: Variable '%s' already declared \n", varId)
// 		// add error
// 		v.Errors = append(v.Errors, Error{
// 			Line:   ctx.GetStart().GetLine(),
// 			Column: ctx.GetStart().GetColumn(),
// 			Msg:    fmt.Sprintf("Error: Variable '%s' already declared", varId),
// 			Type:   "Semantic",
// 		})
// 		return nil
// 	}
// 	// get the type of the variable -> var or let
// 	varType := ctx.Type_declaration().GetText()
// 	// get the type of the value -> Int, Float, String, Boolean, Character
// 	varTypeValue := ctx.Type_().GetText()
// 	// get the value of the variable
// 	// varValue := v.Visit(ctx.Expr()).(values.PRIMITIVE)
// 	varValue := v.Visit(ctx.Expr()).(values.PRIMITIVE)
// 	// cast int to float
// 	if varTypeValue == "Float" && varValue.GetType() == "Int" {
// 		// cast the value to float
// 		// varValue = &values.Float{
// 		// 	Value: float64(varValue.GetValue().(int64)),
// 		// }
// 	} else if varTypeValue != varValue.GetType() { // evaluate if the type of the value is the same as the type of the variable
// 		// add error
// 		log.Printf("Error: Type of the value '%s' is not the same as the type of the variable '%s' \n", varTypeValue, varValue.GetType())
// 		v.Errors = append(v.Errors, Error{
// 			Line:   ctx.GetStart().GetLine(),
// 			Column: ctx.GetStart().GetColumn(),
// 			Msg:    fmt.Sprintf("Error: Type of the value '%s' is not the same as the type of the variable '%s'", varTypeValue, varType),
// 			Type:   "Semantic",
// 		})
// 		return nil
// 	}

// 	// if its int, float or boolean
// 	tempString, stack, heap := v.Generator.GenDeclaration(varValue.GetType(), varId, varValue)

// 	symbol := Symbol{
// 		Id:             varId,
// 		TypeSymbol:     values.Variable_Type,
// 		TypeMutable:    varType,
// 		TypeData:       varTypeValue,
// 		Value:          varValue,
// 		Line:           ctx.GetStart().GetLine(),
// 		Column:         ctx.GetStart().GetColumn(),
// 		TempString:     tempString,
// 		StackDirection: stack,
// 		HeapDirection:  heap,
// 	}

// 	// add the variable to the scope
// 	v.getCurrentScope()[varId] = symbol

// 	// apppend to the TableSymbol
// 	// v.TableSymbol = append(v.TableSymbol, symbol)
// 	// print the symbol table
// 	// fmt.Println("Current scope or symbol table ->", v.getCurrentScope())
// 	fmt.Println("Global scope or symbol table ->", v.SymbolStack)

// 	return nil
// }

// // // Visit type declaration with value or not
// // func (v *Visitor) VisitTypeOptionalValueDeclaration(ctx *parser.TypeOptionalValueDeclarationContext) interface{} {
// // 	// get the id of the variable
// // 	varId := ctx.ID_PRIMITIVE().GetText()
// // 	// verify if the variable is in the scope
// // 	if v.VerifyVariableCurrentScope(varId) {
// // 		log.Printf("Error: Variable '%s' already declared \n", varId)
// // 		// add error
// // 		v.Errors = append(v.Errors, Error{
// // 			Line:   ctx.GetStart().GetLine(),
// // 			Column: ctx.GetStart().GetColumn(),
// // 			Msg:    fmt.Sprintf("Error: Variable '%s' already declared", varId),
// // 			Type:   "Semantic",
// // 		})
// // 		return nil
// // 	}
// // 	// get the type of the variable -> var or let
// // 	varType := ctx.Type_declaration().GetText()
// // 	// get the type of the value -> Int, Float, String, Boolean, Character
// // 	varTypeValue := ctx.Type_().GetText()
// // 	// get the question mark
// // 	questionMark := ctx.QUESTION_MARK()

// // 	// evaluate if the var or let
// // 	if varType == "var" {
// // 		// save the value of the variable as nil
// // 		varValue := &values.Nil{
// // 			Value: nil,
// // 		}
// // 		// generate temp
// // 		// newTemp := v.Generator.NewTemp()
// // 		// if its int, float or boolean
// // 		tempString, stack, heap := v.Generator.GenDeclaration(var.val, varId, varValue)

// // 		symbol := Symbol{
// // 			Id:             varId,
// // 			TypeSymbol:     values.Variable_Type,
// // 			TypeMutable:    varType,
// // 			TypeData:       varTypeValue,
// // 			Value:          varValue,
// // 			Line:           ctx.GetStart().GetLine(),
// // 			Column:         ctx.GetStart().GetColumn(),
// // 			TempString:     tempString,
// // 			StackDirection: stack,
// // 			HeapDirection:  heap,
// // 		}

// // 		// add the variable to the scope
// // 		v.getCurrentScope()[varId] = symbol

// // 		// apppend to the TableSymbol
// // 		// v.TableSymbol = append(v.TableSymbol, symbol)
// // 		// print the symbol table
// // 		fmt.Println("Global scope or symbol table ->", v.SymbolStack)

// // 	} else if varType == "let" {
// // 		if questionMark != nil {
// // 			log.Printf("Error: Variable '%s' is a constant and cannot be optional \n", varId)
// // 			// add error
// // 			v.Errors = append(v.Errors, Error{
// // 				Line:   ctx.GetStart().GetLine(),
// // 				Column: ctx.GetStart().GetColumn(),
// // 				Msg:    fmt.Sprintf("Error: Variable '%s' is a constant and cannot be optional", varId),
// // 				Type:   "Semantic",
// // 			})
// // 			return nil
// // 		}
// // 	}

// // 	return nil
// // }

// // // Visit type declaration without type value (infer)

// // func (v *Visitor) VisitValueDeclaration(ctx *parser.ValueDeclarationContext) interface{} {
// // 	// get the id of the variable
// // 	varId := ctx.ID_PRIMITIVE().GetText()
// // 	// verify if the variable is in the scope
// // 	if v.VerifyVariableCurrentScope(varId) {
// // 		log.Printf("Error: Variable '%s' already declared \n", varId)
// // 		// add error
// // 		v.Errors = append(v.Errors, Error{
// // 			Line:   ctx.GetStart().GetLine(),
// // 			Column: ctx.GetStart().GetColumn(),
// // 			Msg:    fmt.Sprintf("Error: Variable '%s' already declared", varId),
// // 			Type:   "Semantic",
// // 		})
// // 		return nil
// // 	}
// // 	// get the type of the variable -> var or let
// // 	varType := ctx.Type_declaration().GetText()
// // 	// get the type of the value -> Int, Float, String, Boolean, Character
// // 	// get the value of the variable
// // 	varValue := v.Visit(ctx.Expr()).(values.PRIMITIVE)
// // 	// get the type of the value
// // 	varTypeValue := varValue.GetType()

// // 	// generate temp
// // 	newTemp := v.Generator.NewTemp()
// // 	// if its int, float or boolean
// // 	tempString, stack, heap := v.Generator.GenDeclaration(newTemp, varId, varValue)

// // 	symbol := Symbol{
// // 		Id:             varId,
// // 		TypeSymbol:     values.Variable_Type,
// // 		TypeMutable:    varType,
// // 		TypeData:       varTypeValue,
// // 		Value:          varValue,
// // 		Line:           ctx.GetStart().GetLine(),
// // 		Column:         ctx.GetStart().GetColumn(),
// // 		TempString:     tempString,
// // 		StackDirection: stack,
// // 		HeapDirection:  heap,
// // 	}

// // 	// add the variable to the scope
// // 	v.getCurrentScope()[varId] = symbol

// // 	// apppend to the TableSymbol
// // 	// v.TableSymbol = append(v.TableSymbol, symbol)
// // 	// print the symbol table
// // 	fmt.Println("Global scope or symbol table ->", v.SymbolStack)

// // 	// apppend to the TableSymbol
// // 	// v.TableSymbol = append(v.TableSymbol, symbol)

// // 	// print the symbol table
// // 	// fmt.Println("Current scope or symbol table ->", v.getCurrentScope())
// // 	fmt.Println("Global scope or symbol table ->", v.SymbolStack)

// // 	return nil

// // }
