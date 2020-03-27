package main

import (
	"fmt"
	"log"
	"net/http"
	"psm/psmd"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	s := &psmd.HTTPServer{
		Router: router,
	}
	router.Handle("GET", "/message/subscribe/send", s.SendSubscribeMsg)
	fmt.Println("Service started successfully! Listening :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
