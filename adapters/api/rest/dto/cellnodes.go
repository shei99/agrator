package dto

type CurrentCellnode struct {
	SystemId uint16 `json:"systemId,omitempty"`
	HubId    uint16 `json:"hubId,omitempty"`
	Nodes    []Cell `json:"nodes,omitempty"`
}
