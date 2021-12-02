package ast

type NodeType int

const (
	NodeType_RESERVED NodeType = iota

	NodeType_STRUCT
	NodeType_ENUM
	NodeType_ALIAS
)

type Node struct {
	Type NodeType
	Name string
	File string
	Line int
	Col  int

	// Struct
	Struct *Struct
	// Enum
	Enum *Enum
	// Alias
	Alias *Alias
}

type Field struct {
	Name string
	Type *Node
}

type Struct struct {
	Fields []*Field
}

type Enum struct {
	Enums []string
}

type Alias struct {
	Type *Node
}
