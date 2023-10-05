package compiler

// import (
// 	"fmt"
// 	"log"
// 	"server/compiler/compiler/values"
// 	"server/compiler/parser"
// )

// func (v *Visitor) VisitValueAssignment(ctx *parser.ValueAssignmentContext) interface{} {

// 	// get the variable name
// 	idName := ctx.ID_PRIMITIVE().GetText()
// 	// get the expr
// 	expr := v.Visit(ctx.Expr()).(values.PRIMITIVE)
// 	// get the value
// 	valueFromScope, ok := v.VerifyScope(idName)
// 	// evaluate if the name is declared
// 	if ok {
// 		// get the value
// 		value := valueFromScope.(Symbol)
// 		// evaluate if its a constant
// 		if value.TypeMutable == "let" {
// 			log.Fatalf("Error: Variable '%s' is a constant, cannot be assign", idName)
// 			v.Errors = append(v.Errors, Error{
// 				Line:   ctx.GetStart().GetLine(),
// 				Column: ctx.GetStart().GetColumn(),
// 				Msg:    "Variable '" + idName + "' is a constant",
// 				Type:   "VariableError",
// 			})
// 			return nil
// 		}
// 		// get the type of the variable

// 		typeVariable := value.Value.(values.PRIMITIVE).GetType()
// 		// evaluate if the type of the variable is the same of the expr
// 		if typeVariable == expr.GetType() || typeVariable == values.NilType {
// 			// update the value
// 			// gen c3d
// 			tempString, stack, heap := v.Generator.GenAssignment(idName, value.StackDirection, expr)
// 			// update the value
// 			value.Value = expr
// 			// update the stack and heap
// 			value.StackDirection = stack
// 			value.HeapDirection = heap
// 			value.TempString = tempString
// 			// update the value no matter the scope
// 			v.UpdateVariable(idName, value)

// 			fmt.Println("Global scope or symbol table ->", v.SymbolStack)
// 			return nil
// 		} else {
// 			log.Fatalf("Error: Variable '%s' is of type %s, cannot be assign to %s", idName, typeVariable, expr.GetType())
// 			v.Errors = append(v.Errors, Error{
// 				Line:   ctx.GetStart().GetLine(),
// 				Column: ctx.GetStart().GetColumn(),
// 				Msg:    "Variable '" + idName + "' is of type " + typeVariable + ", cannot be assign to " + expr.GetType(),
// 				Type:   "TypeError",
// 			})
// 			return nil
// 		}

// 	} else {
// 		// add the error to the errors
// 		log.Fatalf("Error: Variable '%s' not declared", idName)
// 		v.Errors = append(v.Errors, Error{
// 			Line:   ctx.GetStart().GetLine(),
// 			Column: ctx.GetStart().GetColumn(),
// 			Msg:    "Variable '" + idName + "' not declared",
// 			Type:   "VariableError",
// 		})
// 	}

// 	return nil

// }
