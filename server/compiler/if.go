package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

// VisitIfElseStmt
func (v *Visitor) VisitIfElseStmt(ctx *parser.IfElseStmtContext) interface{} {
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
	labelTrue := v.Generator.NewLabel()
	labelFalse := v.Generator.NewLabel()
	fmt.Println("Condition: ", conditionExpr)
	// add comment
	v.Generator.GenComment("If condition")
	// generate if
	v.Generator.AddIf("(int)"+conditionExpr.GetValue(), "1", "==", labelTrue)
	v.Generator.GoTo(labelFalse)

	// // evaluate the condition
	// if conditionExpr.GetValue().(bool) {
	// 	// return the block
	// 	return v.Visit(ctx.Block(0))

	// } else if ctx.ELSE() != nil {
	// 	// return the else block
	// 	return v.Visit(ctx.Block(1))
	// }

	return nil
}
