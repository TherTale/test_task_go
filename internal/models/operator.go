package models

import (
	"github.com/google/uuid"
)

// Operator представляет сущность оператора в системе.
type Operator struct {
	ID          uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	FirstName   string    `bun:",notnull" json:"firstName"`
	LastName    string    `bun:",notnull" json:"lastName"`
	MiddleName  string    `bun:",notnull" json:"middleName"`
	City        string    `bun:",notnull" json:"city"`
	PhoneNumber string    `bun:",notnull,unique" json:"phoneNumber"`
	Email       string    `bun:",notnull,unique" json:"email"`
	Password    string    `bun:",notnull" json:"password"`
}
