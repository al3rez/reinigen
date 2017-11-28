package handler

import (
	"net/http"
	"strconv"

	"github.com/pquerna/ffjson/ffjson"

	"github.com/go-zoo/bone"

	. "github.com/azbshiri/reinigen/datastore"
	. "github.com/azbshiri/reinigen/use_case"
)

var Appointments = map[string]http.HandlerFunc{
	"Index": Index(),
	"Show":  Show(),
}

var AppointmentDataStore AppointmentDataStorer

func Index() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		appointments, err := IndexAppointment(AppointmentDataStore)
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		ffjson.NewEncoder(w).Encode(appointments)
		return
	}
}

func Show() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(bone.GetValue(r, "id"))
		if err != nil {
			http.Error(w, "Invalid Id", 401)
			return
		}

		appointment, err := ShowAppointment(AppointmentDataStore, int64(id))
		if err != nil {
			http.Error(w, err.Error(), 404)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		ffjson.NewEncoder(w).Encode(appointment)
		return
	}
}
