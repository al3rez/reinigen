package use_case

import (
	"github.com/pkg/errors"

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

func IndexAppointment(dataStore AppointmentDataStorer) ([]*Appointment, error) {
	appointments, err := dataStore.All()
	if err != nil {
		return nil, errors.Wrap(err, "IndexAppointment")
	}
	return appointments, nil
}
