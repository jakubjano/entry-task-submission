package main

import (
	"case_study/views"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	homeView *views.View
)

type Events struct {
	Name      string `json:"name"`
	CreatedAt int64  `json:"timestamp"`
}

var eventList []Events

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func NewEvent(w http.ResponseWriter, r *http.Request) {
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

	eventList = append(eventList, event)
	fmt.Println(eventList)
	return
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/click", NewEvent)
	log.Fatal(http.ListenAndServe(":8080", r))
}
