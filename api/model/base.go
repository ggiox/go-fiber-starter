package model

import (
	"time"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

// User struct
// To use uuid_generate_v4() in postgres, you need the uuid-ossp extension
// If it does not exist, execute this command: CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
type Base struct {
	UUID      uuid.UUID `gorm:"primary_key; unique; type:uuid; column:uuid; default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
