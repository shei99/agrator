package dto

type ShuntCriticalWindow struct {
	SystemId                   uint16  `json:"systemId,omitempty"`
	HubId                      uint16  `json:"hubId,omitempty"`
	ControlCriticalShuntVoltLo float32 `json:"controlCriticalShuntVoltLo,omitempty"`
	ControlCriticalShuntVoltHi float32 `json:"controlCriticalShuntVoltHi,omitempty"`
}
