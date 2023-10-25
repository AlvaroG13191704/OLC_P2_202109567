package generator

func (g *Generator) GenPrint(typePrint string, value string) {
	if g.FunctionCode {
		g.FuncionUser = append(g.FuncionUser, "printf(\"%"+typePrint+"\", "+value+");\n")
	} else if g.MainCode {
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
	g.AssignTemp(g.GeneratorNativeVariables.PrintNative.TempToPrint, tempIndex)

	if g.FunctionCode {
		// call the function
		g.FuncionUser = append(g.FuncionUser, "_printString();\n")
	} else {
		// call the function
		g.Code = append(g.Code, "_printString();\n")
	}
}
