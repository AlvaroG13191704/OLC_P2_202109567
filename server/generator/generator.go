package generator

type Generator struct {
	Temp          int           // manage the temp variables
	Label         int           // manage the labels
	Code          []interface{} // list of code
	TempList      []interface{} // list of temp variables
	Compiled      []interface{} // list of compiled code
	FinalCode     []interface{} // list of final code
	FuncCode      []interface{}
	BreakLabel    string // break label
	ContinueLabel string // continue label
	MainCode      bool
}

func NewGenerator() *Generator {
	return &Generator{
		Temp:          0,
		Label:         0,
		BreakLabel:    "",
		ContinueLabel: "",
		MainCode:      true,
	}
}

func (g Generator) GetCode() []interface{} {
	return g.Code
}

func (g Generator) GetFinalCode() []interface{} {
	return g.FinalCode
}

// add break lvl
func (g *Generator) AddBreak(lvl string) {
	g.BreakLabel = lvl
}

// add continue lvl
func (g *Generator) AddContinue(lvl string) {
	g.ContinueLabel = lvl
}

// add a \n
func (g *Generator) AddNewLine() {
	if g.MainCode {
		g.Code = append(g.Code, "\n")
	} else {
		g.FuncCode = append(g.FuncCode, "\n")
	}
}
