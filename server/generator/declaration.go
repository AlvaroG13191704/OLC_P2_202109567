package generator

import (
	"fmt"
	"server/compiler/compiler/values"
)

// Function to generate c3d when a variable is declared
func (g *Generator) GenDeclaration(tempString, id string, value values.PRIMITIVE) (string, int, int) {

	// evaluate if it's a string
	if value.GetType() == values.StringType {
		// add comment
		g.GenComment("*** Declaration of a string ***")
		// generate a new temp and assing the heap pointer
		// tempString := g.NewTemp()
		g.AssignTemp(tempString, "H")
		// iterate the string and save the chars in the heap
		for _, char := range value.GetValue().(string) {
			// convert to ascii
			charString := fmt.Sprintf("%d", char)
			// save in the heap
			g.SaveHeap(charString)
			// increment pointer heap
			g.IncPointerHeap()
		}
		// save the end of the string
		g.SaveHeap("-1")
		// increment pointer heap
		g.IncPointerHeap()
		// add comment
		g.GenComment("Declaration of variable: ", id, ":", value.GetType(), "with value: ", "'", value.GetValue().(string), "'")
		// add the variable to the stack
		g.SaveStack(id, tempString)
		// increment pointer stack
		g.IncPointerStack()
		return tempString, g.StackCounter, g.HeapCounter
	}

	// value
	var valueString string

	fmt.Println("value.GetType(): ", value.GetType())

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

	case values.NilType:
		valueString = "9999999827968.00"
	}

	// create temp
	// tempString := g.NewTemp()
	// add comment
	g.GenComment("Declaration of variable: ", id, ":", value.GetType(), "with value: ", valueString)
	// assing to the temp the value
	g.AssignTemp(tempString, valueString)
	// add the variable to the stack
	g.SaveStack(id, tempString)
	// increment pointer stack
	g.IncPointerStack()

	// return
	return tempString, g.StackCounter, g.HeapCounter
}
