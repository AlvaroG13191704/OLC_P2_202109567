package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

// VisitVectorDeclaration
func (v *Visitor) VisitVectorDeclaration(ctx *parser.VectorDeclarationContext) interface{} {
	// get the id of the variable
	varId := ctx.AllID_PRIMITIVE()[0].GetText()
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
	// get the type of the value -> Int, Float, String, Boolean, Character or Vector
	varTypeValue := ctx.Type_().GetText()

	// if it's a copy of the vector
	if len(ctx.AllID_PRIMITIVE()) == 2 {
		// get the symbol table of the variable
		symbol, ok := v.VerifyScope(ctx.AllID_PRIMITIVE()[1].GetText())
		// evaluate if the variable is in the scope
		if !ok {
			log.Printf("Error: Variable '%s' not declared \n", ctx.AllID_PRIMITIVE()[1].GetText())
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Variable '%s' not declared", ctx.AllID_PRIMITIVE()[1].GetText()),
				Type:   "Semantic",
			})
			return nil
		}
		// evaluate if the type of the variable is not a vector
		symbolTable := symbol.(Symbol)
		if symbolTable.TypeSymbol != values.Vector_Type {
			log.Printf("Error: Variable '%s' is not a vector \n", ctx.AllID_PRIMITIVE()[1].GetText())
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Variable '%s' is not a vector", ctx.AllID_PRIMITIVE()[1].GetText()),
				Type:   "Semantic",
			})
			return nil
		}

		// gen c3d
		v.Generator.GenComment("Vector copy Declaration")

		// temps
		tempPointer := v.Generator.NewTemp()

		// get the value from the stack
		v.Generator.GetStack(tempPointer, fmt.Sprintf("%d", symbolTable.StackDirection))

		// declare
		stack, heap := v.Generator.GenDeclarationVector(tempPointer, varId, false, v.SizeFunction)
		// if it's a function declaration increase the size of the pointer
		if v.Generator.FunctionCode {
			v.SizeFunction++
		}
		newSymbol := Symbol{
			Id:             varId,
			TypeSymbol:     values.Vector_Type,
			TypeMutable:    varType,
			TypeData:       varTypeValue,
			Value:          tempPointer,
			LenghtVector:   symbolTable.LenghtVector,
			Line:           ctx.GetStart().GetLine(),
			Column:         ctx.GetStart().GetColumn(),
			StackDirection: stack,
			HeapDirection:  heap,
		}

		// increment the stack
		v.Generator.CounterStack("+")

		// add the variable to the scope
		v.getCurrentScope()[varId] = newSymbol

		fmt.Println("Global scope or symbol table ->", v.SymbolStack)
		return nil
	}
	// get the list of expresions
	var listTempsString []string
	var lenght int
	// temps
	tempPointer := v.Generator.NewTemp()
	if ctx.ExprList() != nil {

		exprList := ctx.ExprList().(*parser.ExprListContext)
		expressions := exprList.AllExpr()

		// iterate the expressions
		for _, expr := range expressions {
			// get the value of the expression
			value := v.Visit(expr).(*values.C3DPrimitive)
			// validate
			if value.GetType() != varTypeValue {
				log.Printf("Error: Type of the value '%s' is not the same as the type of the variable '%s' \n", varTypeValue, value.GetType())
				// add error
				v.Errors = append(v.Errors, Error{
					Line:   ctx.GetStart().GetLine(),
					Column: ctx.GetStart().GetColumn(),
					Msg:    fmt.Sprintf("Error: Type of the value '%s' is not the same as the type of the variable '%s'", varTypeValue, value.GetType()),
					Type:   "Semantic",
				})
				return nil
			}
			if value.GetType() == values.StringType {
				tempString := v.Generator.GenString(value.GetValue())
				// add to the list of temps
				listTempsString = append(listTempsString, tempString)
			}

		}
		// gen c3d
		v.Generator.GenComment("Vector Declaration")
		v.Generator.AssignTemp(tempPointer, "H")
		// save the length of the vector in the first position
		v.Generator.SaveHeap(fmt.Sprintf("%d", len(expressions)))
		lenght = len(expressions)
		// increase the heap
		v.Generator.IncPointerHeap()
		v.Generator.GenComment("Start saving vector values")
		// iterate the expressions
		for i, expr := range expressions {
			// get the value of the expression
			value := v.Visit(expr).(*values.C3DPrimitive)
			// append value
			// listExpr = append(listExpr, value)
			// validate
			if value.GetType() != varTypeValue {
				log.Printf("Error: Type of the value '%s' is not the same as the type of the variable '%s' \n", varTypeValue, value.GetType())
				// add error
				v.Errors = append(v.Errors, Error{
					Line:   ctx.GetStart().GetLine(),
					Column: ctx.GetStart().GetColumn(),
					Msg:    fmt.Sprintf("Error: Type of the value '%s' is not the same as the type of the variable '%s'", varTypeValue, value.GetType()),
					Type:   "Semantic",
				})
				return nil
			}
			if value.GetType() == values.StringType {
				// save the value in the heap
				v.Generator.SaveHeap(listTempsString[i])
				// increase the heap
				v.Generator.IncPointerHeap()
			} else {
				// save the value in the heap
				v.Generator.SaveHeap(value.GetValue())
				// increase the heap
				v.Generator.IncPointerHeap()
			}

		}

	} else {
		// add nothing
		// gen c3d
		v.Generator.GenComment("Vector Declaration")
		v.Generator.AssignTemp(tempPointer, "H")
		// save the length of the vector in the first position
		v.Generator.SaveHeap("0")
		// increase the heap
		v.Generator.IncPointerHeap()
	}

	// declare
	stack, heap := v.Generator.GenDeclarationVector(tempPointer, varId, false, v.SizeFunction)
	// if it's a function declaration increase the size of the pointer
	if v.Generator.FunctionCode {
		v.SizeFunction++
	}

	symbol := Symbol{
		Id:             varId,
		TypeSymbol:     values.Vector_Type,
		TypeMutable:    varType,
		TypeData:       varTypeValue,
		Value:          tempPointer,
		LenghtVector:   lenght,
		Line:           ctx.GetStart().GetLine(),
		Column:         ctx.GetStart().GetColumn(),
		StackDirection: stack,
		HeapDirection:  heap,
	}

	// increment the stack
	v.Generator.CounterStack("+")

	// add the variable to the scope
	v.getCurrentScope()[varId] = symbol

	// // apppend to the TableSymbol
	v.TableSymbol = append(v.TableSymbol, symbol)

	fmt.Println("Global scope or symbol table ->", v.SymbolStack)
	return nil

}
