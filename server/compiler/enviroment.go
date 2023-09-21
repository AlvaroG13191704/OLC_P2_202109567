package compiler

import (
	"server/compiler/generator"
	"server/compiler/parser"
)

type Symbol struct {
	Id          string // the name of the variable
	TypeSymbol  string // the type of the symbol variable or function, vector, reference, or matrix
	TypeMutable string // the type of the variable -> var or let
	TypeData    string // the type of the data -> Int, Float, String, Boolean, Character
	StructOf    string // the name of the struct that contains the variable
	Value       interface{}
	ListParams  interface{}
	Mutating    bool
	Line        int
	Column      int
}

type Error struct {
	Line   int
	Column int
	Msg    string
	Type   string
}

type Visitor struct {
	parser.BaseGrammarVisitor
	Generator   *generator.Generator
	SymbolStack []map[string]Symbol
	Errors      []Error
}

// Manage the scopes
func (v *Visitor) pushScope() {
	v.SymbolStack = append(v.SymbolStack, make(map[string]Symbol))
}

func (v *Visitor) popScope() {
	if len(v.SymbolStack) > 1 {
		v.SymbolStack = v.SymbolStack[:len(v.SymbolStack)-1]
	}
}

func (v *Visitor) getCurrentScope() map[string]Symbol {
	return v.SymbolStack[len(v.SymbolStack)-1]
}
