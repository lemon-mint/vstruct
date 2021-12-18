package ir

type FieldType int

type TypeInfo struct {
	Size      int
	IsDynamic bool

	FieldType FieldType
}

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
	TypeInfo TypeInfo

	Name string

	Offset int
	Type   string
}

type Struct struct {
	Name string
	Size int

	TotalFixedFieldSize int
	FixedFields         []*Field

	DynamicHead             int
	DynamicFieldHeadOffsets []int
	DynamicFields           []*Field
}

type Alias struct {
	Name string

	OriginalType string
}

type CompileOptions struct {
	UseUnsafe bool
}

type IR struct {
	FileName string

	Options CompileOptions

	Structs []*Struct
	Enums   []*Enum
	Aliases []*Alias
}
