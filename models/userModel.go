package models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName      string    `gorm:"type:varchar(255); not null"`
	LastName       string    `gorm:"type:varchar(255); not null"`
	Password       string    `gorm:"type:varchar(255); not null" json:"-"`
	Height         int       `gorm:"not null"`
	Gender         string    `gorm:"type:varchar(255); not null"`
	Email          string    `gorm:"type:varchar(255); unique; not null"`
	Birthday       time.Time `gorm:"type:date; not null"`
	CommonName     string    `gorm:"type:varchar(255)"`
	DisplayPicture string    `gorm:"type:varchar(255)"`
	OtherInfo      datatypes.JSON
	DiscSkills     DiscSkills
}

func ParseBirthday(birthday string) (time.Time, error) {
	parsedBirthday, err := time.Parse("2006-01-02", birthday)
	return parsedBirthday, err
}
