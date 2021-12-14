package ir

type FieldType int

const (
	FieldType_RESERVED FieldType = iota

	FieldType_Raw
	FieldType_Enum
	FieldType_Array // unused
	FieldType_Map   // unused
)

type Enum struct {
	Name string
	Size int

	Options []string
}

type Field struct {
	Name string
	Type FieldType

	RawType string
	//A int
	EnumType *Enum
	//M map[string]string
}

type Struct struct {
	Name   string
	Fields []*Field
}

type IR struct {
	FileName string

	Structs []*Struct
	Enums   []*Enum
}
