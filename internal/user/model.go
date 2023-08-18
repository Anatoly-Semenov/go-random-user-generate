package user

import (
	"gorm.io/gorm"
	"time"
)

type Sex string

const (
	MALE   Sex = "male"
	FEMALE Sex = "female"
)

type Model struct {
	gorm.Model

	Id        int       `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string    `gorm:"size:255;null" json:"firstName"`
	LastName  string    `gorm:"size:255;null" json:"lastName"`
	Dob       string    `gorm:"size:255;null" json:"dob"`
	Age       int       `json:"age"`
	Sex       Sex       `json:"sex"`
	Avatar    string    `gorm:"size:255;null" json:"avatar"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (Model) TableName() string {
	return "users"
}
