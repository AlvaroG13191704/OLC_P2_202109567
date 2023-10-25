package generator

import "fmt"

// Callfunction
func (g *Generator) CallFunctionIntFloat(value string) {
	// add comment
	g.GenComment("*** embbeded convert to Int or Float ***")
	// assign temps
	g.AssignTemp(
		g.GeneratorNativeVariables.EmbebbedNative.TempDigitStringConvertoToIntOrFloat,
		value,
	)

	if g.FunctionCode {
		// call the function
		g.FuncionUser = append(g.FuncionUser, "_convertStringToIntOrFloat();\n")
	} else {
		// call the function
		g.Code = append(g.Code, "_convertStringToIntOrFloat();\n")
	}

	g.AddNewLine()
}

// Callfunction
func (g *Generator) CallFunctionString(value string) string {
	tempReturn := g.NewTemp()
	// add comment
	g.GenComment("*** embbeded convert to String ***")
	// assign temps
	g.AssignTemp(
		g.GeneratorNativeVariables.EmbebbedNative.TempDigitToString,
		value,
	)
	if g.FunctionCode {
		// call the function
		g.FuncionUser = append(g.FuncionUser, "_convertToStringDigit();\n")
	} else {
		// call the function
		g.Code = append(g.Code, "_convertToStringDigit();\n")
	}
	g.AssignTemp(tempReturn, g.GeneratorNativeVariables.EmbebbedNative.TempReturnValue2)
	g.AddNewLine()

	return tempReturn
}

func (g *Generator) GenNativeIntFloat() string {
	var temBool bool
	// set false main code
	temBool = g.FunctionCode
	g.MainCode = false
	if g.FunctionCode {
		g.FunctionCode = false
	}

	// declare temps
	temp1 := g.NewTemp()
	g.GeneratorNativeVariables.EmbebbedNative.TempDigitStringConvertoToIntOrFloat = temp1
	temp2 := g.NewTemp()
	temp3 := g.NewTemp()
	temp4 := g.NewTemp()
	temp5 := g.NewTemp()
	temp6 := g.NewTemp()
	temp7 := g.NewTemp()
	// generate labels
	label0 := g.NewLabel()
	label1 := g.NewLabel()
	label2 := g.NewLabel()
	label3 := g.NewLabel()
	label4 := g.NewLabel()
	label5 := g.NewLabel()
	label6 := g.NewLabel()
	label7 := g.NewLabel()
	label8 := g.NewLabel()
	label9 := g.NewLabel()
	label10 := g.NewLabel()
	label11 := g.NewLabel()
	// gen code
	g.FuncCode = append(g.FuncCode, "void _convertStringToIntOrFloat() { \n")
	// assign temps
	g.GetHeap(temp2, temp1)
	g.AssignTemp(temp3, "0")
	g.AssignTemp(temp4, "0")
	g.AssignTemp(temp5, "0")
	g.AssignTemp(temp6, "1")
	g.AssignTemp(temp7, "0")
	// add label
	g.AddLabel(label0)
	// add if's
	g.AddIf(temp2, "-1", "==", label10)
	g.AddIf(temp2, "46", "==", label4)

	// add label 1
	g.AddLabel(label1)
	// add if's
	g.AddIf(temp2, "47", ">", label2)
	g.GoTo(label9) // error

	// add label 2
	g.AddLabel(label2)
	// add if's
	g.AddIf(temp2, "58", "<", label3)
	g.GoTo(label9) // error

	// add label 3
	g.AddLabel(label3)
	// sum
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp3, temp3, "*", "10"))
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp4, temp2, "-", "48"))
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp3, temp3, "+", temp4))
	// next character
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp1, temp1, "+", "1"))
	g.GetHeap(temp2, temp1)
	// goto L0
	g.GoTo(label0)

	// add label4
	g.AddLabel(label4)
	// jump dot
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp1, temp1, "+", "1"))
	g.GetHeap(temp2, temp1)

	// add label5
	g.AddLabel(label5)
	// add if
	g.AddIf(temp2, "-1", "==", label10)

	// add label 6
	g.AddLabel(label6)
	// add if's
	g.AddIf(temp2, "47", ">", label7)
	g.GoTo(label9) // error

	// add label7
	g.AddLabel(label7)
	// add if's
	g.AddIf(temp2, "58", "<", label8)
	g.GoTo(label9) // error

	// add label8
	g.AddLabel(label8)
	// sum
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp5, temp5, "*", "10"))
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp4, temp2, "-", "48"))
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp5, temp5, "+", temp4))
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp6, temp6, "*", "10"))
	// next character
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp1, temp1, "+", "1"))
	g.GetHeap(temp2, temp1)
	// goto L5
	g.GoTo(label5)

	// add label9
	g.AddLabel(label9)
	// add temp6
	g.AssignTemp(temp7, "9999999827968.00")
	g.GoTo(label11)

	// add label10
	g.AddLabel(label10)
	// add temp6
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp7, temp5, "/", temp6))
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp7, temp7, "+", temp3))

	// add label11
	g.AddLabel(label11)
	// return
	g.FuncCode = append(g.FuncCode, "return;\n")
	// close function
	g.FuncCode = append(g.FuncCode, "}\n")

	// return to main code
	g.MainCode = true
	if g.FunctionCode {
		g.FunctionCode = true
	}
	if temBool {
		g.FunctionCode = true
	}

	return temp7
}

// Function to convert into String

func (g *Generator) GenNativeDigitsToString() string {
	var temBool bool
	// set false main code
	temBool = g.FunctionCode
	g.MainCode = false
	if g.FunctionCode {
		g.FunctionCode = false
	}
	// declare temps
	temp0 := g.NewTemp()
	g.GeneratorNativeVariables.EmbebbedNative.TempDigitToString = temp0
	temp1 := g.NewTemp()
	temp2 := g.NewTemp()
	temp3 := g.NewTemp()
	temp4 := g.NewTemp()
	temp5 := g.NewTemp()
	temp6 := g.NewTemp()
	temp7 := g.NewTemp()
	temp8 := g.NewTemp()
	temp9 := g.NewTemp()
	temp10 := g.NewTemp()
	temp11 := g.NewTemp()
	temp12 := g.NewTemp()
	temp13 := g.NewTemp()
	// generate labels
	label0 := g.NewLabel()
	label1 := g.NewLabel()
	label2 := g.NewLabel()
	label3 := g.NewLabel()
	label4 := g.NewLabel()
	label5 := g.NewLabel()
	// gen code
	g.FuncCode = append(g.FuncCode, "void _convertToStringDigit() { \n")
	// assign temps
	g.AssignTemp(temp1, "(int)"+temp0)
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp2, temp0, "-", "(float)"+temp1))
	g.AssignTemp(temp3, "H")
	g.IncPointerHeap()
	g.AssignTemp(temp4, temp3)

	// add label
	g.AddLabel(label0)
	// add if
	g.AddIf(temp1, "0", "<=", label1)
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp5, "(int) "+temp1, "%", "10"))
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp6, temp5, "+", "48"))
	g.SaveHeapIndex(temp6, temp4)
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp1, "(int)"+temp1, "/", "10"))
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp4, temp4, "+", "1"))
	g.GoTo(label0)

	// add label1
	g.AddLabel(label1)
	g.AssignTemp(temp7, temp3)
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp8, temp4, "-", "1"))

	// add label2
	g.AddLabel(label2)
	// add if
	g.AddIf(temp7, temp8, ">=", label3)
	g.GetHeap(temp9, temp7)
	g.GetHeap(temp10, temp8)
	g.SaveHeapIndex(temp10, temp7)
	g.SaveHeapIndex(temp9, temp8) // pending
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp7, temp7, "+", "1"))
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp8, temp8, "-", "1"))
	g.GoTo(label2)

	// add label3
	g.AddLabel(label3)
	// add if
	g.AddIf(temp2, "0", "==", label5)
	g.SaveHeapIndex("46", temp4)
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp4, temp4, "+", "1"))
	g.AssignTemp(temp10, "1")
	g.AssignTemp(temp11, "6")

	// add label4
	g.AddLabel(label4)
	// add if
	g.AddIf(temp11, "0", "==", label5)
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp2, temp2, "*", "10"))
	g.AssignTemp(temp12, "(int)"+temp2)
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp13, temp12, "+", "48"))
	g.SaveHeapIndex(temp13, temp4)
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp2, temp2, "-", "(float)"+temp12))
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp11, temp11, "-", "1"))
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", temp4, temp4, "+", "1"))
	g.GoTo(label4)

	// add label5
	g.AddLabel(label5)
	g.SaveHeapIndex("-1", temp4)
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", "H", temp4, "+", "1"))

	// return
	g.FuncCode = append(g.FuncCode, "return;\n")
	// close function
	g.FuncCode = append(g.FuncCode, "}\n")

	g.MainCode = true
	if g.FunctionCode {
		g.FunctionCode = true
	}
	if temBool {
		g.FunctionCode = true
	}

	return temp3
}
