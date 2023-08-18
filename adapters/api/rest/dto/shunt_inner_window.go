package dto

type ShuntInnerWindow struct {
	SystemId                       uint16  `json:"systemId,omitempty"`
	HubId                          uint16  `json:"hubId,omitempty"`
	ControlChargeTargetLimpVolt    float32 `json:"controlChargeTargetLimpVolt,omitempty"`
	ControlDischargeTargetLimpVolt float32 `json:"controlDischargeTargetLimpVolt,omitempty"`
}
