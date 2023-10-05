package generator

import "fmt"

// Global Function to print strings
func (g *Generator) GenerateFuncToPrint() {

	// set false main code
	g.MainCode = false

	g.TempToPrint = g.NewTemp()
	g.TempToEvaluateEndString = g.NewTemp()
	// generate label
	labelStart := g.NewLabel()
	labelEnd := g.NewLabel()
	// add the code
	g.FuncCode = append(g.FuncCode, "void _printString() { \n")
	// add label
	g.AddLabel(labelStart)
	// get the position of the string using the temp
	g.GetHeap(g.TempToEvaluateEndString, g.TempToPrint)
	// evaluate if the string is not the end
	g.AddIf(g.TempToEvaluateEndString, "-1", "==", labelEnd)
	// print the string
	g.GenPrint("c", "(int)"+g.TempToEvaluateEndString)
	// increment the temp
	g.GenArithmetic(g.TempToPrint, g.TempToPrint, "1", "+")
	// go to the start
	g.GoTo(labelStart)
	// add label
	g.AddLabel(labelEnd)
	// end of the function
	g.FuncCode = append(g.FuncCode, "return; \n")
	g.FuncCode = append(g.FuncCode, "} \n")
	// set false main code
	g.MainCode = true

}

// Function to generate c3d when a variable is declared
func (g *Generator) GenString(value string) string {

	// add comment
	g.GenComment("*** working with strings ***")
	// generate a new temp and assing the heap pointer
	tempString := g.NewTemp()
	g.AssignTemp(tempString, "H")
	// iterate the string and save the chars in the heap
	for _, char := range value {
		// convert to ascii
		charString := fmt.Sprintf("%d", char)
		// save in the heap
		g.SaveHeap(charString)
		// increment pointer heap
		g.IncPointerHeap()
	}
	// save the end of the string
	g.SaveHeap("-1")
	// increment pointer heap
	g.IncPointerHeap()

	return tempString
}
