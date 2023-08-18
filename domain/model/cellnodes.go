package model

type Cellnode struct {
	BatriumId BatriumIdentifier
	Cells     []Cell
}

func NewCellnode() *Cellnode {
	return &Cellnode{}
}
