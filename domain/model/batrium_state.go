package model

type BatriumState struct {
	CellInnerWindow     Window
	CellCriticalWindow  Window
	ShuntInnerWindow    Window
	ShuntCriticalWindow Window
}

func NewBatriumState() *BatriumState {
	return &BatriumState{}
}
