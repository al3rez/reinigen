package datastore

import . "github.com/azbshiri/reinigen/domain_object"

type AppointmentDataStorer interface {
	Find(id int64) (*Appointment, error)
}
