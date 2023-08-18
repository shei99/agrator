package dto

type Cell struct {
	Id          uint8   `json:"id,omitempty"`
	MinCellVolt float32 `json:"minCellVolt,omitempty"`
	MaxCellVolt float32 `json:"maxCellVolt,omitempty"`
	CellTemp    uint8   `json:"cellTemp,omitempty"`
	Status      uint8   `json:"status,omitempty"`
}
