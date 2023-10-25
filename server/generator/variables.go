package generator

// This struct will work to manage the temp variables for eac native function
type GeneratorNativeVariables struct {
	// print
	PrintNative PrintNative
	// arithmetic
	ZeroNative ZeroNative
	// concat strings
	ConcatNative ConcatNative
	// print bool
	PrintBoolNative PrintBoolNative
	// compare strings
	CompareStringsNative CompareStringsNative
	// logical
	LogicalNative LogicalNative
	// embebbed
	EmbebbedNative EmbebbedNative
}

type EmbebbedNative struct {
	IsEmbebbedFunc1                     bool   // if is embebbed func
	IsEmbebbedFunc2                     bool   // if is embebbed func
	TempDigitStringConvertoToIntOrFloat string // temp to convert string to int or float
	TempDigitToString                   string // temp to convert string to bool
	TempReturnValue1                    string // temp to return value
	TempReturnValue2                    string // temp to return value
}

type PrintNative struct {
	TempToPrint             string // temp to print
	TempToEvaluateEndString string // temp to evaluate the end of the string
}

type ZeroNative struct {
	TempDivisionZero      string // temp to evaluate division zero
	IsPosibleDivisionZero bool   // if is posible division zero
}

type ConcatNative struct {
	IsFirstTimeConcat      bool   // if is concat string
	IsFirstTimePrintString bool   // if is print string
	TempInitConcat         string // temp to concat strings
	TempFirstConcat        string // temp to concat strings
	TempSecondConcat       string // temp to concat strings
}

type PrintBoolNative struct {
	TempBoolFunc string // temp to bool func
	IsBoolFunc   bool   // if is bool func
}

type CompareStringsNative struct {
	TempFirstStringCompare     string // temp to compare string func
	TempSecondStringCompare    string // temp to compare string func
	TempFirstStringCompareGLE  string // temp to compare string func
	TempSecondStringCompareGLE string // temp to compare string func
	IsCompareStringFuncGLE     bool   // if is compare string func
	TempFirstStringCompareGL   string // temp to compare string func
	TempSecondStringCompareGL  string // temp to compare string func
	IsCompareStringFuncGL      bool   // if is compare string func
	IsCompareStringFunc        bool   // if is compare string func
}

type LogicalNative struct {
	// and
	IsAndFunc      bool   // if is logical func
	TempAndPointer string // temp to logical func
	Temp1And       string // temp to logical func
	Temp2And       string // temp to logical func
	// or
	IsOrFunc      bool   // if is logical func
	TempOrPointer string // temp to logical func
	Temp1Or       string // temp to logical func
	Temp2Or       string // temp to logical func
	// not
	IsNotFunc     bool   // if is logical func
	TempNoPointer string // temp to logical func
	TempNot       string // temp to logical func
}
