package models

import (
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	Name           string `gorm:"type:varchar(255); not null"`
	Email          string `gorm:"type:varchar(255); unique; not null"`
	Password       string `gorm:"type:varchar(255); not null" json:"-"`
	City           string `gorm:"type:varchar(255); not null"`
	Est            int    `gorm:"not null"`
	Facebook       string `gorm:"type:varchar(255)"`
	Instagram      string `gorm:"type:varchar(255)"`
	DisplayPicture string `gorm:"type:varchar(255)"`
}
