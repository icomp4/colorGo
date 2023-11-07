package models

import "gorm.io/gorm"

type Page struct {
	gorm.Model
	URL string
}

// Making sure the url isnt blank 
func (p *Page) isValid() bool {
	return p.URL != ""
}