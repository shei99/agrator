package model

type Window struct {
	CellVoltHi float32
	CellVoltLo float32
	CellTempHi uint8
	CellTempLo uint8
}

func NewWindow() *Window {
	return &Window{}
}
