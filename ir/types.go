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

	Offset int
	Type   string
}

type Struct struct {
	Name string
	Size int

	TotalFixedFieldSize int
	FixedFields         []*Field
	DynamicFields       []*Field
}

type Alias struct {
	Name string

	OriginalType string
}

type IR struct {
	FileName string

	Structs []*Struct
	Enums   []*Enum
	Aliases []*Alias
}
