package compiler

import (
	"server/compiler/compiler/values"
	"server/compiler/generator"

	"server/compiler/parser"
)

type Symbol struct {
	// general
	TypeSymbol  string // the type of the symbol variable or function, vector, reference, or matrix
	TypeMutable string // the type of the variable -> var or let
	TypeData    string // the type of the data -> Int, Float, String, Boolean, Character
	Line        int
	Column      int
	// Function attributes
	sizeFunction int
	ListParams   []Symbol
	ReturnType   string
	ReturnTemp   string
	ReturnLabel  string

	// variables attributes
	Id             string // the name of the variable
	Value          interface{}
	LenghtVector   int
	CodeSentence   string
	StackDirection int
	HeapDirection  int
}

type LoopContext struct {
	TypeLoop      string
	ContinueFound bool
	BreakFound    bool
	StartLabel    string
	EndLabel      string
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
	// manage loop context
	loopContexts []LoopContext
	Errors       []Error
	// function
	FunctionContext []string
	FirstPass       bool
	SizeFunction    int
	ReturnLabel     string
	ReturnTemp      string
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

// Verify if the variale is in the scope only return true or false
func (v *Visitor) VerifyVariable(varName string) bool {

	for i := len(v.SymbolStack) - 1; i >= 0; i-- {
		scope := v.SymbolStack[i]
		if _, ok := scope[varName]; ok {
			return true
		}
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

// udpate a variable in the scope
func (v *Visitor) UpdateVariable(varName string, value interface{}) {
	for i := len(v.SymbolStack) - 1; i >= 0; i-- {
		scope := v.SymbolStack[i]
		if _, ok := scope[varName]; ok {
			scope[varName] = value.(Symbol)
			return
		}
	}
}

// AssignVariable assign a variable to the current scope -> works for loops
func (v *Visitor) AssignVariable(varName string, value interface{}) {
	scope := v.getCurrentScope()
	scope[varName] = value.(Symbol)
}

// DeleteVariable delete a variable from the current scope -> works for loops
func (v *Visitor) DeleteVariable(varName string) {
	scope := v.getCurrentScope()
	delete(scope, varName)
}

// PushLoopContext push a loop context
func (v *Visitor) PushLoopContext(typeLoop string, startLabel string, endLabel string) {
	v.loopContexts = append(v.loopContexts, LoopContext{
		TypeLoop:      typeLoop,
		ContinueFound: false,
		BreakFound:    false,
		StartLabel:    startLabel,
		EndLabel:      endLabel,
	})
}

// PopLoopContext pop a loop context
func (v *Visitor) PopLoopContext() {
	if len(v.loopContexts) > 0 {
		v.loopContexts = v.loopContexts[:len(v.loopContexts)-1]
	}
}

// ExistsLoopContext check if a loop context exists
func (v *Visitor) ExistsLoopContext() bool {
	return len(v.loopContexts) > 0
}

// update the current loop context
func (v *Visitor) UpdateLoopContext(ctx LoopContext) {
	v.loopContexts[len(v.loopContexts)-1] = ctx
}

// GetLoopContext get the current loop context
func (v *Visitor) GetLoopContext() LoopContext {

	return v.loopContexts[len(v.loopContexts)-1]
}

// VerifyFunctionScope verify if the function is already declared in the current scope
func (v *Visitor) VerifyFunctionScope(varName string) bool {

	scope := v.SymbolStack[len(v.SymbolStack)-1]
	if variable, ok := scope[varName]; ok {
		// evaluate if the variable is a function
		return variable.TypeSymbol == values.Function_Type

	}
	return false
}

// GetFunction from the scope
func (v *Visitor) GetFunction(varName string) (Symbol, bool) {
	for i := len(v.SymbolStack) - 1; i >= 0; i-- {
		scope := v.SymbolStack[i]
		if val, ok := scope[varName]; ok {
			if val.TypeSymbol == values.Function_Type {
				return val, true
			}
		}
	}
	return Symbol{}, false
}

// ExistsFunctionContext check if a function context exists
func (v *Visitor) ExistsFunctionContext() bool {
	return len(v.FunctionContext) > 0
}

// PushFunctionContext push a function context
func (v *Visitor) PushFunctionContext(functionContext string) {
	v.FunctionContext = append(v.FunctionContext, functionContext)
}

// PopFunctionContext pop a function context
func (v *Visitor) PopFunctionContext() {
	if len(v.FunctionContext) > 0 {
		v.FunctionContext = v.FunctionContext[:len(v.FunctionContext)-1]
	}
}
