package use_case

import (
	. "github.com/azbshiri/first-example/domain_object"
)

var appointments = []*Appointment{
	NewAppointment(WithID(1)),
	NewAppointment(WithID(2)),
	NewAppointment(WithID(3)),
	NewAppointment(WithID(4)),
}

// CreateAppointment ...
func CreateAppointment(attrs ...AppointmentAttr) (*Appointment, error) {
	appointemnt := NewAppointment(attrs...)
	return appointemnt, nil
}
