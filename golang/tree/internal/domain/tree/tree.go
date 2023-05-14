package tree

type PayloadType interface {
	int32 | int64
}

type Node[P PayloadType] struct {
	Name    string      `json:"name"`
	Nodes   *[]*Node[P] `json:"nodes"`
	Payload P           `json:"payload"`
}
