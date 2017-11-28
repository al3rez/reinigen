package postgres_test

import (
	"testing"
	"time"

	"aahframework.org/test.v0/assert"
	. "github.com/azbshiri/reinigen/datastore/postgres"
	. "github.com/azbshiri/reinigen/domain_object"
	"github.com/go-pg/pg"
)

var db = pg.Connect(&pg.Options{
	User:     "postgres",
	Password: "postgres",
})

func TestPostgresAppointmentDataStore_Find(t *testing.T) {
	var err error
	date := time.Now()
	err = db.CreateTable(&Appointment{}, nil)
	if err != nil {
		db.DropTable(&Appointment{}, nil)
		db.CreateTable(&Appointment{}, nil)
	}
	err = db.Insert(NewAppointment(WithID(1), WithCreatedAt(date)))
	if err != nil {
		panic(err)
	}

	dataStore := NewPostgresAppointmentDataStore(db)

	appointment, _ := dataStore.Find(1)

	assert.Equal(t, appointment.CreatedAt.Day(), date.Day())
	assert.Equal(t, appointment.CreatedAt.Month(), date.Month())
	assert.Equal(t, appointment.CreatedAt.Year(), date.Year())
}

func TestPostgresAppointmentDataStore_Find_NotFound(t *testing.T) {
	var err error
	date := time.Now()
	err = db.CreateTable(&Appointment{}, nil)
	if err != nil {
		db.DropTable(&Appointment{}, nil)
		db.CreateTable(&Appointment{}, nil)
	}
	err = db.Insert(NewAppointment(WithID(1), WithCreatedAt(date)))
	if err != nil {
		panic(err)
	}

	dataStore := NewPostgresAppointmentDataStore(db)

	_, err = dataStore.Find(2)

	assert.Equal(t, err.Error(), "Appointment not found")
}

func TestPostgresAppointmentDataStore_All(t *testing.T) {
	var err error
	date := time.Now()
	err = db.CreateTable(&Appointment{}, nil)
	if err != nil {
		db.DropTable(&Appointment{}, nil)
		db.CreateTable(&Appointment{}, nil)
	}
	err = db.Insert(NewAppointment(WithID(1), WithCreatedAt(date)))
	err = db.Insert(NewAppointment(WithID(2), WithCreatedAt(date)))
	err = db.Insert(NewAppointment(WithID(3), WithCreatedAt(date)))
	if err != nil {
		panic(err)
	}

	dataStore := NewPostgresAppointmentDataStore(db)

	appointments, err := dataStore.All()

	assert.Equal(t, len(appointments), 3)
}
