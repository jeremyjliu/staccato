package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"time"
)

type Staccato struct {
	ID        uuid.UUID    `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	Open      bool         `json:"open" db:"open"`
	TimeOpen  time.Time    `json:"time_open" db:"time_open"`
	TimeClose time.Time    `json:"time_close" db:"time_close"`
	Link      nulls.String `json:"link" db:"link"`
}

// String is not required by pop and may be deleted
func (s Staccato) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Staccatos is not required by pop and may be deleted
type Staccatos []Staccato

// String is not required by pop and may be deleted
func (s Staccatos) String() string {
	js, _ := json.Marshal(s)
	return string(js)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (s *Staccato) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.TimeIsPresent{Field: s.TimeOpen, Name: "TimeOpen"},
		&validators.TimeIsPresent{Field: s.TimeClose, Name: "TimeClose"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (s *Staccato) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (s *Staccato) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
