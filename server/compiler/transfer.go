package compiler

import (
	"fmt"
	"log"
	"server/compiler/compiler/values"
	"server/compiler/parser"
)

func (v *Visitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {
	// evaluate if the expression is not null
	fmt.Println("RETURN STMT")
	// if there is return value ++ the size
	v.SizeFunction = v.SizeFunction + 1
	// evaluate if the return is inside a function
	if !v.ExistsFunctionContext() {
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Return statement must be inside a function",
			Type:   "Semantic",
		})
		log.Println("Return statement must be inside a function")
		return nil
	}
	if ctx.Expr() != nil {
		// evaluate the expression
		expr := v.Visit(ctx.Expr()).(*values.C3DPrimitive)

		if !expr.GetIsTemp() && expr.GetType() == values.StringType {
			// add comment
			v.Generator.GenComment("Return statement")
			// gen string
			tempString := v.Generator.GenString(expr.GetValue())
			// assign to the return type
			v.Generator.AssignTemp(v.ReturnTemp, tempString)
			// add goto
			v.Generator.GoTo(v.ReturnLabel)
			return nil
		}

		// gen comment
		v.Generator.GenComment("Return statement")

		// assign to the return type
		v.Generator.AssignTemp(v.ReturnTemp, expr.GetValue())

		// add goto
		v.Generator.GoTo(v.ReturnLabel)

	} else {

	}
	return nil
}

// VisitBreak
func (v *Visitor) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {

	// evaluate if the current context is inside a loop
	if v.ExistsLoopContext() {

		// update the loop context
		loopContext := v.GetLoopContext()

		if loopContext.TypeLoop == "while" || loopContext.TypeLoop == "for" || loopContext.TypeLoop == "switch" {
			// update the loop context
			loopContext.BreakFound = true

			// update the loop context
			v.UpdateLoopContext(loopContext)

			// add comment
			v.Generator.GenComment("Break statement")
			// go to end
			v.Generator.GoTo(loopContext.EndLabel)

		} else {
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Break statement must be inside a loop",
				Type:   "Semantic",
			})
			log.Println("Break statement must be inside a loop")

		}

	} else {
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Break statement must be inside a loop",
			Type:   "Semantic",
		})
		log.Println("Break statement must be inside a loop")
	}

	return nil
}

// VisitContinue
func (v *Visitor) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {

	// evaluate if the current context is inside a loop
	if v.ExistsLoopContext() {
		fmt.Println("continue found--------")
		// update the loop context
		loopContext := v.GetLoopContext()

		if loopContext.TypeLoop == "while" || loopContext.TypeLoop == "for" {
			// update the loop context
			loopContext.ContinueFound = true

			// update the loop context
			v.UpdateLoopContext(loopContext)

			// add comment
			v.Generator.GenComment("Continue statement")
			// go to start
			v.Generator.GoTo(loopContext.StartLabel)

		} else {
			v.Errors = append(v.Errors, Error{
				Line:   ctx.GetStart().GetLine(),
				Column: ctx.GetStart().GetColumn(),
				Msg:    "Continue statement must be inside a loop",
				Type:   "Semantic",
			})
			log.Println("Continue statement must be inside a loop")

		}

	} else {
		v.Errors = append(v.Errors, Error{
			Line:   ctx.GetStart().GetLine(),
			Column: ctx.GetStart().GetColumn(),
			Msg:    "Continue statement must be inside a loop",
			Type:   "Semantic",
		})
		log.Println("Continue statement must be inside a loop")
	}

	return nil
}
