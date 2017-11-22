package main

import (
	"net/http"

	"github.com/azbshiri/first-example/delivery_mechanism/http/handler"
	"github.com/go-zoo/bone"
)

func main() {
	router := bone.New()
	router.Post("/appointments", handler.Appointments["Create"])
	http.ListenAndServe(":6000", router)
}
