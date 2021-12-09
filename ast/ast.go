package ast

func NewFile(filename string) *File {
	return &File{
		Name:    filename,
		Nodes:   nil,
		Globals: make(map[string]*Node),
	}
}

type File struct {
	Name  string
	Nodes []*Node

	Globals map[string]*Node
}

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
	// RawType
	RawType *RawType
}

type Field struct {
	Name    string
	StrType string
	Type    *Node
}

type Struct struct {
	Fields []*Field
}

type Enum struct {
	Enums []string
}

type Alias struct {
	StrType string
	Type    *Node
}

type RawType struct {
	StrType string
	Type    string
}

func NewNode(nodeType NodeType) *Node {
	return &Node{
		Type: nodeType,
	}
}
