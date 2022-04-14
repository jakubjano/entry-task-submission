package logic

import (
	"fmt"
)

type click struct {
	Click int `json:"click"`
}

type Event struct {
	Name      string `json:"name"`
	CreatedAt int64  `json:"timestamp"`
	//ID        int64  `json:"id"`
}

type EventList []Event

func (e *EventList) SaveEvent(event Event) {
	//todo validate events saved
	*e = append(*e, event)
}

func (e *EventList) AggregateEvent(from int64, to int64, t int64) (clicks []click) {
	//fmt.Println("e from aggregate events: ")
	//fmt.Println(e)
	//fmt.Println("pointer to e")
	//fmt.Println(*e)
	var subset []Event
	for _, item := range *e {
		//fmt.Println("item from aggregate events")
		//fmt.Println(item)
		if item.CreatedAt >= from && item.CreatedAt <= to {
			subset = append(subset, item)
		}
	}
	fmt.Println("subset")
	fmt.Println(subset)
	// j is for indexing in slice eventCount
	// i adds up to the last unix timestamp in from to time range
	// if last iteration i+t > to redeclare t

	for i, j := from, 0; i < to; i, j = i+t, j+1 {
		if i+t > to {
			t = to - i
		}
		//todo last residual interval

		fmt.Println("from:", i)
		fmt.Println("to", i+t)

		counter := 0
		for _, item := range subset {
			if item.CreatedAt < i+t && item.CreatedAt >= i {
				counter += 1

			}

		}
		fmt.Println(counter)
		clicks = append(clicks, click{
			counter,
		})
	}

	return clicks
}
