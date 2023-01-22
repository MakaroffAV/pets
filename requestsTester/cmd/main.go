package main

import (
	"log"
	"net/http"
	"requestsTester/internal/h"
)

func main() {

	/*
		Closed:	False
		Author:	Makarov Aleksei
		Target:	Program start point
	*/

	http.HandleFunc("/", h.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
