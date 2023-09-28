package generator

import (
	"fmt"
	"server/compiler/compiler/values"
)

// Function to generate c3d when a variable is declared
func (g *Generator) GenDeclaration(id string, value values.PRIMITIVE) (string, int, int) {

	// value
	var valueString string

	// evaluate the type of the value
	switch value.GetType() {
	case values.BooleanType:
		// cast the value to float
		if value.GetValue().(bool) {
			valueString = "1"
		} else {
			valueString = "0"
		}
	case values.IntType:
		intvalue := value.GetValue().(int64)
		// conver to string
		valueString = fmt.Sprintf("%d", intvalue)

	case values.FloatType:
		floatvalue := value.GetValue().(float64)
		// conver to string
		valueString = fmt.Sprintf("%f", floatvalue)

	case values.CharType:
		charvalue := value.GetValue().(string)
		// conver to ascii
		valueString = fmt.Sprintf("%d", charvalue[0])
	}

	// create temp
	tempString := g.NewTemp()
	// assing to the temp the value
	g.AssignTemp(tempString, valueString)
	// add the variable to the stack
	g.VariableDeclaration(id, tempString)
	// increment pointer stack
	g.IncPointerStack()

	// return
	return tempString, g.StackCounter, g.HeapCounter
}
