package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

func (v *Visitor) VisitValueAssignment(ctx *parser.ValueAssignmentContext) interface{} {

	// get the variable name
	idName := ctx.ID_PRIMITIVE().GetText()
	// get the expr
	expr := v.Visit(ctx.Expr()).(*values.C3DPrimitive)
	// get the value
	valueFromScope, ok := v.VerifyScope(idName)
	// evaluate if the name is declared
	if ok {
		// get the value
		value := valueFromScope.(Symbol)
		// evaluate if its a constant
		if value.TypeMutable == "let" {
			log.Fatalf("Error: Variable '%s' is a constant, cannot be assign", idName)
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Variable '" + idName + "' is a constant",
				Type:   "VariableError",
			})
			return nil
		}
		// get the type of the variable

		typeVariable := value.Value.(*values.C3DPrimitive).GetType()
		// evaluate if the type of the variable is the same of the expr
		if typeVariable == expr.GetType() || typeVariable == values.NilType {
			// gen c3d
			fmt.Println("value assignment -> ", value)
			v.Generator.GenAssignment(
				value.Value.(*values.C3DPrimitive).GetValue(),
				idName, value.StackDirection,
				value.Value.(*values.C3DPrimitive),
				expr,
			)
			// update the value
			value.Value = expr

			// update the value no matter the scope
			v.UpdateVariable(idName, value)

			fmt.Println("Global scope or symbol table ->", v.SymbolStack)
			return nil
		} else {
			log.Fatalf("Error: Variable '%s' is of type %s, cannot be assign to %s", idName, typeVariable, expr.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Variable '" + idName + "' is of type " + typeVariable + ", cannot be assign to " + expr.GetType(),
				Type:   "TypeError",
			})
			return nil
		}

	} else {
		// add the error to the errors
		log.Fatalf("Error: Variable '%s' not declared", idName)
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Variable '" + idName + "' not declared",
			Type:   "VariableError",
		})
	}

	return nil

}

func (v *Visitor) VisitPlusAssignment(ctx *parser.PlusAssignmentContext) interface{} {
	// get the variable name
	idName := ctx.ID_PRIMITIVE().GetText()
	// get the expr
	expr := v.Visit(ctx.Expr()).(*values.C3DPrimitive)
	// get the value
	valueFromScope, ok := v.VerifyScope(idName)
	// evaluate if the name is declared
	if ok {
		// get the value
		value := valueFromScope.(Symbol)
		// evaluate if its a constant
		if value.TypeMutable == values.LetMutable {
			log.Fatalf("Error: Variable '%s' is a constant, cannot be assign", idName)
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Variable '" + idName + "' is a constant",
				Type:   "VariableError",
			})
			return nil
		}
		// get the type of the variable

		typeVariable := value.Value.(*values.C3DPrimitive).GetType()
		// evaluate if the type of the variable is the same of the expr
		// if the expr is an integer, cast it to float
		if typeVariable == expr.GetType() || (typeVariable == values.FloatType && expr.GetType() == values.IntType) || (typeVariable == values.IntType && expr.GetType() == values.FloatType) {

			fmt.Println("value -> ", value)
			// gen c3d
			v.Generator.GenAssignmentPlusMiuns(
				"+",
				fmt.Sprintf("%d", value.StackDirection),
				expr,
			)
			// update the value no matter the scope
			v.UpdateVariable(idName, value)

			fmt.Println("Global scope or symbol table ->", v.SymbolStack)

			return nil
		} else {
			log.Fatalf("Error: Variable '%s' is of type %s, cannot be assign to %s", idName, typeVariable, expr.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Variable '" + idName + "' is of type " + typeVariable + ", cannot be assign to " + expr.GetType(),
				Type:   "TypeError",
			})
			return nil
		}

	} else {
		// add the error to the errors
		log.Fatalf("Error: Variable '%s' not declared", idName)
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Variable '" + idName + "' not declared",
			Type:   "VariableError",
		})
	}

	return nil
}

func (v *Visitor) VisitMinusAssignment(ctx *parser.MinusAssignmentContext) interface{} {
	// get the variable name
	idName := ctx.ID_PRIMITIVE().GetText()
	// get the expr
	expr := v.Visit(ctx.Expr()).(*values.C3DPrimitive)
	// get the value
	valueFromScope, ok := v.VerifyScope(idName)
	// evaluate if the name is declared
	if ok {
		// get the value
		value := valueFromScope.(Symbol)
		// evaluate if its a constant
		if value.TypeMutable == "let" {
			log.Fatalf("Error: Variable '%s' is a constant, cannot be assign", idName)
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Variable '" + idName + "' is a constant",
				Type:   "VariableError",
			})
			return nil
		}
		// get the type of the variable

		typeVariable := value.Value.(*values.C3DPrimitive).GetType()
		// evaluate if the type of the variable is the same of the expr
		// if the expr is an integer, cast it to float
		if typeVariable == expr.GetType() || (typeVariable == values.FloatType && expr.GetType() == values.IntType) || (typeVariable == values.IntType && expr.GetType() == values.FloatType) {

			// gen c3d
			v.Generator.GenAssignmentPlusMiuns(
				"-",
				fmt.Sprintf("%d", value.StackDirection),
				expr,
			)
			// update the value no matter the scope
			v.UpdateVariable(idName, value)

			return nil
		} else {
			log.Fatalf("Error: Variable '%s' is of type %s, cannot be assign to %s", idName, typeVariable, expr.GetType())
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Variable '" + idName + "' is of type " + typeVariable + ", cannot be assign to " + expr.GetType(),
				Type:   "TypeError",
			})
			return nil
		}

	} else {
		// add the error to the errors
		log.Fatalf("Error: Variable '%s' not declared", idName)
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Variable '" + idName + "' not declared",
			Type:   "VariableError",
		})
	}

	return nil
}
