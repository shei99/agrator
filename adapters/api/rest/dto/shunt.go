package dto

type CurrentShunt struct {
	SystemId     uint16  `json:"systemId,omitempty"`
	HubId        uint16  `json:"hubId,omitempty"`
	ShuntTemp    uint8   `json:"shuntTemp,omitempty"`
	ShuntVoltage float32 `json:"shuntVoltage,omitempty"`
}
