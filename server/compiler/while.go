package compiler

import (
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

func (v *Visitor) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {
	// temps
	startLabel := v.Generator.NewLabel()
	labelEnd := v.Generator.NewLabel()
	// add comment
	v.Generator.GenComment("While condition")
	// labels
	// add label
	v.Generator.AddLabel(startLabel)

	// push a loop to the loop context
	v.PushLoopContext("while", startLabel, labelEnd)
	defer v.PopLoopContext() // pop the loop context after the execution

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

	// generate if
	v.Generator.AddIf("(int)"+conditionExpr.GetValue(), "0", "==", labelEnd)
	// generate the block
	v.Visit(ctx.Block())
	// go to start
	v.Generator.GoTo(startLabel)

	// go to end
	v.Generator.AddLabel(labelEnd)

	return nil
}
