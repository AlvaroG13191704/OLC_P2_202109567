package generator

import (
	"fmt"
	"server/compiler/compiler/values"
)

// Function to generate c3d when a variable is declared
func (g *Generator) GenDeclaration(tempString, id string, value *values.C3DPrimitive, isLoopContext bool, relativePointer int) (int, int) {

	// evaluate if it's inside a function
	if g.FunctionCode {

		if isLoopContext {

			if value.GetType() == values.StringType && value.IsTemp {
				g.GenComment("Declaration of variable inside function")
				// add the variable to the stack
				temp := g.NewTemp()
				// assign temp
				g.GenArithmetic(temp, "P", fmt.Sprintf("%d", relativePointer), "+")
				// save the value
				g.SaveStack(temp, tempString)

				return relativePointer, g.HeapCounter
			}

			if value.GetType() == values.StringType {
				tempPointer := g.GenString(value.GetValue())
				g.GenComment("Declaration of variable inside function")
				// add the variable to the stack
				temp := g.NewTemp()
				// assign temp
				g.GenArithmetic(temp, "P", fmt.Sprintf("%d", relativePointer), "+")
				// add the variable to the stack
				g.SaveStack(temp, tempPointer)
				// return
				return relativePointer, g.HeapCounter
			}
			g.GenComment("Declaration of variable inside function")
			// add the variable to the stack
			temp := g.NewTemp()
			// assign temp
			g.GenArithmetic(temp, "P", fmt.Sprintf("%d", relativePointer), "+")
			// save the value
			g.SaveStack(temp, value.GetValue())

			return relativePointer, g.HeapCounter

		} else {

			if value.GetType() == values.StringType && value.IsTemp {
				g.GenComment("Declaration of variable inside function")
				temp := g.NewTemp()
				// assign temp
				g.GenArithmetic(temp, "P", fmt.Sprintf("%d", relativePointer), "+")
				// save the value
				g.SaveStack(temp, tempString)
				return relativePointer, g.HeapCounter
			}

			if value.GetType() == values.StringType {
				tempPointer := g.GenString(value.GetValue())
				g.GenComment("Declaration of variable inside function")
				// add the variable to the stack
				temp := g.NewTemp()
				// assign temp
				g.GenArithmetic(temp, "P", fmt.Sprintf("%d", relativePointer), "+")
				// add the variable to the stack
				g.SaveStack(temp, tempPointer)
				// return
				return relativePointer, g.HeapCounter
			}
			g.GenComment("Declaration of variable inside function")
			// add the variable to the stack
			temp := g.NewTemp()
			// assign temp
			g.GenArithmetic(temp, "P", fmt.Sprintf("%d", relativePointer), "+")
			// add the variable to the stack
			g.SaveStack(temp, value.GetValue())

			return relativePointer, g.HeapCounter
		}

	} else {

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

			return g.StackCounter, g.HeapCounter

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

			return g.StackCounter, g.HeapCounter
		}
	}

}

// Function to generate c3d when a variable is declared
func (g *Generator) GenDeclarationVector(tempString, id string, isLoopContext bool, relativePointer int) (int, int) {

	if g.FunctionCode {
		if isLoopContext {
			// add comment
			g.GenComment("Declaration of variable: ", id, ":", "Vector inside fucntion")
			temp := g.NewTemp()
			// assign temp
			g.GenArithmetic(temp, "P", fmt.Sprintf("%d", relativePointer), "+")
			// save the value
			g.SaveStack(temp, tempString)

			return relativePointer, g.HeapCounter

		} else {
			// add comment
			g.GenComment("Declaration of variable: ", id, ":", "Vector inside fucntion")
			temp := g.NewTemp()
			// assign temp
			g.GenArithmetic(temp, "P", fmt.Sprintf("%d", relativePointer), "+")
			// save the value
			g.SaveStack(temp, tempString)

			return relativePointer, g.HeapCounter
		}

	} else {

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

			return g.StackCounter, g.HeapCounter
		}
	}

}
