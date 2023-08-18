package dto

type CellCriticalWindow struct {
	SystemId                  uint16  `json:"systemId,omitempty"`
	HubId                     uint16  `json:"hubId,omitempty"`
	ControlCriticalCellVoltLo float32 `json:"controlCriticalCellVoltLo,omitempty"`
	ControlCriticalCellVoltHi float32 `json:"controlCriticalCellVoltHi,omitempty"`
	ControlCriticalCellTempLo uint8   `json:"controlCriticalCellTempLo,omitempty"`
	ControlCriticalCellTempHi uint8   `json:"controlCriticalCellTempHi,omitempty"`
}
