package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	// Use uuid.UUID as the primary key
	Id        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	Title     string    `json:"title"`
	Content   *string   `json:"content,omitempty"`               // Optional field, pointer to string
	Done      bool      `json:"done" gorm:"default:false"`       // Default to false                        // Changed type to bool for true/false
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"` // Auto-set on creation
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"` // Auto-set on update
}

type Tasks struct {
	Tasks []Task `json:"tasks"`
}

func (task *Task) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	task.Id = uuid.New()
	return
}
