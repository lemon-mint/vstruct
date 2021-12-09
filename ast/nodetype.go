package ast

type NodeType int

const (
	NodeType_RESERVED NodeType = iota

	NodeType_STRUCT
	NodeType_ENUM
	NodeType_ALIAS

	NodeType_RAWTYPE
)

var NodeType_names = map[NodeType]string{
	NodeType_RESERVED: "RESERVED",

	NodeType_STRUCT: "STRUCT",
	NodeType_ENUM:   "ENUM",
	NodeType_ALIAS:  "ALIAS",

	NodeType_RAWTYPE: "RAWTYPE",
}

func (n NodeType) String() string {
	return NodeType_names[n]
}

func (n NodeType) Printf(format string, a ...interface{}) string {
	return n.String()
}
