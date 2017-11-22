package handler

import (
	"net/http"

	. "github.com/azbshiri/first-example/domain_object"
	. "github.com/azbshiri/first-example/use_case"
	"github.com/pquerna/ffjson/ffjson"
)

var Appointments = map[string]http.HandlerFunc{
	"Create": Create(),
}

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var appointemnt Appointment
		err := ffjson.NewDecoder().DecodeReader(r.Body, &appointemnt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		appointment, err := CreateAppointment(WithID(appointemnt.ID),
			WithCreatedAt(appointemnt.CreatedAt),
			WithUpdatedAt(appointemnt.UpdatedAt))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		ffjson.NewEncoder(w).Encode(appointment)
		return
	}
}
