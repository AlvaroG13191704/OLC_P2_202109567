package generator

import "fmt"

// Global Function to print strings
func (g *Generator) GenerateFuncToPrint() string {
	// create temp
	tempToIterateString := g.NewTemp()
	g.TempToPrint = tempToIterateString
	tempToEvaluateString := g.NewTemp()
	g.TempToEvaluateEndString = tempToEvaluateString
	// generate label
	labelStart := g.NewLabel()
	labelEnd := g.NewLabel()
	// add the code
	functionString := fmt.Sprintf("void printString() {\n")
	// add label
	functionString += fmt.Sprintf("%s: \n", labelStart)
	// get the position of the string using the temp
	functionString += fmt.Sprintf("%s = heap[(int)%s]; \n", tempToIterateString, tempToEvaluateString)
	// evaluate if the string is not the end
	functionString += fmt.Sprintf("if (%s != -1) goto %s; \n", tempToIterateString, labelEnd)
	// print the string
	g.GenPrint("c", tempToIterateString)
	// increment the temp
	functionString += fmt.Sprintf("%s = %s + 1; \n", tempToEvaluateString, tempToEvaluateString)
	// go to the start
	functionString += fmt.Sprintf("goto %s; \n", labelStart)
	// add label
	functionString += fmt.Sprintf("%s: \n", labelEnd)
	// end of the function
	functionString += fmt.Sprintf("return; \n")
	functionString += fmt.Sprintf("} \n")

	return functionString
}

func (g *Generator) GenPrint(typePrint string, value string) {
	if g.MainCode {
		g.Code = append(g.Code, "printf(\"%"+typePrint+"\", "+value+");\n")
	} else {
		g.FuncCode = append(g.FuncCode, "printf(\"%"+typePrint+"\", "+value+");\n")
	}
}

// GenPrintString
func (g *Generator) GenPrintString(tempIndex string) {
	// add comment
	g.GenComment("Print a string")
	// assing the temp to the tempString
	g.AssignTemp(g.TempToPrint, tempIndex)
	// call the function
	g.FuncCode = append(g.FuncCode, "printString();\n")
	// add a %c 32
	g.GenPrint("c", "32")
	// add a new line
	g.AddNewLine()
}
