package repository

import (
	"sort"

	. "github.com/azbshiri/reinigen/domain_object"
	"github.com/pkg/errors"
)

type AppointmentRepository struct {
	root []*Appointment
}

func (ar *AppointmentRepository) Len() int {
	return len(ar.root)
}

func (ar *AppointmentRepository) Swap(i, j int) {
	ar.root[i], ar.root[j] = ar.root[j], ar.root[i]
}

func (ar *AppointmentRepository) Less(i, j int) bool {
	return ar.root[i].ID < ar.root[j].ID
}

func (ar *AppointmentRepository) Create(attrs ...AppointmentAttr) *Appointment {
	appointment := NewAppointment(attrs...)
	ar.root = append(ar.root, appointment)
	return appointment
}

func (ar *AppointmentRepository) Update(id int64, attrs ...AppointmentAttr) (*Appointment, error) {
	a, err := ar.Find(id)
	if err != nil {
		return nil, errors.Wrap(err, "Update")
	}

	for _, attr := range attrs {
		a = attr(a)
	}
	return a, nil
}

func (ar *AppointmentRepository) Delete(id int64) (*Appointment, error) {
	i, err := ar.IndexOf(id)
	if err != nil {
		return nil, errors.Wrap(err, "Delete")
	}
	a := ar.root[i]
	ar.root = append(ar.root[:i], ar.root[i+1:]...)
	return a, nil
}

func (ar *AppointmentRepository) All() []*Appointment {
	sort.Sort(ar)
	return ar.root
}

func (ar *AppointmentRepository) Find(id int64) (*Appointment, error) {
	for _, a := range ar.root {
		if a.ID == id {
			return a, nil
		}
	}
	return nil, errors.New("Appointment not found")
}

func (ar *AppointmentRepository) First() *Appointment {
	sort.Sort(ar)
	return ar.root[0]
}

func (ar *AppointmentRepository) IndexOf(id int64) (int, error) {
	for i, a := range ar.root {
		if a.ID == id {
			return i, nil
		}
	}
	return 0, errors.New("Appointment not found")
}

func (ar *AppointmentRepository) Last() *Appointment {
	sort.Sort(ar)
	return ar.root[len(ar.root)-1]
}

func (ar *AppointmentRepository) Clear() []*Appointment {
	deleted := ar.root
	ar.root = nil
	return deleted
}
