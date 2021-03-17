package models

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	ID       uint
	Rating   string
	Title    string
	Duration int
	Year     int
}
