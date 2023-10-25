package generator

import "fmt"

func (g *Generator) GenArithmetic(target, left, right, operator string) {

	// evaluate if is / or % to add dviision by zero
	if operator == "/" || operator == "%" {
		if !g.GeneratorNativeVariables.ZeroNative.IsPosibleDivisionZero {
			// gen zero
			g.GenDivisionZero()
			g.GeneratorNativeVariables.ZeroNative.IsPosibleDivisionZero = true
		}

		if g.FunctionCode {
			// generate label
			labelEnd := g.NewLabel()
			// assign the temp to the tempDivisionZero
			g.AssignTemp(g.GeneratorNativeVariables.ZeroNative.TempDivisionZero, right)
			// call the function
			g.FuncionUser = append(g.FuncionUser, "_zero_division();\n")
			// evaluate if the tempDivisionZero is 1
			g.AddIf(g.GeneratorNativeVariables.ZeroNative.TempDivisionZero, "1", "==", labelEnd)
			// add the code
			g.FuncionUser = append(g.FuncionUser, fmt.Sprintf("%s = %s %s %s;\n", target, left, operator, right))
			// add label
			g.AddLabel(labelEnd)
		} else {

			// generate label
			labelEnd := g.NewLabel()
			// assign the temp to the tempDivisionZero
			g.AssignTemp(g.GeneratorNativeVariables.ZeroNative.TempDivisionZero, right)
			// call the function
			g.Code = append(g.Code, "_zero_division();\n")
			// evaluate if the tempDivisionZero is 1
			g.AddIf(g.GeneratorNativeVariables.ZeroNative.TempDivisionZero, "1", "==", labelEnd)
			// add the code
			g.Code = append(g.Code, fmt.Sprintf("%s = %s %s %s;\n", target, left, operator, right))
			// add label
			g.AddLabel(labelEnd)
		}

	} else {
		if g.FunctionCode {
			g.FuncionUser = append(g.FuncionUser, fmt.Sprintf("%s = %s %s %s;\n", target, left, operator, right))
		} else if g.MainCode {
			g.Code = append(g.Code, fmt.Sprintf("%s = %s %s %s;\n", target, left, operator, right))
		} else {
			g.FuncCode = append(g.FuncCode, fmt.Sprintf("%s = %s %s %s;\n", target, left, operator, right))
		}
	}

}

func (g *Generator) GenDivisionZero() {

	// set false main
	var temBool bool
	// set false main code
	temBool = g.FunctionCode
	g.MainCode = false
	if g.FunctionCode {
		g.FunctionCode = false
	}

	// g.TempDivisionZero = tempZero
	g.GeneratorNativeVariables.ZeroNative.TempDivisionZero = g.NewTemp()

	// generate labels
	labelIfTrue := g.NewLabel()
	labelEnd := g.NewLabel()

	// add the code
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("void _zero_division() {\n"))
	// add if
	g.AddIf("(int)"+g.GeneratorNativeVariables.ZeroNative.TempDivisionZero, "0", "!=", labelIfTrue)
	// iterate over the string to conver to ascii code
	for _, char := range "Error: Division by zero" {
		// generate c3d
		g.GenPrint("c", fmt.Sprintf("%v", int(char)))
	}
	g.GenPrint("c", "32")
	g.GenPrint("c", "10")

	// assign 1 to the temp
	g.AssignTemp(g.GeneratorNativeVariables.ZeroNative.TempDivisionZero, "1")
	// go to the end
	g.GoTo(labelEnd)
	// add label
	g.AddLabel(labelIfTrue)
	// assign 0 to the temp
	g.AssignTemp(g.GeneratorNativeVariables.ZeroNative.TempDivisionZero, "0")
	// add label
	g.AddLabel(labelEnd)
	// end of the function
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("return; \n"))
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("} \n"))

	g.MainCode = true
	if g.FunctionCode {
		g.FunctionCode = true
	}
	if temBool {
		g.FunctionCode = true
	}
}
