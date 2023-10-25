package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

// Visit type declaration with value or not
func (v *Visitor) VisitTypeValueDeclaration(ctx *parser.TypeValueDeclarationContext) interface{} {
	// get the id of the variable
	varId := ctx.ID_PRIMITIVE().GetText()
	// verify if the variable is in the scope
	if v.VerifyVariableCurrentScope(varId) {
		log.Printf("Error: Variable '%s' already declared \n", varId)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' already declared", varId),
			Type:   "Semantic",
		})
		return nil
	}
	// get the type of the variable -> var or let
	varType := ctx.Type_declaration().GetText()
	// get the type of the value -> Int, Float, String, Boolean, Character
	varTypeValue := ctx.Type_().GetText()
	// get the value of the variable
	varValue := v.Visit(ctx.Expr()).(*values.C3DPrimitive)
	// cast int to float
	if varTypeValue == values.FloatType && varValue.GetType() == values.IntType {
		// cast the value to float
		varValue.TypeData = values.FloatType

	} else if varTypeValue != varValue.GetType() { // evaluate if the type of the value is the same as the type of the variable
		// add error
		log.Printf("Error: Type of the value '%s' is not the same as the type of the variable '%s' \n", varTypeValue, varValue.GetType())
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Type of the value '%s' is not the same as the type of the variable '%s'", varTypeValue, varType),
			Type:   "Semantic",
		})
		return nil
	}

	// gen c3d
	stack, heap := v.Generator.GenDeclaration(varValue.GetType(), varId, varValue, v.ExistsLoopContext())

	symbol := Symbol{
		Id:             varId,
		TypeSymbol:     values.Variable_Type,
		TypeMutable:    varType,
		TypeData:       varTypeValue,
		Value:          varValue,
		Line:           ctx.GetStart().GetLine(),
		Column:         ctx.GetStart().GetColumn(),
		StackDirection: stack,
		HeapDirection:  heap,
	}
	// increment the stack
	v.Generator.CounterStack("+")

	// add the variable to the scope
	v.getCurrentScope()[varId] = symbol

	// apppend to the TableSymbol
	// v.TableSymbol = append(v.TableSymbol, symbol)
	fmt.Println("Global scope or symbol table ->", v.SymbolStack)

	return nil
}

// Visit type declaration with value or not
func (v *Visitor) VisitTypeOptionalValueDeclaration(ctx *parser.TypeOptionalValueDeclarationContext) interface{} {
	// get the id of the variable
	varId := ctx.ID_PRIMITIVE().GetText()
	// verify if the variable is in the scope
	if v.VerifyVariableCurrentScope(varId) {
		log.Printf("Error: Variable '%s' already declared \n", varId)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' already declared", varId),
			Type:   "Semantic",
		})
		return nil
	}
	// get the type of the variable -> var or let
	varType := ctx.Type_declaration().GetText()
	// get the type of the value -> Int, Float, String, Boolean, Character
	varTypeValue := ctx.Type_().GetText()
	// get the question mark
	questionMark := ctx.QUESTION_MARK()

	// evaluate if the var or let
	if varType == "var" {
		// save the value of the variable as nil
		varValue := values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
		// generate temp
		stack, heap := v.Generator.GenDeclaration(varValue.GetType(), varId, varValue, v.ExistsLoopContext())

		symbol := Symbol{
			Id:             varId,
			TypeSymbol:     values.Variable_Type,
			TypeMutable:    varType,
			TypeData:       varTypeValue,
			Value:          varValue,
			Line:           ctx.GetStart().GetLine(),
			Column:         ctx.GetStart().GetColumn(),
			StackDirection: stack,
			HeapDirection:  heap,
		}

		// increment the stack
		v.Generator.CounterStack("+")

		// add the variable to the scope
		v.getCurrentScope()[varId] = symbol

		// apppend to the TableSymbol
		// v.TableSymbol = append(v.TableSymbol, symbol)
		fmt.Println("Global scope or symbol table ->", v.SymbolStack)

	} else if varType == "let" {
		if questionMark != nil {
			log.Printf("Error: Variable '%s' is a constant and cannot be optional \n", varId)
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Variable '%s' is a constant and cannot be optional", varId),
				Type:   "Semantic",
			})
			return nil
		}
	}

	return nil
}

// Visit type declaration without type value (infer)

func (v *Visitor) VisitValueDeclaration(ctx *parser.ValueDeclarationContext) interface{} {
	// get the id of the variable
	varId := ctx.ID_PRIMITIVE().GetText()
	// verify if the variable is in the scope
	if v.VerifyVariableCurrentScope(varId) {
		log.Printf("Error: Variable '%s' already declared \n", varId)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' already declared", varId),
			Type:   "Semantic",
		})
		return nil
	}
	// get the type of the variable -> var or let
	varType := ctx.Type_declaration().GetText()
	// get the type of the value -> Int, Float, String, Boolean, Character
	varValue := v.Visit(ctx.Expr()).(*values.C3DPrimitive)
	// get the type of the value
	varTypeValue := varValue.GetType()

	fmt.Println("value -->>", v.Generator.FunctionCode)
	// if its int, float or boolean
	stack, heap := v.Generator.GenDeclaration(varValue.Value, varId, varValue, v.ExistsLoopContext())

	symbol := Symbol{
		Id:             varId,
		TypeSymbol:     values.Variable_Type,
		TypeMutable:    varType,
		TypeData:       varTypeValue,
		Value:          varValue,
		Line:           ctx.GetStart().GetLine(),
		Column:         ctx.GetStart().GetColumn(),
		StackDirection: stack,
		HeapDirection:  heap,
	}

	// increment the stack
	v.Generator.CounterStack("+")

	// add the variable to the scope
	v.getCurrentScope()[varId] = symbol

	// apppend to the TableSymbol
	// v.TableSymbol = append(v.TableSymbol, symbol)

	fmt.Println("Global scope or symbol table ->", v.SymbolStack)

	return nil

}
