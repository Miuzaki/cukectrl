package models

import "github.com/google/uuid"

type Bot struct {
	ID    uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Token string    `json:"token" gorm:"type:varchar(255);not null;unique"`
}
