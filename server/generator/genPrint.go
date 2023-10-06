package generator

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
	g.Code = append(g.Code, "_printString();\n")
	// add a %c 32
	g.GenPrint("c", "32")
	g.GenPrint("c", "10")
	// add a new line
	g.AddNewLine()
}
