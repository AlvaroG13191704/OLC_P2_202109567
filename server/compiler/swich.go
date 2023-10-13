package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

func (v *Visitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {
	// temps
	var listLabelCases []string
	startLabel := v.Generator.NewLabel()
	labelEnd := v.Generator.NewLabel()
	var defaultLabel string

	if ctx.DefaultBlock() != nil {
		// create a new label
		defaultLabel = v.Generator.NewLabel()
	}
	// add comment
	v.Generator.GenComment("switch condition")
	v.PushLoopContext("switch", startLabel, labelEnd)
	defer v.PopLoopContext() // pop the loop context after the execution

	// get the condition
	conditionExpr := v.Visit(ctx.Expr()).(*values.C3DPrimitive)

	// get list of cases
	casesList := ctx.AllCaseBlock()

	// iterate to create the labels
	for range casesList {
		// create a new label
		listLabelCases = append(listLabelCases, v.Generator.NewLabel())
	}

	// iterate over the cases
	for i, caseBlock := range casesList {

		// gen comment
		v.Generator.GenComment("case condition")
		// add label
		v.Generator.AddLabel(listLabelCases[i])

		// get the case condition
		caseCondition := v.Visit(caseBlock.Expr()).(*values.C3DPrimitive)
		// verify if the cases are the same type
		if caseCondition.GetType() == conditionExpr.GetType() {

			// verify if it's the end of the list
			if i == len(casesList)-1 {
				// evaluate if there is a default block
				if ctx.DefaultBlock() != nil {
					// generate if
					v.Generator.AddIf(conditionExpr.GetValue(), caseCondition.GetValue(), "!=", defaultLabel)
				} else {
					// generate if
					v.Generator.AddIf(conditionExpr.GetValue(), caseCondition.GetValue(), "!=", labelEnd)
				}
			} else {
				// generate if
				v.Generator.AddIf(conditionExpr.GetValue(), caseCondition.GetValue(), "!=", listLabelCases[i+1])
			}
			// visit block
			v.Visit(caseBlock.Block())
			// go to end
			v.Generator.GoTo(labelEnd)

		} else {
			log.Fatal("Error: switch cases must be the same type")
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Error: switch cases must be the same type",
				Type:   "Semantic",
			})
		}

	}

	// if there is a default block
	if ctx.DefaultBlock() != nil {
		// add label
		v.Generator.AddLabel(defaultLabel)
		// return the default block
		v.Visit(ctx.DefaultBlock())
		// go to end
		v.Generator.GoTo(labelEnd)
	}

	// ADD LABEL END
	v.Generator.AddLabel(labelEnd)

	fmt.Printf("LoopStack: %v\n", v.loopContexts)

	return nil

}

// visit default block
func (v *Visitor) VisitDefaultBlock(ctx *parser.DefaultBlockContext) interface{} {

	return v.Visit(ctx.Block())
}
