package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct{
	ID 			uint `json:"id" gorm:"primaryKey"`
	Name 		string `json:"name"`
	Email 		string `json:"email"`
	City 		string `json:"city"`
	Province 	string `json:"province"`
	Created_at 	time.Time `json:"created_at,omitempty" db:"created_at"`
	Updated_at 	time.Time `json:"updated_at,omitempty" db:"updated_at"`
	Deleted_at 	gorm.DeletedAt `json:"deleted_at" gorm:"index"`
} 