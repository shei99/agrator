package model

type AdditionalData struct {
	BatriumId BatriumIdentifier
	Type      string
	Data      map[string]float32
}

func NewAdditionalData() *AdditionalData {
	return &AdditionalData{}
}
