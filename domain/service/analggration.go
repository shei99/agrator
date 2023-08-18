package service

import "github.com/reactivex/rxgo/v2"

type Analggration interface {
	InnerWindowEvaluation(observable rxgo.Observable)
	CriticalWindowEvaluation(observable rxgo.Observable)
}
