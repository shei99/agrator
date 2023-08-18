package model

type Cell struct {
	Id       uint8
	CellVolt float32
	CellTemp uint8
	Status   uint8
}

func NewCell() *Cell {
	return &Cell{}
}
