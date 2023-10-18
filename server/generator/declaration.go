package generator

import (
	"server/compiler/compiler/values"
)

// Function to generate c3d when a variable is declared
func (g *Generator) GenDeclaration(tempString, id string, value *values.C3DPrimitive, isLoopContext bool) (int, int) {

	if isLoopContext {

		if value.GetType() == values.StringType && value.IsTemp {
			// add the variable to the stack
			g.SaveStack("P", tempString)
			return g.StackCounter, g.HeapCounter
		}

		if value.GetType() == values.StringType {
			tempPointer := g.GenString(value.GetValue())
			// add comment
			g.GenComment("Declaration of variable: ", id, ":", value.GetType(), "= ", "'", value.GetValue(), "'")
			// add the variable to the stack
			g.SaveStack("P", tempPointer)
			// return
			return g.StackCounter, g.HeapCounter
		}

		// add comment
		g.GenComment("Declaration of variable: ", id, ":", value.GetType(), "= ", value.GetValue())
		// add the variable to the stack
		g.SaveStack("P", value.GetValue())

	} else {

		if value.GetType() == values.StringType && value.IsTemp {
			// add the variable to the stack
			g.SaveStack("P", tempString)
			// increment pointer stack
			g.IncPointerStack()
			return g.StackCounter, g.HeapCounter
		}

		if value.GetType() == values.StringType {
			tempPointer := g.GenString(value.GetValue())
			// add comment
			g.GenComment("Declaration of variable: ", id, ":", value.GetType(), "= ", "'", value.GetValue(), "'")
			// add the variable to the stack
			g.SaveStack("P", tempPointer)
			// increment pointer stack
			g.IncPointerStack()

			// return
			return g.StackCounter, g.HeapCounter
		}

		// add comment
		g.GenComment("Declaration of variable: ", id, ":", value.GetType(), "= ", value.GetValue())
		// add the variable to the stack
		g.SaveStack("P", value.GetValue())
		// increment pointer stack
		g.IncPointerStack()
	}

	// return
	return g.StackCounter, g.HeapCounter
}

// Function to generate c3d when a variable is declared
func (g *Generator) GenDeclarationVector(tempString, id string, isLoopContext bool) (int, int) {

	if isLoopContext {
		// add comment
		g.GenComment("Declaration of variable: ", id, ":", "Vector")
		// add the variable to the stack
		g.SaveStack("P", tempString)

		return g.StackCounter, g.HeapCounter

	} else {
		// add comment
		g.GenComment("Declaration of variable: ", id, ":", "Vector")
		// add the variable to the stack
		g.SaveStack("P", tempString)
		// increment pointer stack
		g.IncPointerStack()
	}
	// return
	return g.StackCounter, g.HeapCounter
}
