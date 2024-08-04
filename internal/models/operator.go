package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"regexp"
)

// Operator представляет сущность оператора в системе.
type Operator struct {
	ID          uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()" json:"id"`
	FirstName   string    `bun:",notnull" json:"firstName"`
	LastName    string    `bun:",notnull" json:"lastName"`
	MiddleName  string    `bun:",notnull" json:"middleName"`
	City        string    `bun:",notnull" json:"city"`
	PhoneNumber string    `bun:",notnull,unique" json:"phoneNumber"`
	Email       string    `bun:",notnull,unique" json:"email"`
	Password    string    `bun:",notnull" json:"password"`

	Projects []*Project `bun:"m2m:project_assignments,join:Operator=Project" json:"projects,omitempty"`
}

func (o *Operator) UnmarshalJSON(data []byte) error {
	type Alias Operator
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(o),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if !isValidName(aux.FirstName) || !isValidName(aux.LastName) || !isValidName(aux.MiddleName) {
		return fmt.Errorf("invalid name format")
	}
	if aux.City == "" {
		return fmt.Errorf("city cannot be empty")
	}
	if !isValidPhoneNumber(aux.PhoneNumber) {
		return fmt.Errorf("invalid phone number format")
	}
	if !isValidEmail(aux.Email) {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

func (o *Operator) GetUpdateFields() []string {
	var updateFields = []string{
		"city",
		"phone_number",
		"first_name",
		"last_name",
		"email",
		"middle_name",
	}
	return updateFields
}

// Валидация формата имени
func isValidName(name string) bool {
	re := regexp.MustCompile(`^[a-zA-Zа-яА-ЯёЁ\s-]+$`)
	return re.MatchString(name) && len(name) > 0 && len(name) <= 255
}

// Валидация формата номера телефона (8**********)
func isValidPhoneNumber(phone string) bool {
	re := regexp.MustCompile(`^8\d{10}$`)
	return re.MatchString(phone)
}

// Валидация формата Email
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}
