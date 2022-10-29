package models

import (
	"gorm.io/gorm"
)

type DiscSkills struct {
	gorm.Model
	UserID              uint   `gorm:"index; not null"`
	PrimaryRole         int    `gorm:"unsigned; not null"`
	Throwing            int    `gorm:"unsigned; not null"`
	Catching            int    `gorm:"unsigned; not null"`
	OffensiveStrategies int    `gorm:"unsigned; not null"`
	DefensiveStrategies int    `gorm:"unsigned; not null"`
	Public              string `gorm:"default:false"`
}
