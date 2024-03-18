package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	TeamId uuid.UUID
	Name   string `json:"name"`
	Users  []User
}
