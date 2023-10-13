package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
	"strconv"
)

// visit for statement
func (v *Visitor) VisitForRangeExpr(ctx *parser.ForRangeExprContext) interface{} {
	// get the range
	rangeValue := v.Visit(ctx.ForRange())
	// gett he left and right value
	leftValue := rangeValue.([]*values.C3DPrimitive)[0]
	rightValue := rangeValue.([]*values.C3DPrimitive)[1]
	// declare the loop variable
	var loopVarName string = "_"
	if ctx.ID_PRIMITIVE() != nil {
		// get the loopVarName
		loopVarName = ctx.ID_PRIMITIVE().GetText()
	}
	// declare the variable c3d
	// generate temp
	stack, heap := v.Generator.GenDeclaration(leftValue.GetTemp(), loopVarName, leftValue)

	// increment the stack
	v.Generator.CounterStack("+")
	// assign the value to the loopVarName
	v.AssignVariable(loopVarName, Symbol{
		Id:             loopVarName,
		TypeSymbol:     values.Variable_Type,
		TypeMutable:    values.LetMutable,
		TypeData:       leftValue.GetType(),
		Value:          leftValue,
		Line:           ctx.GetStart().GetLine(),
		Column:         ctx.GetStart().GetColumn(),
		StackDirection: stack,
		HeapDirection:  heap,
	})

	// temps
	startLabel := v.Generator.NewLabel()
	labelEnd := v.Generator.NewLabel()
	v.LabelLoop = labelEnd
	// add comment
	v.Generator.GenComment("--------- For range condition ---------")
	// labels
	// add label
	v.Generator.AddLabel(startLabel)
	// push a loop to the loop context
	v.PushLoopContext("for", startLabel, labelEnd)
	defer v.PopLoopContext() // pop the loop context after the execution

	// verify if the left and right value are integers
	if leftValue.GetType() == values.NilType || rightValue.GetType() == values.NilType {
		log.Printf("Error: %s", "left and right value must be an integer")
		// add the error to the errors
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "left and right value in the range must be an Integer",
			Type:   "TypeError",
		})
		return nil
	}

	// generate if
	v.Generator.AddIf(leftValue.GetTemp(), rightValue.GetTemp(), ">", labelEnd)
	// visit block
	v.Visit(ctx.Block())
	// increment the loop variable
	v.Generator.GenArithmetic(leftValue.GetTemp(), leftValue.GetTemp(), "1", "+")
	// assign the value
	v.Generator.SaveStack(fmt.Sprintf("%d", stack), leftValue.GetTemp())
	// go to start
	v.Generator.GoTo(startLabel)
	// go to end
	v.Generator.AddLabel(labelEnd)
	v.Generator.GenComment("--------- END For range condition ---------")
	// delete the variable from the scope
	v.DeleteVariable(loopVarName)

	fmt.Println("Global scope or symbol table ->", v.SymbolStack)

	return nil
}

// visit forRange
func (v *Visitor) VisitForRange(ctx *parser.ForRangeContext) interface{} {
	// get left expr
	leftExpr := v.Visit(ctx.GetLeft()).(*values.C3DPrimitive)
	rightExpr := v.Visit(ctx.GetRight()).(*values.C3DPrimitive)

	fmt.Println("leftExpr: ", leftExpr)
	fmt.Println("rightExpr: ", rightExpr)

	// convert string to int
	leftExprInt, _ := strconv.ParseInt(leftExpr.GetValue(), 10, 64)
	rightExprInt, _ := strconv.ParseInt(rightExpr.GetValue(), 10, 64)

	// evaluate if the left expr is an int and the right expr is an int
	if leftExpr.GetType() == values.IntType && rightExpr.GetType() == values.IntType {
		if leftExprInt > rightExprInt {
			// throw an error
			log.Printf("Error: %s", "left expr must be less than right expr")
			// add the error to the errors
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "left expr must be less than right expr",
				Type:   "TypeError",
			})
			return []*values.C3DPrimitive{
				values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false),
				values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false),
			}
		}
		// add the c3d
		v.Generator.GenComment("--------- for range --------- ")
		// temps
		temp1 := v.Generator.NewTemp()
		temp2 := v.Generator.NewTemp()
		// add the c3d
		v.Generator.AssignTemp(temp1, leftExpr.GetValue())
		v.Generator.AssignTemp(temp2, rightExpr.GetValue())
		v.Generator.GenComment("--------- end for range --------- ")

		// return an array of Primitive Intenger struct
		return []*values.C3DPrimitive{
			values.NewC3DPrimitive(leftExpr.GetValue(), temp1, values.IntType, true),
			values.NewC3DPrimitive(rightExpr.GetValue(), temp2, values.IntType, true),
		}

	} else {
		return []*values.C3DPrimitive{
			values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false),
			values.NewC3DPrimitive("9999999827968.00", "nil", values.NilType, false),
		}
	}
}
