package compiler

import (
	"server/compiler/compiler/values"
	"server/compiler/generator"
	"server/compiler/parser"
)

type Symbol struct {
	TypeSymbol  string // the type of the symbol variable or function, vector, reference, or matrix
	TypeMutable string // the type of the variable -> var or let
	TypeData    string // the type of the data -> Int, Float, String, Boolean, Character
	StructOf    string // the name of the struct that contains the variable
	ListParams  interface{}
	Mutating    bool
	Line        int
	Column      int
	// the following attributes are for the generator
	Id             string // the name of the variable
	TempString     string
	Value          interface{}
	CodeSentence   string
	StackDirection int
	HeapDirection  int
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

// VerifyVariableCurrentScope verify if the variable is already declared in the current scope
func (v *Visitor) VerifyVariableCurrentScope(varName string) bool {

	scope := v.SymbolStack[len(v.SymbolStack)-1]
	// fmt.Println("Current scope ->", scope)
	if variable, ok := scope[varName]; ok {
		// evaluate if the variable is not a function
		return variable.TypeSymbol == values.Variable_Type

	}
	return false
}

// VerifyScope verify if the variable is in the scope
func (v *Visitor) VerifyScope(varName string) (interface{}, bool) {

	for i := len(v.SymbolStack) - 1; i >= 0; i-- {
		scope := v.SymbolStack[i]
		if val, ok := scope[varName]; ok {
			return val, true
		}
	}
	return nil, false
}
