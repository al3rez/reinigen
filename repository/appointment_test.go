package repository_test

import (
	"testing"
	"time"

	. "github.com/azbshiri/reinigen/domain_object"
	. "github.com/azbshiri/reinigen/repository"

	"github.com/stretchr/testify/assert"
)

func TestAppointmentRepository_Create(t *testing.T) {
	date := time.Now()
	expected := Appointment{}
	expected.ID = 1
	expected.CreatedAt = date
	expected.UpdatedAt = date

	r := AppointmentRepository{}
	actual := r.Create(WithID(1), WithCreatedAt(date), WithUpdatedAt(date))

	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.UpdatedAt, actual.UpdatedAt)
	assert.Equal(t, expected.CreatedAt, actual.CreatedAt)
}

func TestAppointmentRepository_Update(t *testing.T) {
	date := time.Now()
	r := AppointmentRepository{}

	r.Create(WithID(1), WithCreatedAt(date), WithUpdatedAt(date))
	r.Create(WithID(2), WithCreatedAt(date), WithUpdatedAt(date))
	r.Create(WithID(3), WithCreatedAt(date), WithUpdatedAt(date))

	actual, _ := r.Update(1, WithID(2))

	assert.Equal(t, int64(2), actual.ID)
}

func TestAppointmentRepository_Update_NotFound(t *testing.T) {
	r := AppointmentRepository{}
	_, err := r.Update(1, WithID(2))

	assert.Equal(t, err.Error(), "Update: Appointment not found")
}

func TestAppointmentRepository_All(t *testing.T) {
	var root []*Appointment
	date := time.Now()
	r := AppointmentRepository{}
	root = append(root, r.Create(WithID(1), WithCreatedAt(date), WithUpdatedAt(date)))
	root = append(root, r.Create(WithID(2), WithCreatedAt(date), WithUpdatedAt(date)))
	root = append(root, r.Create(WithID(3), WithCreatedAt(date), WithUpdatedAt(date)))
	root = append(root, r.Create(WithID(4), WithCreatedAt(date), WithUpdatedAt(date)))

	assert.Equal(t, r.All(), root)
}

func TestAppointmentRepository_First(t *testing.T) {
	date := time.Now()
	r := AppointmentRepository{}
	r.Create(WithID(2), WithCreatedAt(date), WithUpdatedAt(date))
	r.Create(WithID(4), WithCreatedAt(date), WithUpdatedAt(date))
	r.Create(WithID(3), WithCreatedAt(date), WithUpdatedAt(date))

	first := r.Create(WithID(1), WithCreatedAt(date), WithUpdatedAt(date))

	assert.Equal(t, r.First(), first)
}

func TestAppointmentRepository_Last(t *testing.T) {
	date := time.Now()
	r := AppointmentRepository{}
	r.Create(WithID(1), WithCreatedAt(date), WithUpdatedAt(date))
	r.Create(WithID(4), WithCreatedAt(date), WithUpdatedAt(date))
	r.Create(WithID(3), WithCreatedAt(date), WithUpdatedAt(date))

	last := r.Create(WithID(5), WithCreatedAt(date), WithUpdatedAt(date))

	assert.Equal(t, r.Last(), last)
}

func TestAppointmentRepository_Clear(t *testing.T) {
	date := time.Now()
	r := AppointmentRepository{}
	r.Create(WithID(1), WithCreatedAt(date), WithUpdatedAt(date))
	r.Create(WithID(4), WithCreatedAt(date), WithUpdatedAt(date))
	r.Create(WithID(3), WithCreatedAt(date), WithUpdatedAt(date))
	r.Create(WithID(5), WithCreatedAt(date), WithUpdatedAt(date))

	r.Clear()

	assert.Equal(t, r.Len(), 0)
}

func TestAppointmentRepository_Delete(t *testing.T) {
	date := time.Now()
	r := AppointmentRepository{}
	r.Create(WithID(2), WithCreatedAt(date), WithUpdatedAt(date))
	r.Create(WithID(4), WithCreatedAt(date), WithUpdatedAt(date))
	r.Create(WithID(3), WithCreatedAt(date), WithUpdatedAt(date))

	r.Delete(2)

	assert.Equal(t, r.Len(), 2)
}

func TestAppointmentRepository_Delete_NotFound(t *testing.T) {
	date := time.Now()
	r := AppointmentRepository{}
	r.Create(WithID(2), WithCreatedAt(date), WithUpdatedAt(date))
	r.Create(WithID(4), WithCreatedAt(date), WithUpdatedAt(date))
	r.Create(WithID(3), WithCreatedAt(date), WithUpdatedAt(date))

	_, err := r.Delete(1)

	assert.Equal(t, err.Error(), "Delete: Appointment not found")
}
