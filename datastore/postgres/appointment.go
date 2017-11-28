package postgres

import (
	. "github.com/azbshiri/reinigen/datastore"
	. "github.com/azbshiri/reinigen/domain_object"
	"github.com/go-pg/pg"
	"github.com/pkg/errors"
)

type PostgresAppointmentDataStore struct {
	db *pg.DB
}

func NewPostgresAppointmentDataStore(db *pg.DB) AppointmentDataStorer {
	dataStore := PostgresAppointmentDataStore{}
	dataStore.db = db
	return &dataStore
}

func (p *PostgresAppointmentDataStore) Find(id int64) (*Appointment, error) {
	appointment := Appointment{ID: id}
	err := p.db.Select(&appointment)
	if err != nil {
		return nil, errors.New("Appointment not found")
	}

	return &appointment, nil
}

func (p *PostgresAppointmentDataStore) All() ([]*Appointment, error) {
	var appointments []*Appointment
	err := p.db.Model(&appointments).Select()
	if err != nil {
		return nil, errors.Wrap(err, "PostgresAppointmentDataStore.All")
	}

	return appointments, nil
}
