package use_case

import (
	"errors"

	. "github.com/azbshiri/reinigen/datastore"
	. "github.com/azbshiri/reinigen/domain_object"
)

// ShowAppointment ...
func ShowAppointment(dataStore AppointmentDataStorer, id int64) (*Appointment, error) {
	appointment, err := dataStore.Find(id)
	if err != nil {
		return nil, errors.New("Appointment not found")
	}
	return appointment, nil
}
