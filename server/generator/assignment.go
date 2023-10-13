package generator

import (
	"fmt"
	"server/compiler/compiler/values"
)

// Function to generate c3d when a variable is assign
func (g *Generator) GenAssignment(tempString, id string, pointerStack int, value *values.C3DPrimitive, expr *values.C3DPrimitive) {

	fmt.Println("Value to assign ->", value)
	if value.GetType() == values.StringType && expr.IsTemp {
		g.SaveStack(fmt.Sprintf("%d", pointerStack), expr.GetValue())

		return
	}

	if value.GetType() == values.StringType {
		tempPointer := g.GenString(expr.GetValue())
		// add comment
		g.GenComment("Assigment of variable: ", id, ":", value.GetType(), "=", "'", expr.GetValue(), "'")
		// add the variable to the stack
		g.SaveStack(fmt.Sprintf("%d", pointerStack), tempPointer)

		// return
		return
	}

	// add comment
	g.GenComment("Assigment of variable: ", id, ":", value.GetType(), "= ", "'", expr.GetValue(), "'")
	// add the variable to the stack
	g.SaveStack(fmt.Sprintf("%d", pointerStack), expr.GetValue())

}

// Function to generate c3d when a variable is assign += or -=
func (g *Generator) GenAssignmentPlusMiuns(target, stackDirection string, expr *values.C3DPrimitive) {

	if target == "+" && expr.GetType() == values.StringType {
		tempHeap := g.GenString(expr.GetValue())
		// gen code to concat
		if !g.GeneratorNativeVariables.ConcatNative.IsFirstTimeConcat {
			initPointer := g.ConcantStrings()
			g.GeneratorNativeVariables.ConcatNative.IsFirstTimeConcat = true
			// t1 = stack[ (int) 0];
			temp := g.NewTemp()
			g.AccessStack(temp, stackDirection)
			// gen code to assign the pointer
			g.GenConcatString(temp, tempHeap)

			//save the value
			g.SaveStack(stackDirection, initPointer)
		} else {
			// t1 = stack[ (int) 0];
			temp := g.NewTemp()
			g.AccessStack(temp, stackDirection)
			// gen code to assign the pointer
			g.GenConcatString(temp, tempHeap)

			//save the value
			g.SaveStack(stackDirection, g.GeneratorNativeVariables.ConcatNative.TempInitConcat)
		}
		return
	}
	// add comment
	g.GenComment("Assigment of variable: += or -= ")
	// t1 = stack[ (int) 0];
	temp := g.NewTemp()
	g.AccessStack(temp, stackDirection)
	// t2 = t1 + 1;
	temp2 := g.NewTemp()
	g.GenArithmetic(temp2, temp, expr.GetValue(), target)

	// save the value
	g.SaveStack(stackDirection, temp2)

}
