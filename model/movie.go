package model

import "gorm.io/gorm"

type Movie struct {
	gorm.Model
	Title       string
	Slug        string
	Description string
	Duration    int
	Image       string
}
