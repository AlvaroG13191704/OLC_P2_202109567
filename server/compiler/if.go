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
	labelEnd := v.Generator.NewLabel()
	fmt.Println("Condition: ", conditionExpr)
	// add comment
	v.Generator.GenComment("If else condition")
	// generate if
	v.Generator.AddIf("(int)"+conditionExpr.GetValue(), "1", "==", labelTrue)
	v.Generator.GoTo(labelFalse)
	v.Generator.AddLabel(labelTrue)
	// visit block 1
	v.Visit(ctx.Block(0))
	// go to end
	v.Generator.GoTo(labelEnd)
	// add label false
	v.Generator.AddLabel(labelFalse)
	// visit block 2
	if ctx.Block(1) != nil {
		v.Visit(ctx.Block(1))
	}
	// add label end
	v.Generator.AddLabel(labelEnd)

	return nil
}

// VisitElseIfStmt
func (v *Visitor) VisitElseIfStmt(ctx *parser.ElseIfStmtContext) interface{} {
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
	labelEnd := v.Generator.NewLabel()
	fmt.Println("Condition: ", conditionExpr)
	// add comment
	v.Generator.GenComment("Else if condition")
	// generate if
	v.Generator.AddIf("(int)"+conditionExpr.GetValue(), "1", "==", labelTrue)
	v.Generator.GoTo(labelFalse)
	v.Generator.AddLabel(labelTrue)
	// visit block 1
	v.Visit(ctx.Block())
	// go to end
	v.Generator.GoTo(labelEnd)
	// add label false
	v.Generator.AddLabel(labelFalse)
	// visit block 2
	if ctx.Ifstmt() != nil {
		v.Visit(ctx.Ifstmt())
	}
	// add label end
	v.Generator.AddLabel(labelEnd)

	return nil
}
