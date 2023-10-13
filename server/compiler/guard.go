package compiler

import (
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

func (v *Visitor) VisitGuardStmt(ctx *parser.GuardStmtContext) interface{} {

	// get the condition
	conditionExpr := v.Visit(ctx.Expr()).(*values.C3DPrimitive)
	// verify if the condition is a Bool
	if conditionExpr.GetType() != values.BooleanType {
		// add error
		log.Printf("Error: Condition must be a boolean in if  \n")
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Error: Condition must be a boolean",
			Type:   "Semantic",
		})
		return nil
	}

	// c3d code
	labelEnd := v.Generator.NewLabel()
	// add comment
	v.Generator.GenComment("Guard condition")
	// generate if
	v.Generator.AddIf("(int)"+conditionExpr.GetValue(), "1", "==", labelEnd)
	// visit block
	v.Visit(ctx.Block())
	// go to end
	v.Generator.GoTo(labelEnd)
	// add label end
	v.Generator.AddLabel(labelEnd)

	// return
	return nil
}
