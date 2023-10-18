package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
	"strconv"
)

// VisitAccessVector
func (v *Visitor) VisitAccessVector(ctx *parser.AccessVectorContext) interface{} {
	// gen comment
	v.Generator.GenComment("Access Vector")

	// get the vector name
	vectorName := ctx.ID_PRIMITIVE().GetText()

	// evaluate if the vector exists
	value, ok := v.VerifyScope(vectorName)
	if !ok {
		log.Printf("Error: Variable '%s' not declared \n", vectorName)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' not declared", vectorName),
			Type:   "Semantic",
		})

		return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
	}
	symbolValue := value.(Symbol)

	// get the index
	index := v.Visit(ctx.Expr()).(*values.C3DPrimitive)

	// conver to int
	indexInt, _ := strconv.ParseInt(index.GetValue(), 10, 64)

	// evaluate if the index is out of bounds
	if indexInt < 0 || indexInt >= int64(symbolValue.LenghtVector) {
		log.Printf("Error: Index '%d' is out of bounds \n", index.GetValue())
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Index '%d' is out of bounds", index.GetValue()),
			Type:   "Semantic",
		})
		return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
	}

	// evaluate if the index is a number
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

	// primitive
	tempIterator := v.Generator.NewTemp()
	tempValue := v.Generator.NewTemp()
	tempLenght := v.Generator.NewTemp()
	// get the value from the stack
	v.Generator.GetStack(tempIterator, fmt.Sprintf("%d", symbolValue.StackDirection))
	// increment the iterator
	v.Generator.GenArithmetic(tempLenght, tempIterator, "1", "+")

	// sum the index to the iterator
	v.Generator.GenArithmetic(tempIterator, tempIterator, fmt.Sprintf("%d", indexInt+1), "+")
	// get the value from the heap
	v.Generator.GetHeap(tempValue, tempIterator)

	return values.NewC3DPrimitive(tempValue, tempIterator, symbolValue.TypeData, true)
}

// VisitAppendVector -> append an expression to an existing vector
func (v *Visitor) VisitAppendVector(ctx *parser.AppendVectorContext) interface{} {
	// get the vector name
	vectorName := ctx.ID_PRIMITIVE().GetText()

	// evaluate if the vector exists
	value, ok := v.VerifyScope(vectorName)
	if !ok {
		log.Printf("Error: Variable '%s' not declared \n", vectorName)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' not declared", vectorName),
			Type:   "Semantic",
		})
		return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
	}

	symbolValue := value.(Symbol)
	// fmt.Println("Symbol value before -> ", symbolValue)

	// get the expression
	expr := v.Visit(ctx.Expr()).(*values.C3DPrimitive)

	// evaluate if the expr type is the same as the symbol value
	if symbolValue.TypeData != expr.GetType() {
		log.Printf("Error: Type '%s' is not assignable to type '%s' \n", expr.GetType(), symbolValue.TypeData)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Type '%s' is not assignable to type '%s'", expr.GetType(), symbolValue.TypeData),
			Type:   "Semantic",
		})
		return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
	}

	var TempsString string
	if expr.GetType() == values.StringType {
		tempString := v.Generator.GenString(expr.GetValue())
		// add to the list of temps
		TempsString = tempString
	}
	// append the value to the vector
	v.Generator.GenComment("Append Vector")

	// c3d code
	// generate temps
	tempLenght := v.Generator.NewTemp()
	tempPointerStack := v.Generator.NewTemp()
	tempPointerHeap := v.Generator.NewTemp()
	newTempPointer := v.Generator.NewTemp()

	// get the value from the stack
	v.Generator.GetStack(tempPointerStack, fmt.Sprintf("%d", symbolValue.StackDirection))

	// get the value from the heap
	v.Generator.GetHeap(tempPointerHeap, tempPointerStack)

	// assign the value to the tempLenght
	v.Generator.AssignTemp(tempLenght, tempPointerHeap)

	// increment the and the lenght
	v.Generator.GenArithmetic(tempLenght, tempLenght, "1", "+")            // the new lenght because is going to be added a new value
	v.Generator.GenArithmetic(tempPointerHeap, tempPointerStack, "1", "+") // start the values of the vector

	// gencomment
	v.Generator.GenComment("Save length of the vector")
	// iterate the vector to save it in a new position
	v.Generator.AssignTemp(newTempPointer, "H")
	// save the length of the vector in the first position
	v.Generator.SaveHeap(tempLenght)
	// increase the heap
	v.Generator.IncPointerHeap()

	// iterate the vector
	iteratorTemp := v.Generator.NewTemp()
	v.Generator.GenComment("Start saving vector values")
	for i := 0; i < symbolValue.LenghtVector; i++ {
		// get the value from the heap
		v.Generator.GetHeap(iteratorTemp, tempPointerHeap)
		// save the value in the heap
		v.Generator.SaveHeap(iteratorTemp)
		// increment the iterator
		v.Generator.GenArithmetic(tempPointerHeap, tempPointerHeap, "1", "+")
		// increase the heap
		v.Generator.IncPointerHeap()
	}

	if expr.GetType() == values.StringType {
		// save the value in the heap
		v.Generator.SaveHeap(TempsString)

		// increase the heap
		v.Generator.IncPointerHeap()
	} else {

		// save the value in the heap
		v.Generator.SaveHeap(expr.GetValue())

		// increase the heap
		v.Generator.IncPointerHeap()
	}

	// declare
	stack, heap := v.Generator.GenDeclarationVector(newTempPointer, symbolValue.Id, false)

	// update the symbol table
	symbolValue.Value = newTempPointer
	symbolValue.StackDirection = stack
	symbolValue.HeapDirection = heap
	symbolValue.LenghtVector = symbolValue.LenghtVector + 1

	// fmt.Println("Symbol value after -> ", symbolValue)

	fmt.Println("Global scope or symbol table ->", v.SymbolStack)
	// update the symbol table
	v.UpdateVariable(vectorName, symbolValue)

	// increment the stack
	v.Generator.CounterStack("+")
	return nil
}

// VisitRemoveLastVector
func (v *Visitor) VisitRemoveLastVector(ctx *parser.RemoveLastVectorContext) interface{} {
	// gen comment
	v.Generator.GenComment("Remove Last Vector")
	// get the vector name
	vectorName := ctx.ID_PRIMITIVE().GetText()

	// evaluate if the vector exists
	value, ok := v.VerifyScope(vectorName)
	if !ok {
		log.Printf("Error: Variable '%s' not declared \n", vectorName)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' not declared", vectorName),
			Type:   "Semantic",
		})
		return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
	}
	symbolValue := value.(Symbol)
	// evaluate if the vector is empty
	if symbolValue.LenghtVector == 0 {
		log.Printf("Error: Vector '%s' is empty \n", vectorName)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Vector '%s' is empty", vectorName),
			Type:   "Semantic",
		})
		return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
	}

	// heap[INICIO_VECTOR + LONGITUD_VECTOR + 1]
	// c3d code
	// generate temps
	tempLenght := v.Generator.NewTemp()
	tempPointerStack := v.Generator.NewTemp()
	tempPointerHeap := v.Generator.NewTemp()

	// get the value from the stack
	v.Generator.GetStack(tempPointerStack, fmt.Sprintf("%d", symbolValue.StackDirection))

	// get the value from the heap
	v.Generator.GetHeap(tempPointerHeap, tempPointerStack)

	// assign the value to the tempLenght
	v.Generator.AssignTemp(tempLenght, tempPointerHeap)

	// increment the and the lenght
	v.Generator.GenArithmetic(tempLenght, tempLenght, "1", "-") // the new lenght because is going to be added a new value
	// gencomment
	v.Generator.GenComment("re-assign length of the vector")
	// save the length of the vector in the first position
	v.Generator.SaveHeapIndex(tempLenght, fmt.Sprintf("%d", symbolValue.StackDirection))

	// gen comment
	v.Generator.GenComment("set 0 last value")
	// set t3 = t3 + lenght + 1
	v.Generator.GenArithmetic(tempPointerHeap, tempPointerHeap, fmt.Sprintf("%d", symbolValue.LenghtVector+1), "+")
	// set 0 in the last position
	v.Generator.SaveHeapIndex("0", tempPointerHeap)

	// update the symbol table
	symbolValue.LenghtVector = symbolValue.LenghtVector - 1

	// decrease the heap
	v.Generator.DecPointerHeap()

	// update the symbol table
	v.UpdateVariable(vectorName, symbolValue)

	return nil
}

// VisitIsEmptyVector
func (v *Visitor) VisitIsEmptyVector(ctx *parser.IsEmptyVectorContext) interface{} {
	// gen comment
	v.Generator.GenComment("Is Empty Vector")
	// get the vector name
	vectorName := ctx.ID_PRIMITIVE().GetText()

	// evaluate if the vector exists
	value, ok := v.VerifyScope(vectorName)
	if !ok {
		log.Printf("Error: Variable '%s' not declared \n", vectorName)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' not declared", vectorName),
			Type:   "Semantic",
		})
		return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
	}
	symbolValue := value.(Symbol)

	// evaluate if the vector is empty
	// c3d code
	// generate temps
	tempValue := v.Generator.NewTemp()
	tempPointerStack := v.Generator.NewTemp()
	tempPointerHeap := v.Generator.NewTemp()
	// generate label
	labelEnd := v.Generator.NewLabel()

	// get the value from the stack
	v.Generator.GetStack(tempPointerStack, fmt.Sprintf("%d", symbolValue.StackDirection))

	// get the value from the heap
	v.Generator.GetHeap(tempPointerHeap, tempPointerStack)

	// assing 1 to the temp value
	v.Generator.AssignTemp(tempValue, "0")
	// evaluate if the lenght is 0
	v.Generator.AddIf(tempPointerHeap, "0", "!=", labelEnd)
	// if the lenght is 0
	v.Generator.AssignTemp(tempValue, "1")
	// add goto
	v.Generator.AddLabel(labelEnd)

	return values.NewC3DPrimitive(tempValue, "nil", values.BooleanType, true)
}

// VisitCountVector
func (v *Visitor) VisitCountVector(ctx *parser.CountVectorContext) interface{} {
	// gen comment
	v.Generator.GenComment("Count Vector")
	// get the vector name
	vectorName := ctx.ID_PRIMITIVE().GetText()

	// evaluate if the vector exists
	value, ok := v.VerifyScope(vectorName)
	if !ok {
		log.Printf("Error: Variable '%s' not declared \n", vectorName)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' not declared", vectorName),
			Type:   "Semantic",
		})
		return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
	}
	symbolValue := value.(Symbol)

	// c3d code
	// generate temps
	tempPointerStack := v.Generator.NewTemp()
	tempPointerHeap := v.Generator.NewTemp()

	// get the value from the stack
	v.Generator.GetStack(tempPointerStack, fmt.Sprintf("%d", symbolValue.StackDirection))

	// get the value from the heap
	v.Generator.GetHeap(tempPointerHeap, tempPointerStack)

	return values.NewC3DPrimitive(tempPointerHeap, "nil", values.IntType, true)
}

// VisitRemoveAtVector
func (v *Visitor) VisitRemoveAtVector(ctx *parser.RemoveAtVectorContext) interface{} {
	// gen comment
	v.Generator.GenComment("Remove At Vector")
	// get the vector name
	vectorName := ctx.ID_PRIMITIVE().GetText()

	// evaluate if the vector exists
	value, ok := v.VerifyScope(vectorName)
	if !ok {
		log.Printf("Error: Variable '%s' not declared \n", vectorName)
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Variable '%s' not declared", vectorName),
			Type:   "Semantic",
		})
		return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
	}
	symbolValue := value.(Symbol)

	// get the index
	index := v.Visit(ctx.Expr()).(*values.C3DPrimitive)

	// evaluate if the index is a number
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

	// convert the index from int64 to int
	numINdex, _ := strconv.ParseInt(index.GetValue(), 10, 64)

	// evaluate if the index is out of bounds
	if numINdex < 0 || numINdex >= int64(symbolValue.LenghtVector) {
		log.Printf("Error: Index '%d' is out of bounds \n", index.GetValue())
		// add error
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    fmt.Sprintf("Error: Index '%d' is out of bounds", index.GetValue()),
			Type:   "Semantic",
		})
		return values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false)
	}
	// c3d code
	// generate temps
	tempLenght := v.Generator.NewTemp()
	tempPointerStack := v.Generator.NewTemp()
	tempPointerHeap := v.Generator.NewTemp()
	newTempPointer := v.Generator.NewTemp()

	// get the value from the stack
	v.Generator.GetStack(tempPointerStack, fmt.Sprintf("%d", symbolValue.StackDirection))

	// get the value from the heap
	v.Generator.GetHeap(tempPointerHeap, tempPointerStack)

	// assign the value to the tempLenght
	v.Generator.AssignTemp(tempLenght, tempPointerHeap)

	// increment the and the lenght
	v.Generator.GenArithmetic(tempLenght, tempLenght, "1", "-")            // the new lenght because is going to be added a new value
	v.Generator.GenArithmetic(tempPointerHeap, tempPointerStack, "1", "+") // start the values of the vector

	// gencomment
	v.Generator.GenComment("Save length of the vector")
	// iterate the vector to save it in a new position
	v.Generator.AssignTemp(newTempPointer, "H")
	// save the length of the vector in the first position
	v.Generator.SaveHeap(tempLenght)
	// increase the heap
	v.Generator.IncPointerHeap()
	// iterate the vector
	iteratorTemp := v.Generator.NewTemp()
	v.Generator.GenComment("Start saving new vector values")
	for i := 0; i < symbolValue.LenghtVector; i++ {
		if i == int(numINdex) {
			// increment the iterator
			v.Generator.GenArithmetic(tempPointerHeap, tempPointerHeap, "1", "+")
			continue
		}
		// get the value from the heap
		v.Generator.GetHeap(iteratorTemp, tempPointerHeap)
		// save the value in the heap
		v.Generator.SaveHeap(iteratorTemp)
		// increment the iterator
		v.Generator.GenArithmetic(tempPointerHeap, tempPointerHeap, "1", "+")
		// increase the heap
		v.Generator.IncPointerHeap()
	}

	// declare
	stack, heap := v.Generator.GenDeclarationVector(newTempPointer, symbolValue.Id, false)

	// update the symbol table
	symbolValue.Value = newTempPointer
	symbolValue.StackDirection = stack
	symbolValue.HeapDirection = heap
	symbolValue.LenghtVector = symbolValue.LenghtVector - 1

	// fmt.Println("Symbol value after -> ", symbolValue)

	fmt.Println("Global scope or symbol table ->", v.SymbolStack)
	// update the symbol table
	v.UpdateVariable(vectorName, symbolValue)
	// decrease the heap
	v.Generator.DecPointerHeap()
	// increment the stack
	v.Generator.CounterStack("+")
	return nil

}
