package compiler

import (
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

func (v *Visitor) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {

	// push a loop to the loop context
	v.PushLoopContext("while")
	defer v.PopLoopContext() // pop the loop context after the execution

	// c3d code
	// labels
	startLabel := v.Generator.NewLabel()
	// add label
	v.Generator.AddLabel(startLabel)

	// get the condition
	conditionExpr := v.Visit(ctx.Expr()).(*values.C3DPrimitive)
	// verify if the condition is a Bool
	if conditionExpr.GetType() != values.BooleanType {
		// add error
		log.Printf("Error: Condition must be a boolean in while  \n")
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Error: Condition must be a boolean",
			Type:   "Semantic",
		})
		return nil
	}

	// labels
	labelTrue := v.Generator.NewLabel()
	labelFalse := v.Generator.NewLabel()
	labelEnd := v.Generator.NewLabel()

	// add comment
	v.Generator.GenComment("While condition")
	// generate if
	v.Generator.AddIf("(int)"+conditionExpr.GetValue(), "1", "==", labelTrue)
	v.Generator.GoTo(labelFalse)
	v.Generator.AddLabel(labelTrue)
	// visit block 1
	v.Visit(ctx.Block())

	// TODO: evaluate if there is a continue

	// evaluate if there is a break in the loop
	if isBreak := v.GetLoopContext().BreakFound; isBreak {
		// add comment
		v.Generator.GenComment("Break stmt")
		// add label end
		v.Generator.GoTo(labelEnd)
	}
	// go to end
	v.Generator.GoTo(startLabel)
	// add label false
	v.Generator.AddLabel(labelFalse)
	// add label end
	v.Generator.AddLabel(labelEnd)

	return nil
}
