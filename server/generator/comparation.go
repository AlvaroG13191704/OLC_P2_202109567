package generator

import "fmt"

func (g *Generator) GenPrintBoolFunc() {
	// set false main
	var temBool bool
	// set false main code
	temBool = g.FunctionCode
	g.MainCode = false
	if g.FunctionCode {
		g.FunctionCode = false
	}

	g.GeneratorNativeVariables.PrintBoolNative.TempBoolFunc = g.NewTemp()
	// generate labels
	labelIfTrue := g.NewLabel()
	labelEnd := g.NewLabel()
	// code
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("void _print_bool() {\n"))
	// add if
	g.AddIf(g.GeneratorNativeVariables.PrintBoolNative.TempBoolFunc, "0", "==", labelIfTrue)
	// generate "true"
	for _, char := range "true" {
		g.GenPrint("c", fmt.Sprintf("%v", int(char)))
	}
	// add goto
	g.GoTo(labelEnd)
	// add label
	g.AddLabel(labelIfTrue)
	// generate "false"
	for _, char := range "false" {
		g.GenPrint("c", fmt.Sprintf("%v", int(char)))
	}
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

// t1 = 1;
// if (left == right ) goto L1;
// t1 = 0;
// L1:
// t2 = t1;
// __print_bool();
// Func to generate the comparation c3d
func (g *Generator) GenComparation(left, right, sign string) string {
	// generate labels
	labelEnd := g.NewLabel()
	// generate temp
	temp := g.NewTemp()
	// assign true to the temp
	g.AssignTemp(temp, "1")
	// add if
	g.AddIf(left, right, sign, labelEnd)
	// assign false to the temp
	g.AssignTemp(temp, "0")
	// add label
	g.AddLabel(labelEnd)

	return temp

}

// Func to compare strings
func (g *Generator) GenComparationStringFunc() string {
	// set false main
	var temBool bool
	// set false main code
	temBool = g.FunctionCode
	g.MainCode = false
	if g.FunctionCode {
		g.FunctionCode = false
	}

	// gen comment
	g.GenComment("Comparation String")
	// generate labels
	labelL1 := g.NewLabel()
	labelL2 := g.NewLabel()
	label3 := g.NewLabel()
	labelL4 := g.NewLabel()
	// generate temps
	temp1 := g.NewTemp()
	g.GeneratorNativeVariables.CompareStringsNative.TempFirstStringCompare = temp1
	temp2 := g.NewTemp()
	g.GeneratorNativeVariables.CompareStringsNative.TempSecondStringCompare = temp2
	temp3 := g.NewTemp()
	temp4 := g.NewTemp()
	temp5 := g.NewTemp()
	// code
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("void _compare_strings() {\n"))
	// assign true to the temp
	g.AssignTemp(temp3, "1")
	// add label
	g.AddLabel(labelL1)
	// t4 = heap[(int) t1];
	g.GetHeap(temp4, temp1)
	// t5 = heap[(int) t2];
	g.GetHeap(temp5, temp2)
	// if (t4 == -1) goto L2;
	g.AddIf(temp4, "-1", "==", labelL2)
	// if (t5 == -1) goto L4;
	g.AddIf(temp5, "-1", "==", label3)
	// if (t4 != t5) goto L3;
	g.AddIf(temp4, temp5, "!=", label3)
	// t1 = t1 + 1;
	g.GenArithmetic(temp1, temp1, "1", "+")
	// t2 = t2 + 1;
	g.GenArithmetic(temp2, temp2, "1", "+")
	// goto L1;
	g.GoTo(labelL1)
	// add label
	g.AddLabel(labelL2)
	// if(t5 == -1) goto L4;
	g.AddIf(temp5, "-1", "==", labelL4)
	// t3 = 0;
	g.AssignTemp(temp3, "0")
	// add label
	g.AddLabel(label3)
	// t3 = 0;
	g.AssignTemp(temp3, "0")
	// add label
	g.AddLabel(labelL4)

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
	return temp3
}

// Func to compare strings
func (g *Generator) GenComparationStringFuncGLE() string {

	// set false main
	var temBool bool
	// set false main code
	temBool = g.FunctionCode
	g.MainCode = false
	if g.FunctionCode {
		g.FunctionCode = false
	}

	// create temps
	temp0 := g.NewTemp()
	g.GeneratorNativeVariables.CompareStringsNative.TempFirstStringCompareGLE = temp0
	temp1 := g.NewTemp()
	g.GeneratorNativeVariables.CompareStringsNative.TempSecondStringCompareGLE = temp1
	temp2 := g.NewTemp()
	temp3 := g.NewTemp()
	temp4 := g.NewTemp()
	temp5 := g.NewTemp()
	temp6 := g.NewTemp()

	// create labels
	labelL0 := g.NewLabel()
	labelL1 := g.NewLabel()
	labelL2 := g.NewLabel()
	labelL3 := g.NewLabel()
	labelL4 := g.NewLabel()

	// code
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("void _compare_stringsGLE() {\n"))

	// t2 = t0;
	g.AssignTemp(temp2, temp0)
	// t3 = t1;
	g.AssignTemp(temp3, temp1)
	// t4 = heap[(int) t2];
	g.GetHeap(temp4, temp2)
	// t5 = heap[(int) t3];
	g.GetHeap(temp5, temp3)
	// add label
	g.AddLabel(labelL0)
	// if(t4 > t5) goto L2;
	g.AddIf(temp4, temp5, ">", labelL2)
	// if(t4 == t5) goto L1;
	g.AddIf(temp4, temp5, "==", labelL1)
	// goto L3;
	g.GoTo(labelL3)
	// add label
	g.AddLabel(labelL1)
	// if(t4 == -1) goto L2;
	g.AddIf(temp4, "-1", "==", labelL2)
	// t2 = t2 + 1;
	g.GenArithmetic(temp2, temp2, "1", "+")
	// t3 = t3 + 1;
	g.GenArithmetic(temp3, temp3, "1", "+")
	// t4 = heap[(int) t2];
	g.GetHeap(temp4, temp2)
	// t5 = heap[(int) t3];
	g.GetHeap(temp5, temp3)
	// goto L0;
	g.GoTo(labelL0)
	// add label
	g.AddLabel(labelL3)
	// t6 = 0;
	g.AssignTemp(temp6, "0")
	// goto L4;
	g.GoTo(labelL4)
	// add label
	g.AddLabel(labelL2)
	// t6 = 1;
	g.AssignTemp(temp6, "1")
	// add label
	g.AddLabel(labelL4)

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
	return temp6
}

// Func to compare strings
func (g *Generator) GenComparationStringFuncGL() string {
	// set false main
	var temBool bool
	// set false main code
	temBool = g.FunctionCode
	g.MainCode = false
	if g.FunctionCode {
		g.FunctionCode = false
	}

	// create temps
	temp0 := g.NewTemp()
	g.GeneratorNativeVariables.CompareStringsNative.TempFirstStringCompareGL = temp0
	temp1 := g.NewTemp()
	g.GeneratorNativeVariables.CompareStringsNative.TempSecondStringCompareGL = temp1
	temp2 := g.NewTemp()
	temp3 := g.NewTemp()
	temp4 := g.NewTemp()
	temp5 := g.NewTemp()
	temp6 := g.NewTemp()

	// create labels
	labelL0 := g.NewLabel()
	labelL1 := g.NewLabel()
	labelL2 := g.NewLabel()
	labelL3 := g.NewLabel()
	labelL4 := g.NewLabel()

	// code
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("void _compare_stringsGL() {\n"))

	// t2 = t0;
	g.AssignTemp(temp2, temp0)
	// t3 = t1;
	g.AssignTemp(temp3, temp1)
	// t4 = heap[(int) t2];
	g.GetHeap(temp4, temp2)
	// t5 = heap[(int) t3];
	g.GetHeap(temp5, temp3)
	// add label
	g.AddLabel(labelL0)
	// if(t4 > t5) goto L2;
	g.AddIf(temp4, temp5, ">", labelL2)
	// if(t4 == t5) goto L1;
	g.AddIf(temp4, temp5, "==", labelL1)
	// goto L3;
	g.GoTo(labelL3)
	// add label
	g.AddLabel(labelL1)
	// if(t4 == -1) goto L2;
	g.AddIf(temp4, "-1", "==", labelL3)
	// t2 = t2 + 1;
	g.GenArithmetic(temp2, temp2, "1", "+")
	// t3 = t3 + 1;
	g.GenArithmetic(temp3, temp3, "1", "+")
	// t4 = heap[(int) t2];
	g.GetHeap(temp4, temp2)
	// t5 = heap[(int) t3];
	g.GetHeap(temp5, temp3)
	// goto L0;
	g.GoTo(labelL0)
	// add label
	g.AddLabel(labelL3)
	// t6 = 0;
	g.AssignTemp(temp6, "0")
	// goto L4;
	g.GoTo(labelL4)
	// add label
	g.AddLabel(labelL2)
	// t6 = 1;
	g.AssignTemp(temp6, "1")
	// add label
	g.AddLabel(labelL4)

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

	return temp6
}
