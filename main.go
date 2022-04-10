package main

import (
	"case_study/views"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	homeView *views.View
)

type Events struct {
	Name      string `json:"name"`
	CreatedAt int64  `json:"timestamp"`
	//ID        int64  `json:"id"`
}

type eventList []Events

//func home(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "text/html")
//	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
//	if err != nil {
//		panic(err)
//	}
//}

func (e *eventList) saveEvent(event Events) {
	//todo validate events saved
	*e = append(*e, event)
}

//
func (e *eventList) aggregateEvent(from int64, to int64, t int64) (clicks []click) {
	//fmt.Println("e from aggregate events: ")
	//fmt.Println(e)
	//fmt.Println("pointer to e")
	//fmt.Println(*e)
	var subset []Events
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

	for i, j := from, 0; i <= to; i, j = i+t, j+1 {

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

type click struct {
	Click int `json:"click"`
}

// handler post
func (e *eventList) PostEventHandler(w http.ResponseWriter, r *http.Request) {
	//todo validate methods and  routes
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Other methods than POST not allowed"))
		return
	}

	var event Events
	w.Header().Set("Content-type", "application/json")
	event = Events{
		Name:      "click",
		CreatedAt: time.Now().Unix(),
	}
	jbytes, err := json.Marshal(event)
	if err != nil {
		panic(err)
	}
	w.Write(jbytes)

	e.saveEvent(event)
	fmt.Println(e)

	return
}

func (e *eventList) GetEventHandler(w http.ResponseWriter, r *http.Request) {
	//todo validate methods and  routes

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Other methods than GET not allowed"))
		return
	}

	from, to, t := parseQueryParams(r.URL.Query())
	clicks := e.aggregateEvent(from, to, t)
	clicksBytes, err := json.Marshal(clicks)
	if err != nil {
		panic(err)
	}
	w.Write(clicksBytes)
	fmt.Println(string(clicksBytes))

}

//func (e *eventList) Server(w http.ResponseWriter, r *http.Request) {
//	//todo switch for validating routes and matching path
//}
//
//func (e *eventList) wrongPath(w http.ResponseWriter, r *http.Request) {
//
//}

func parseQueryParams(params url.Values) (from, to, t int64) {

	from = stringConv(params.Get("from"))
	to = stringConv(params.Get("to"))
	t = stringConv(params.Get("t"))

	//todo validate these query params
	// validate here or in function implementing logic of aggregation ?
	return from, to, t

}

func stringConv(s string) (i int64) {

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	r := http.NewServeMux()
	server := &eventList{}
	//fmt.Println(time.Now().Unix())
	//fmt.Println(time.Now().Add(-time.Minute * 100).Unix())
	r.HandleFunc("/click", server.PostEventHandler)
	r.HandleFunc("/aggregate", server.GetEventHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
