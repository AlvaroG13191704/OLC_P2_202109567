package generator

import (
	"fmt"
	"strings"
)

type Generator struct {
	Temp                    int           // manage the temp variables
	StackCounter            int           // manage the stack
	HeapCounter             int           // manage the heap
	Label                   int           // manage the labels
	TempToPrint             string        // temp to print
	TempToEvaluateEndString string        // temp to evaluate the end of the string
	Code                    []interface{} // the final code
	TempList                []interface{} // list of temp variables
	Compiled                []interface{} // list of compiled code
	FinalCode               []interface{} // list of final code
	FuncCode                []interface{}
	BreakLabel              string // break label
	ContinueLabel           string // continue label
	MainCode                bool
}

func NewGenerator() *Generator {
	return &Generator{
		Temp:          0,
		Label:         0,
		BreakLabel:    "",
		ContinueLabel: "",
		MainCode:      true,
	}
}

func (g Generator) FillCode() string {
	initValues := "#include <stdio.h>\n"
	initValues += "float stack[1000000]; // STACK\n"
	initValues += "float heap[1000000];  // HEAP\n"
	initValues += "float P;              // Stack pointer \n"
	initValues += "float H;              // Heap pointer \n"

	return initValues
}

func (g Generator) FillWithTemp() string {

	if len(g.TempList) == 0 {
		return "\n"
	}

	tempString := "float "
	for _, v := range g.TempList {
		tempString += v.(string) + ", "
	}
	// delete the last comma
	tempString = tempString[:len(tempString)-2]
	tempString += ";\n"
	return tempString
}

func (g Generator) GetCode() []interface{} {
	// add to the bennining of the code the temp variables stack, heap, pointer stack, pointer heap

	return g.Code
}

func (g Generator) GetFinalCode() string {
	var resultString string
	// add to the bennining of the code the temp variables stack, heap, pointer stack, pointer heap
	resultString += g.FillCode()
	// add the list of temporals
	resultString += g.FillWithTemp()

	// add the function to print strings
	// resultString += g.GenerateFuncToPrint()

	// add main
	resultString += "int main() {\n"

	// get the code generated
	codeGenerated := g.GetCode()
	// return string
	for _, v := range codeGenerated {
		resultString += v.(string)
	}

	// add the end of the main
	resultString += "return 0;\n"
	resultString += "}"

	return resultString
}

// increment pointer stack
func (g *Generator) IncPointerStack() {
	g.Code = append(g.Code, "P = P + 1;\n\n")
	g.StackCounter = g.StackCounter + 1
}

// decrement pointer stack
func (g *Generator) DecPointerStack() {
	g.Code = append(g.Code, "P = P - 1;\n")
	g.StackCounter = g.StackCounter + 1
}

// increment pointer heap
func (g *Generator) IncPointerHeap() {
	g.Code = append(g.Code, "H = H + 1;\n")
	g.HeapCounter = g.HeapCounter + 1
}

// // add break lvl
// func (g *Generator) AddBreak(lvl string) {
// 	g.BreakLabel = lvl
// }

// // add continue lvl
// func (g *Generator) AddContinue(lvl string) {
// 	g.ContinueLabel = lvl
// }

// add a \n
func (g *Generator) AddNewLine() {
	if g.MainCode {
		g.Code = append(g.Code, "\n")
	} else {
		g.FuncCode = append(g.FuncCode, "\n")
	}
}

// Generate a new Temp
func (g *Generator) NewTemp() string {
	temp := "t" + fmt.Sprintf("%v", g.Temp)
	g.Temp = g.Temp + 1
	//Lo guardamos para declararlo
	g.TempList = append(g.TempList, temp)
	return temp
}

// Generate a new Label
func (g *Generator) NewLabel() string {
	label := "L" + fmt.Sprintf("%v", g.Label)
	g.Label = g.Label + 1
	return label
}

// generate comment, receive a list of string args
func (g *Generator) GenComment(args ...string) {
	comment := fmt.Sprintf("// %s \n", strings.Join(args, " "))
	g.Code = append(g.Code, comment)
}

// save stack
func (g *Generator) SaveStack(varName, tempString string) {
	saveStack := fmt.Sprintf("stack[(int)P] = %s; \n", tempString)
	g.Code = append(g.Code, saveStack)

}

// Save heap
func (g *Generator) SaveHeap(tempString string) {
	saveHeap := fmt.Sprintf("heap[(int)H] = %s; \n", tempString)
	g.Code = append(g.Code, saveHeap)
}

// Assign temp
func (g *Generator) AssignTemp(tempString string, value string) string {
	assingTemp := fmt.Sprintf("%s = %s;\n", tempString, value)
	g.Code = append(g.Code, assingTemp)
	return assingTemp
}

// Access stack
func (g *Generator) AccessStack(tempString string, index string) string {
	accessStack := fmt.Sprintf("%s = stack[(int)%s];\n", tempString, index)
	g.Code = append(g.Code, accessStack)
	return accessStack
}
