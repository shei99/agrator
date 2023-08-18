package application

import (
	"context"
	"fmt"
	"time"

	"github.com/reactivex/rxgo/v2"

	"github.com/shei99/agrator/domain/model"
	"github.com/shei99/agrator/ports/outgoing"
)

type ShuntAnalggration struct {
	shuntRepository outgoing.ShuntRepository
}

func NewShuntAnalggration(batriumRepo outgoing.ShuntRepository) *ShuntAnalggration {
	return &ShuntAnalggration{shuntRepository: batriumRepo}
}

func (service *ShuntAnalggration) InnerWindowEvaluation(observable rxgo.Observable) {
	sndsub := observable.Filter(func(i interface{}) bool {
		// fmt.Println(i)
		return i.(model.ShuntWithViolation).InnerWindowViolation == false
	}, rxgo.WithBufferedChannel(10))

	gsnd := sndsub.GroupByDynamic(func(i rxgo.Item) string {
		return i.V.(model.ShuntWithViolation).Shunt.BatriumId.Id
	}, rxgo.WithBufferedChannel(10))

	gsnd.Map(func(_ context.Context, i interface{}) (interface{}, error) {
		groupedObservable := i.(rxgo.GroupedObservable)
		obs := groupedObservable.WindowWithTime(rxgo.WithDuration(time.Second * 5)).Observe()

		go service.evaluateGroupedWindowOfShuntInnerWindow(obs, groupedObservable)
		return i, nil
	}, rxgo.WithBufferedChannel(10)).Observe()
}

func (service *ShuntAnalggration) evaluateGroupedWindowOfShuntInnerWindow(
	ch <-chan rxgo.Item,
	groupedObservable rxgo.GroupedObservable,
) {
	for {
		foo := <-ch
		// fmt.Println(foo)
		if foo.V == nil {
			return
		}

		foo.V.(rxgo.Observable).
			Map(func(_ context.Context, i interface{}) (interface{}, error) {
				// fmt.Println("2nd", i, groupedObservable.Key)
				return i, nil
			}).
			Last().
			Map(func(_ context.Context, i interface{}) (interface{}, error) {
				fmt.Println("Saving", i, groupedObservable.Key)
				service.shuntRepository.SaveShunt(i.(model.ShuntWithViolation).Shunt)
				return i, nil
			}).
			Observe()
	}
}

func (service *ShuntAnalggration) CriticalWindowEvaluation(observable rxgo.Observable) {
	fstsub := observable.Filter(func(i interface{}) bool {
		return i.(model.ShuntWithViolation).InnerWindowViolation
	})
	fstsub.DoOnNext(func(i interface{}) {
		fmt.Println("Critical", i)
		service.shuntRepository.SaveShunt(i.(model.ShuntWithViolation).Shunt)
	})
}
