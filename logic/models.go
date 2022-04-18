package logic

type click struct {
	Click int `json:"click"`
}

type Event struct {
	Name      string `json:"name"`
	CreatedAt int64  `json:"timestamp"`
}

type EventList []Event

// SaveEvent saves event to the EventList slice of type Event
func (e *EventList) SaveEvent(event Event) {
	*e = append(*e, event)
}

// AggregateEvent is used to select events specified by a time range from-to
// and to aggregate these events by time interval t.
// Returns aggregated events called clicks where count of events in each specified
// interval t is an item in a slice
func (e *EventList) AggregateEvent(from int64, to int64, t int64) (clicks []click) {
	// filter events from-to
	var subset []Event
	for _, item := range *e {
		if item.CreatedAt >= from && item.CreatedAt < to {
			subset = append(subset, item)
		}
	}

	for i := from; i < to; i = i + t {
		// This if statement handles last iteration where possibility of i+t
		// being greater than to arise
		if i+t > to {
			t = to - i
		}
		counter := 0
		for _, item := range subset {
			if item.CreatedAt < i+t && item.CreatedAt >= i {
				counter += 1
			}
		}
		clicks = append(clicks, click{
			counter,
		})
	}
	return clicks
}
