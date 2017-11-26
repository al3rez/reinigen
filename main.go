package main

import (
	"net/http"

	. "github.com/azbshiri/reinigen/datastore/postgres"
	"github.com/azbshiri/reinigen/delivery_mechanism/http/handler"
	"github.com/go-pg/pg"
	"github.com/go-zoo/bone"
)

func main() {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "postgres",
	})
	handler.AppointmentDataStore = NewPostgresAppointmentDataStore(db)

	router := bone.New()
	router.Post("/appointments", handler.Appointments["Create"])
	router.Get("/appointments/:id", handler.Appointments["Show"])
	http.ListenAndServe(":6000", router)
}
