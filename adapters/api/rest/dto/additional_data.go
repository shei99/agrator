package dto

type AdditionalData struct {
	SystemId uint16             `json:"systemId,omitempty"`
	HubId    uint16             `json:"hubId,omitempty"`
	Type     string             `json:"type,omitempty"`
	Data     map[string]float32 `json:"data,omitempty"`
}
