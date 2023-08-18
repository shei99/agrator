package dto

type CellInnerWindow struct {
	SystemId                   uint16  `json:"systemId,omitempty"`
	HubId                      uint16  `json:"hubId,omitempty"`
	ControlChargeCellVoltHi    float32 `json:"controlChargeCellVoltHi,omitempty"`
	ControlDischargeCellVoltLo float32 `json:"controlDischargeCellVoltLo,omitempty"`
	ControlChargeCellTempLo    uint8   `json:"controlChargeCellTempLo,omitempty"`
	ControlChargeCellTempHi    uint8   `json:"controlChargeCellTempHi,omitempty"`
}
