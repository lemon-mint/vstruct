package ir

type FieldType int

const (
	FieldType_RESERVED FieldType = iota

	FieldType_UINT
	FieldType_INT
	FieldType_FLOAT
	FieldType_BOOL
	FieldType_STRING
	FieldType_BYTES
	FieldType_ENUM
	FieldType_STRUCT
)

type Enum struct {
	Name string
	Size int

	Options []string
}

type Field struct {
	FT FieldType

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
