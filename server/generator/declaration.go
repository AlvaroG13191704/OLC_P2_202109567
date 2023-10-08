package generator

import (
	"server/compiler/compiler/values"
)

// Function to generate c3d when a variable is declared
func (g *Generator) GenDeclaration(tempString, id string, value *values.C3DPrimitive) (string, int, int) {

	if value.GetType() == values.StringType {
		tempPointer := g.GenString(value.GetValue())
		// add comment
		// add comment
		g.GenComment("Declaration of variable: ", id, ":", value.GetType(), ": ", "'", value.GetValue(), "'")
		// add the variable to the stack
		g.SaveStack("P", tempPointer)
		// increment pointer stack
		g.IncPointerStack()

		// return
		return tempPointer, g.StackCounter, g.HeapCounter
	}

	// add comment
	g.GenComment("Declaration of variable: ", id, ":", value.GetType(), ": ", value.GetValue())
	// add the variable to the stack
	g.SaveStack("P", value.GetValue())
	// increment pointer stack
	g.IncPointerStack()

	// return
	return tempString, g.StackCounter, g.HeapCounter
}
