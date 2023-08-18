package application

import (
	"context"
	"fmt"
	"time"

	"github.com/reactivex/rxgo/v2"

	"github.com/shei99/agrator/domain/model"
	"github.com/shei99/agrator/ports/outgoing"
)

type CellAnalggration struct {
	cellRepository outgoing.CellRepository
}

func NewCellAnalggration(batriumRepo outgoing.CellRepository) *CellAnalggration {
	return &CellAnalggration{cellRepository: batriumRepo}
}

func (service *CellAnalggration) InnerWindowEvaluation(observable rxgo.Observable) {
	sndsub := observable.Filter(func(i interface{}) bool {
		// fmt.Println(i)
		return i.(model.CellnodeWithViolation).InnerWindowViolation == false
	}, rxgo.WithBufferedChannel(10))

	gsnd := sndsub.GroupByDynamic(func(i rxgo.Item) string {
		return i.V.(model.CellnodeWithViolation).Cellnode.BatriumId.Id
	}, rxgo.WithBufferedChannel(10))

	gsnd.Map(func(_ context.Context, i interface{}) (interface{}, error) {
		groupedObservable := i.(rxgo.GroupedObservable)
		obs := groupedObservable.WindowWithTime(rxgo.WithDuration(time.Second * 5)).Observe()

		go service.evaluateGroupedWindowOfCellInnerWindow(obs, groupedObservable)
		return i, nil
	}, rxgo.WithBufferedChannel(10)).Observe()
}

func (service *CellAnalggration) evaluateGroupedWindowOfCellInnerWindow(
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
				service.cellRepository.SaveCellnode(i.(model.CellnodeWithViolation).Cellnode)
				return i, nil
			}).
			Observe()
	}
}

func (service *CellAnalggration) CriticalWindowEvaluation(observable rxgo.Observable) {
	fstsub := observable.Filter(func(i interface{}) bool {
		return i.(model.CellnodeWithViolation).InnerWindowViolation
	})
	fstsub.DoOnNext(func(i interface{}) {
		fmt.Println("Critical", i)
		service.cellRepository.SaveCellnode(i.(model.CellnodeWithViolation).Cellnode)
	})
}
