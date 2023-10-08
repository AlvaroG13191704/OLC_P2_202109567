package generator

import "fmt"

/*
c3d code
void _and() {
	if ( t1 ==  0) goto L3;
	if ( t2 ==  0) goto L3;
	t3 = 1;
	goto L4;
	L3:
	t3 = 0;
	L4:
return;
}
*/

// Generate func to compare and
func (g *Generator) GenAndFunc() {
	// set false main
	g.MainCode = false

	// generate labels
	labelEnd := g.NewLabel()
	labelTrue := g.NewLabel()

	// generate temp
	temp1 := g.NewTemp()
	g.GeneratorNativeVariables.LogicalNative.Temp1And = temp1
	temp2 := g.NewTemp()
	g.GeneratorNativeVariables.LogicalNative.Temp2And = temp2
	temp3 := g.NewTemp()
	g.GeneratorNativeVariables.LogicalNative.TempAndPointer = temp3
	// gen comment
	g.GenComment("And")

	// code
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("void _and() {\n"))
	// add if
	g.AddIf(temp1, "0", "==", labelEnd)
	// add if
	g.AddIf(temp2, "0", "==", labelEnd)
	// assign true to the temp
	g.AssignTemp(temp3, "1")
	// add goto
	g.GoTo(labelTrue)
	// add label
	g.AddLabel(labelEnd)
	// assign false to the temp
	g.AssignTemp(temp3, "0")
	// add label
	g.AddLabel(labelTrue)
	// return
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("return;\n"))
	// add end func
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("}\n"))

	g.MainCode = true

}

/*
void __or() {
	if ( t1 !=  0) goto L3;
	if ( t2 !=  0) goto L3;
	t3 = 0;
	goto L4;
	L3:
	t3 = 1;
	L4:
return;
}
*/

// Generate func to compare or
func (g *Generator) GenOrFunc() {
	// set false main
	g.MainCode = false

	// generate labels
	labelEnd := g.NewLabel()
	labelTrue := g.NewLabel()

	// generate temp
	temp1 := g.NewTemp()
	g.GeneratorNativeVariables.LogicalNative.Temp1Or = temp1
	temp2 := g.NewTemp()
	g.GeneratorNativeVariables.LogicalNative.Temp2Or = temp2
	temp3 := g.NewTemp()
	g.GeneratorNativeVariables.LogicalNative.TempOrPointer = temp3
	// gen comment
	g.GenComment("Or")

	// code
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("void _or() {\n"))
	// add if
	g.AddIf(temp1, "0", "!=", labelEnd)
	// add if
	g.AddIf(temp2, "0", "!=", labelEnd)
	// assign true to the temp
	g.AssignTemp(temp3, "0")
	// add goto
	g.GoTo(labelTrue)
	// add label
	g.AddLabel(labelEnd)
	// assign false to the temp
	g.AssignTemp(temp3, "1")
	// add label
	g.AddLabel(labelTrue)
	// return
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("return;\n"))
	// add end func
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("}\n"))

	g.MainCode = true

}

/*
void __not() {
	if ( t1 ==  0) goto L3;
	t2 = 0;
	goto L4;
	L3:
	t2 = 1;
	L4:
return;
}
*/

// Generate not
func (g *Generator) GenNotFunc() {
	// set false main
	g.MainCode = false

	// generate labels
	labelEnd := g.NewLabel()
	labelTrue := g.NewLabel()

	// generate temp
	temp1 := g.NewTemp()
	g.GeneratorNativeVariables.LogicalNative.TempNot = temp1
	temp2 := g.NewTemp()
	g.GeneratorNativeVariables.LogicalNative.TempNoPointer = temp2
	// gen comment
	g.GenComment("Not")

	// code
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("void _not() {\n"))
	// add if
	g.AddIf(temp1, "0", "==", labelEnd)
	// assign false to the temp
	g.AssignTemp(temp2, "0")
	// add goto
	g.GoTo(labelTrue)
	// add label
	g.AddLabel(labelEnd)
	// assign true to the temp
	g.AssignTemp(temp2, "1")
	// add label
	g.AddLabel(labelTrue)
	// return
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("return;\n"))
	// add end func
	g.FuncCode = append(g.FuncCode, fmt.Sprintf("}\n"))

	g.MainCode = true
}

// Generate logical func
func (g *Generator) GenLogical(left, right, temp1, temp2, temp3, funcName string) string {

	// gen comment
	g.GenComment(funcName)
	if funcName == "not" {
		// assing left to temp1
		g.AssignTemp(temp1, left)
	} else {
		// assing left to temp1
		g.AssignTemp(temp1, left)
		// assing right to temp2
		g.AssignTemp(temp2, right)
	}
	// call func
	g.Code = append(g.Code, fmt.Sprintf("_%s();\n", funcName))

	return temp3
}
