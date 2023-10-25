package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
	"strconv"
)

// VisitVectorAssignment
func (v *Visitor) VisitVectorAssignment(ctx *parser.VectorAssignmentContext) interface{} {
	// get the variable name
	idName := ctx.ID_PRIMITIVE().GetText()
	// get the value
	valueFromScope, ok := v.VerifyScope(idName)
	// evaluate if the name is declared
	if !ok {
		// add the error to the errors
		log.Fatalf("Error: Variable '%s' not declared", idName)
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Variable '" + idName + "' not declared",
			Type:   "VariableError",
		})
		return nil
	}
	symbolValue := valueFromScope.(Symbol)
	var numINdex int64
	index := v.Visit(ctx.Expr(0)).(*values.C3DPrimitive)
	if index.GetType() != values.IntType {
		log.Printf("Error: Index '%s' is not a number \n", index.GetValue())
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Index '%s' is not a number", index.GetValue()),
			Type:   "Semantic",
		})
		return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
	}

	if !index.GetIsTemp() {
		// conver to int
		numINdex, _ = strconv.ParseInt(index.GetValue(), 10, 64)
		// evaluate if the index is out of bounds
		if numINdex < 0 || numINdex >= int64(symbolValue.LenghtVector) {
			log.Printf("Error: Index '%d' is out of bounds \n", numINdex)
			// add error
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    fmt.Sprintf("Error: Index '%d' is out of bounds", numINdex),
				Type:   "Semantic",
			})
			return nil
		}
	}

	// get the expr
	expr := v.Visit(ctx.Expr(1)).(*values.C3DPrimitive)

	// evaluate if the expr type is the same as the symbol value
	if expr.GetType() != symbolValue.TypeData {
		log.Printf("Error: Type of the expression is different from the type of the vector")
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Error: Type of the expression is different from the type of the vector",
			Type:   "Semantic",
		})
		return nil
	}
	var tempString string
	// if the expr is a string
	if expr.GetType() == values.StringType {
		// get the temp
		tempString = v.Generator.GenString(expr.GetValue())
	} else {
		tempString = expr.GetValue()
	}

	// gen comment
	v.Generator.GenComment("Vector assignment")
	// gen c3d
	tempPointerStack := v.Generator.NewTemp()
	tempPointerHeap := v.Generator.NewTemp()
	// newTempPointer := v.Generator.NewTemp()

	// get the value from the stack
	v.Generator.GetStack(tempPointerStack, fmt.Sprintf("%d", symbolValue.StackDirection))

	// increment to move to the length of the vector
	v.Generator.GenArithmetic(tempPointerHeap, tempPointerStack, "1", "+") // start the values of the vector

	// increment to move to the index of the vector
	if !index.GetIsTemp() {
		v.Generator.GenArithmetic(tempPointerHeap, tempPointerStack, fmt.Sprintf("%d", numINdex+1), "+") // start the values of the vector
	} else {
		v.Generator.GenArithmetic(tempPointerHeap, tempPointerStack, index.GetValue(), "+") // start the values of the vector
	}

	// assign the value to the vector
	v.Generator.SaveHeapIndex(tempString, tempPointerHeap)

	// update the value no matter the scope
	v.UpdateVariable(idName, symbolValue)

	return nil
}
