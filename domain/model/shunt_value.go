package model

type ShuntValue struct {
	BatriumId    BatriumIdentifier
	ShuntTemp    uint8
	ShuntVoltage float32
}

func NewShuntValue() *ShuntValue {
	return &ShuntValue{}
}
