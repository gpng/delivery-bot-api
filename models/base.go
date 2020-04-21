package models

import (
	"log"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // required for postgres dbs
)

// Model struct to replace gorm Model in order to stop certain fields from being exported
type Model struct {
	ID        string     `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// ModelWithTime which exports create/updated dates to json
type ModelWithTime struct {
	ID        string     `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

// BeforeCreate callback for models to add uuid as id
func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()

	if err != nil {
		log.Printf("uuid NewV4 failed with error: %v", err)
		return err
	}

	scope.SetColumn("ID", id)
	return nil
}
