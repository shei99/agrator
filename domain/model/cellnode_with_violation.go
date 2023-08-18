package model

type CellnodeWithViolation struct {
	Cellnode                Cellnode
	InnerWindowViolation    bool
	CriticalWindowViolation bool
}

func NewCellNodeWithViolation(cellnode Cellnode) *CellnodeWithViolation {
	return &CellnodeWithViolation{Cellnode: cellnode}
}
