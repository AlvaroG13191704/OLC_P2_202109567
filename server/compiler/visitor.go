package compiler

import (
	"fmt"
	"log"
	"server/compiler/generator"
	"server/compiler/parser"

	"github.com/antlr4-go/antlr/v4"
)

func NewVisitor() *Visitor {
	return &Visitor{
		Generator:   generator.NewGenerator(),
		TableSymbol: []Symbol{},
	}

}

// The visit method
func (v *Visitor) Visit(tree antlr.ParseTree) interface{} {

	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		node := tree.Accept(v)
		return node
	}
}

// visit start
func (v *Visitor) VisitStart(ctx *parser.StartContext) interface{} {

	return v.Visit(ctx.Block())
}

// visit block
func (v *Visitor) VisitBlock(ctx *parser.BlockContext) interface{} {
	// push the scope
	v.pushScope()
	// declare all the functions
	if !v.FirstPass {
		fmt.Println("FIRST TRAVEL ")
		for _, stmt := range ctx.AllStmts() {
			if stmt.FunctionStmt() != nil {
				v.Visit(stmt.FunctionStmt())
			}
		}
		v.FirstPass = true

	}

	for _, stmt := range ctx.AllStmts() {
		v.Visit(stmt)
	}

	v.popScope()
	return nil
}

// visit stmts
func (v *Visitor) VisitStmts(ctx *parser.StmtsContext) interface{} {
	// verify if the stmt which stmt is
	if ctx.Declaration() != nil {
		return v.Visit(ctx.Declaration())
	}
	if ctx.Assignment() != nil {
		return v.Visit(ctx.Assignment())
	}
	if ctx.EmbbededFunc() != nil {
		return v.Visit(ctx.EmbbededFunc())
	}
	if ctx.Ifstmt() != nil {
		return v.Visit(ctx.Ifstmt())
	}
	if ctx.GuardStmt() != nil {
		return v.Visit(ctx.GuardStmt())
	}
	if ctx.SwitchStmt() != nil {
		return v.Visit(ctx.SwitchStmt())
	}
	if ctx.WhileStmt() != nil {
		return v.Visit(ctx.WhileStmt())
	}
	if ctx.ForStmt() != nil {
		return v.Visit(ctx.ForStmt())
	}
	// visit transfer
	if ctx.TransferStmt() != nil {
		return v.Visit(ctx.TransferStmt())
	}
	if ctx.CallFunctionStmt() != nil {
		return v.Visit(ctx.CallFunctionStmt())
	}
	// verify if the function is a print
	if ctx.Printstmt() != nil {
		return v.Visit(ctx.Printstmt())
	}
	if ctx.CallBack() != nil {
		return v.Visit(ctx.CallBack())
	}
	// structs
	if ctx.StructStmt() != nil {
		return v.Visit(ctx.StructStmt())
	}

	return nil
}

// visit embbededFunc
func (v *Visitor) VisitEmbbededFunc(ctx *parser.EmbbededFuncContext) interface{} {
	// verify if the function is a int
	if ctx.Intstmt() != nil {
		return v.Visit(ctx.Intstmt())
	}
	// verify if the function is a float
	if ctx.Floatstmt() != nil {
		return v.Visit(ctx.Floatstmt())
	}
	// verify if the function is a string
	if ctx.Stringstmt() != nil {
		return v.Visit(ctx.Stringstmt())
	}

	return nil

}
