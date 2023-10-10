package generator

import (
	"fmt"
	"strings"
)

type Generator struct {
	Temp                     int                      // manage the temp variables
	StackCounter             int                      // manage the stack
	HeapCounter              int                      // manage the heap
	Label                    int                      // manage the labels
	GeneratorNativeVariables GeneratorNativeVariables // start native temps
	Code                     []interface{}            // the final code
	TempList                 []interface{}            // list of temp variables
	Compiled                 []interface{}            // list of compiled code
	FinalCode                []interface{}            // list of final code
	FuncCode                 []interface{}
	BreakLabel               string // break label
	ContinueLabel            string // continue label
	MainCode                 bool
}

func NewGenerator() *Generator {
	return &Generator{
		Temp:          0,
		Label:         0,
		BreakLabel:    "",
		ContinueLabel: "",
		MainCode:      true,
		StackCounter:  0,
		HeapCounter:   0,
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
	return g.Code
}

func (g Generator) GetFuncCode() []interface{} {
	return g.FuncCode
}

func (g Generator) GetFinalCode() string {

	var resultString string
	// add to the bennining of the code the temp variables stack, heap, pointer stack, pointer heap
	resultString += g.FillCode()

	// add the list of temporals
	resultString += g.FillWithTemp()

	// get the code generated
	codeGenerated := g.GetFuncCode()
	// return string
	for _, v := range codeGenerated {
		resultString += v.(string)
	}

	// add main
	resultString += "int main() {\n"

	// get the code generated
	codeGenerated = g.GetCode()
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
	if g.MainCode {
		g.Code = append(g.Code, "P = P + 1;\n")
	} else {
		g.FuncCode = append(g.FuncCode, "P = P + 1;\n")
	}

	// g.StackCounter = g.StackCounter + 1

}

// decrement pointer stack
func (g *Generator) DecPointerStack() {
	if g.MainCode {
		g.Code = append(g.Code, "P = P - 1;\n")
	} else {
		g.FuncCode = append(g.FuncCode, "P = P - 1;\n")
	}
	// g.StackCounter = g.StackCounter - 1
}

// increment pointer heap
func (g *Generator) IncPointerHeap() {
	if g.MainCode {
		g.Code = append(g.Code, "H = H + 1;\n")
	} else {
		g.FuncCode = append(g.FuncCode, "H = H + 1;\n")
	}
	g.HeapCounter = g.HeapCounter + 1
}

// increment counter
func (g *Generator) CounterStack(which string) {
	if which == "+" {
		g.StackCounter = g.StackCounter + 1
	} else {
		g.StackCounter = g.StackCounter - 1
	}

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

// add a goto
func (g *Generator) GoTo(label string) {
	if g.MainCode {
		g.Code = append(g.Code, fmt.Sprintf("goto %s;\n", label))
	} else {
		g.FuncCode = append(g.FuncCode, fmt.Sprintf("goto %s;\n", label))
	}
}

// add a label
func (g *Generator) AddLabel(label string) {
	if g.MainCode {
		g.Code = append(g.Code, fmt.Sprintf("%s: \n", label))
	} else {
		g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s: \n", label))
	}
}

// generate comment, receive a list of string args
func (g *Generator) GenComment(args ...string) {
	if g.MainCode {
		g.Code = append(g.Code, "// "+strings.Join(args, " ")+"\n")
	} else {
		g.FuncCode = append(g.FuncCode, "// "+strings.Join(args, " ")+"\n")
	}
}

// save stack
func (g *Generator) SaveStack(varName, tempString string) {
	if g.MainCode {
		saveStack := fmt.Sprintf("stack[(int)%s] = %s; \n", varName, tempString)
		g.Code = append(g.Code, saveStack)
	} else {
		saveStack := fmt.Sprintf("stack[(int)%s] = %s; \n", varName, tempString)
		g.FuncCode = append(g.FuncCode, saveStack)
	}
}

// Assing stack
func (g *Generator) AssignStack(valueP int, tempString string) {
	if g.MainCode {
		assingStack := fmt.Sprintf("stack[(int)P + %v] = %s; \n", valueP, tempString)
		g.Code = append(g.Code, assingStack)
	} else {
		assingStack := fmt.Sprintf("stack[(int)P + %v] = %s; \n", valueP, tempString)
		g.FuncCode = append(g.FuncCode, assingStack)
	}
}

// get stack
func (g *Generator) GetStack(tempString, valueP string) {
	if g.MainCode {
		getStack := fmt.Sprintf("%s = stack[(int) %s]; \n", tempString, valueP)
		g.Code = append(g.Code, getStack)
	} else {
		getStack := fmt.Sprintf("%s = stack[(int) %s]; \n", tempString, valueP)
		g.FuncCode = append(g.FuncCode, getStack)
	}
}

// get heap
func (g *Generator) GetHeap(tempString, valueH string) {
	if g.MainCode {
		getHeap := fmt.Sprintf("%s = heap[(int) %s]; \n", tempString, valueH)
		g.Code = append(g.Code, getHeap)
	} else {
		getHeap := fmt.Sprintf("%s = heap[(int) %s]; \n", tempString, valueH)
		g.FuncCode = append(g.FuncCode, getHeap)
	}
}

// Save heap
func (g *Generator) SaveHeap(tempString string) {
	if g.MainCode {
		saveHeap := fmt.Sprintf("heap[(int)H] = %s; \n", tempString)
		g.Code = append(g.Code, saveHeap)
	} else {
		saveHeap := fmt.Sprintf("heap[(int)H] = %s; \n", tempString)
		g.FuncCode = append(g.FuncCode, saveHeap)
	}
}

// Assign temp
func (g *Generator) AssignTemp(tempString string, value string) string {

	assingTemp := fmt.Sprintf("%s = %s; \n", tempString, value)
	if g.MainCode {
		g.Code = append(g.Code, assingTemp)
	} else {
		g.FuncCode = append(g.FuncCode, assingTemp)
	}
	return assingTemp
}

// Access stack
func (g *Generator) AccessStack(tempString string, index string) string {
	accessStack := fmt.Sprintf("%s = stack[(int)%s];\n", tempString, index)
	if g.MainCode {
		g.Code = append(g.Code, accessStack)
	} else {
		g.FuncCode = append(g.FuncCode, accessStack)
	}
	return accessStack
}

// INSTRUCTIONS
// add If statement
func (g *Generator) AddIf(left, right, operator, label string) {
	ifStatement := fmt.Sprintf("if (%s %s %s) goto %s;\n", left, operator, right, label)
	if g.MainCode {
		g.Code = append(g.Code, ifStatement)
	} else {

		g.FuncCode = append(g.FuncCode, ifStatement)
	}
}
