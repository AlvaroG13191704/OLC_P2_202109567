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

// Function to concat strings
func (g *Generator) ConcantStrings() string {
	// set false main code
	g.MainCode = false
	// generate temps
	initHeapPointer := g.NewTemp()
	g.TempInitConcat = initHeapPointer
	tempFirstString := g.NewTemp()
	g.TempFirstConcat = tempFirstString
	tempSecondString := g.NewTemp()
	g.TempSecondConcat = tempSecondString
	tempIterator := g.NewTemp()
	// labels
	labelFirstLoop := g.NewLabel()
	labelSecondLoop := g.NewLabel()
	labelEnd := g.NewLabel()
	// add the code
	g.FuncCode = append(g.FuncCode, "void _concatString() { \n")
	// assign the heap pointer
	g.AssignTemp(initHeapPointer, "H")
	// add label
	g.AddLabel(labelFirstLoop)
	// get the heap pointer from the first string
	g.GetHeap(tempIterator, tempFirstString)
	// add the if to evaluate if the string is not the end
	g.AddIf(tempIterator, "-1", "==", labelSecondLoop)
	// save the char in the heap
	g.SaveHeap(tempIterator)
	// increment the heap pointer
	g.IncPointerHeap()
	// increment the temp
	g.GenArithmetic(tempFirstString, tempFirstString, "1", "+")
	// add goto
	g.GoTo(labelFirstLoop)
	// add label
	g.AddLabel(labelSecondLoop)
	// assign the temp iterator the second string
	g.GetHeap(tempIterator, tempSecondString)
	// add the if to evaluate if the string is not the end
	g.AddIf(tempIterator, "-1", "==", labelEnd)
	// save the char in the heap
	g.SaveHeap(tempIterator)
	// increment the heap pointer
	g.IncPointerHeap()
	// increment the temp
	g.GenArithmetic(tempSecondString, tempSecondString, "1", "+")
	// add goto
	g.GoTo(labelSecondLoop)
	// add label
	g.AddLabel(labelEnd)
	// save the end of the string
	g.SaveHeap("-1")
	// increment the heap pointer
	g.IncPointerHeap()
	// end of the function
	g.FuncCode = append(g.FuncCode, "return; \n")
	g.FuncCode = append(g.FuncCode, "} \n")

	// return to main code
	g.MainCode = true

	return initHeapPointer
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

// Function to concat strings
func (g *Generator) GenConcatString(left, right string) {
	// add comment
	g.GenComment("*** concat strings ***")
	// assign temps
	g.AssignTemp(g.TempFirstConcat, left)
	g.AssignTemp(g.TempSecondConcat, right)
	// call the function
	g.Code = append(g.Code, "_concatString();\n")
	g.AddNewLine()
}

// Func to compare strings
func (g *Generator) GenComparationString(left, right string) {
	// add comment
	g.GenComment("*** compare strings ***")
	// assign temps
	g.AssignTemp(g.TempFirstStringCompare, left)
	g.AssignTemp(g.TempSecondStringCompare, right)
	// call the function
	g.Code = append(g.Code, "_compare_strings();\n")
	g.AddNewLine()
}
