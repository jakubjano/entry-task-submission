package endpoints

import (
	"case_study/logic"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type EventRouter struct {
	EventList *logic.EventList
}

func (e *EventRouter) PostEventHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Other methods than POST not allowed"))
		return
	}

	w.Header().Set("Content-type", "application/json")
	event := logic.Event{
		Name:      "click",
		CreatedAt: time.Now().Unix(),
	}
	jbytes, err := json.Marshal(event)
	if err != nil {
		panic(err)
	}
	w.Write(jbytes)

	e.EventList.SaveEvent(event)
	fmt.Println(e.EventList)

	return
}

func (e *EventRouter) GetEventHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Other methods than GET not allowed"))
		return
	}

	from, to, t, err := ParseQueryParams(r.URL.Query())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error parsing data: %s", err.Error())))
		return
	}
	clicks := e.EventList.AggregateEvent(from, to, t)
	clicksBytes, err := json.Marshal(clicks)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(clicksBytes)
	fmt.Println(string(clicksBytes))

}
