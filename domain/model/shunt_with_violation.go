package model

type ShuntWithViolation struct {
	Shunt                   ShuntValue
	InnerWindowViolation    bool
	CriticalWindowViolation bool
}

func NewShuntWithViolation(shunt ShuntValue) *ShuntWithViolation {
	return &ShuntWithViolation{Shunt: shunt}
}
