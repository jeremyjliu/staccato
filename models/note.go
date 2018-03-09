package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"time"
)

type Note struct {
	ID         uuid.UUID     `json:"id" db:"id"`
	CreatedAt  time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at" db:"updated_at"`
	StaccatoID uuid.UUID     `json:"staccato_id" db:"staccato_id"`
	UserID     uuid.UUID     `json:"user_id" db:"user_id"`
	Songs      slices.String `json:"songs" db:"songs"`
}

// String is not required by pop and may be deleted
func (n Note) String() string {
	jn, _ := json.Marshal(n)
	return string(jn)
}

// Notes is not required by pop and may be deleted
type Notes []Note

// String is not required by pop and may be deleted
func (n Notes) String() string {
	jn, _ := json.Marshal(n)
	return string(jn)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (n *Note) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (n *Note) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (n *Note) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
