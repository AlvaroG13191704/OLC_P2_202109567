package values

type PRIMITIVE interface {
	GetValue() interface{}
	GetType() string
}

// INTEGER
type Integer struct {
	Value int64
}

func (i *Integer) GetValue() interface{} {
	return i.Value
}

func (i *Integer) GetType() string {
	return IntType
}

// FLOAT
type Float struct {
	Value float64
}

func (f *Float) GetValue() interface{} {
	return f.Value
}

func (f *Float) GetType() string {
	return FloatType
}

// STRING
type String struct {
	Value string
}

func (s *String) GetValue() interface{} {
	return s.Value
}

func (s *String) GetType() string {
	return StringType
}

// BOOLEAN
type Boolean struct {
	Value bool
}

func (b *Boolean) GetValue() interface{} {
	return b.Value
}

func (b *Boolean) GetType() string {
	return BooleanType
}

// CHARACTER
type Character struct {
	Value string
}

func (c *Character) GetValue() interface{} {
	return c.Value
}

func (c *Character) GetType() string {
	return CharType
}

// NIL
type Nil struct {
	Value interface{}
}

func (n *Nil) GetValue() interface{} {
	return n.Value
}

func (n *Nil) GetType() string {
	return NilType
}

// C3D Primitive
type C3DPrimitive struct {
	Value    string
	Temp     string
	TypeData string
	IsTemp   bool
}

func (c *C3DPrimitive) GetValue() string {
	return c.Value
}

func (c *C3DPrimitive) GetType() string {

	return c.TypeData
}

func (c *C3DPrimitive) GetTemp() string {
	return c.Temp
}

func (c *C3DPrimitive) GetIsTemp() bool {
	return c.IsTemp
}

func NewC3DPrimitive(value string, temp string, typeData string, isTemp bool) *C3DPrimitive {
	return &C3DPrimitive{
		Value:    value,
		Temp:     temp,
		TypeData: typeData,
		IsTemp:   isTemp,
	}
}
