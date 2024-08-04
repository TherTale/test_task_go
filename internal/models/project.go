package models

import (
	"contact-center-system/pkg"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

type ProjectType string

const (
	Incoming     ProjectType = "входящий"
	Outgoing     ProjectType = "исходящий"
	AutoInformer ProjectType = "автоинформатор"
)

// Project представляет сущность проекта в системе.
type Project struct {
	ID   uuid.UUID   `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	Name string      `bun:",notnull" json:"name"`
	Type ProjectType `bun:",notnull" json:"type"`

	Operators []*Operator `bun:"m2m:project_assignments,join:Project=Operator" json:"operators,omitempty"`
}

func (p *Project) UnmarshalJSON(data []byte) error {
	type Alias Project
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	var projectTypes = []ProjectType{
		Incoming,
		Outgoing,
		AutoInformer,
	}
	if !pkg.Contains(projectTypes, aux.Type) {
		return fmt.Errorf("непривальный тип проекта")
	}
	return nil
}

// ProjectAssignment связывает операторов с проектами.
type ProjectAssignment struct {
	ID        uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	ProjectID uuid.UUID `bun:",notnull"`
	Project   *Project  `bun:"rel:belongs-to,join:project_id=id"`

	OperatorID uuid.UUID `bun:",notnull"`
	Operator   *Operator `bun:"rel:belongs-to,join:operator_id=id"`
}

func (p *Project) GetUpdateFields() []string {
	var updateFields = []string{
		"name",
		"type",
	}
	return updateFields
}
