package generator

import "fmt"

func (g *Generator) GenPrintBoolFunc() {
	// set false main
	g.MainCode = false
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

/*
c3d code
t1 = first string address
t2 = second string address
t3 = equal or not temp
t4 = iterator first string
t5 = iterator second string

void _compare_strings() {
	t3 = 1; // Initialize t3 to true

	L1:
		t4 = heap[(int)t1];
		t5 = heap[(int)t2];

		if (t4 == -1) goto L2;
		if (t5 == -1) goto L3;

		if (t4 != t5) goto L3;

		t1 = t1 + 1;
		t2 = t2 + 1;
		goto L1;

	L2:
		if (t5 == -1) goto L4;
		t3 = 0; // Set t3 to false
		goto L4;

	L3:
		t3 = 0; // Set t3 to false

	L4:
		return;
}
*/

// Func to compare strings
func (g *Generator) GenComparationStringFunc() string {
	// set false main
	g.MainCode = false
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
	return temp3
}

/*
void _compare_strings() {
	t3 = 1; // Initialize t3 to true "less than"

	L1:
		t4 = heap[(int)t1];
		t5 = heap[(int)t2];

		if (t4 == -1) goto L2;
		if (t5 == -1) goto L3;

		if (t4 > t5) goto L3;
		if (t4 < t5) goto L2;

		t1 = t1 + 1;
		t2 = t2 + 1;
		goto L1;

	L2:
		if (t5 == -1) goto L4;
		t3 = 0; // Set t3 to false "greater than"
		goto L4;

	L3:
		t3 = 0; // Set t3 to false "greater than"

	L4:
		return;
}
*/

// Func to compare strings
func (g *Generator) GenComparationStringFuncLT() string {
	// set false main
	g.MainCode = false
	// gen comment
	g.GenComment("Comparation String less or than")

	// generate labels
	labelL1 := g.NewLabel()
	labelL2 := g.NewLabel()
	label3 := g.NewLabel()
	labelL4 := g.NewLabel()
	// generate temps
	temp1 := g.NewTemp()
	g.GeneratorNativeVariables.CompareStringsNative.TempFirstStringCompareLT = temp1
	temp2 := g.NewTemp()
	g.GeneratorNativeVariables.CompareStringsNative.TempSecondStringCompareLT = temp2
	temp3 := g.NewTemp()
	temp4 := g.NewTemp()
	temp5 := g.NewTemp()
	// code
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("void _compare_strings_lt() {\n"))
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
	// if (t4 > t5) goto L3;
	g.AddIf(temp4, temp5, ">", label3)
	// if (t4 < t5) goto L2;
	g.AddIf(temp4, temp5, "<", labelL2)
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
	return temp3
}
