package models

import (
	"gorm.io/gorm"
)

type Group struct {
	gorm.Model
	GroupName   string
	Thermostats []Thermostat
}

type Thermostat struct {
	gorm.Model
	Name        string
	CurrentTemp int64
	TempDisplay bool
	GroupID     int64
}
