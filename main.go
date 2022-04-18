package main

import (
	"case_study/endpoints"
	"case_study/logic"
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	el := &logic.EventList{}
	server := &endpoints.EventRouter{
		el,
	}
	// time intervals for postman requests
	//fmt.Println(time.Now().Unix())
	//fmt.Println(time.Now().Add(-time.Minute * 100).Unix())
	r.HandleFunc("/click", server.PostEventHandler)
	r.HandleFunc("/aggregate", server.GetEventHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
