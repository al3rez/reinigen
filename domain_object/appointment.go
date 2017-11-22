package domain_object

import (
	"time"
)

// Appointment ...
type Appointment struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// AppointmentAttr ...
type AppointmentAttr func(Appointment) Appointment

// NewAppointment ...
func NewAppointment(attrs ...AppointmentAttr) *Appointment {
	appointment := Appointment{}
	for _, attr := range attrs {
		appointment = attr(appointment)
	}
	return &appointment
}

// WithID ...
var WithID = func(id int64) AppointmentAttr {
	return func(a Appointment) Appointment {
		a.ID = id
		return a
	}
}

// WithCreatedAt ...
var WithCreatedAt = func(createdAt time.Time) AppointmentAttr {
	return func(a Appointment) Appointment {
		a.CreatedAt = createdAt
		return a
	}
}

// WithUpdatedAt ...
var WithUpdatedAt = func(updatedAt time.Time) AppointmentAttr {
	return func(a Appointment) Appointment {
		a.UpdatedAt = updatedAt
		return a
	}
}
