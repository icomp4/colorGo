package controllers

import (
	"github.com/zipqt/goScrape/database"
	"github.com/zipqt/goScrape/models"
)

func Save(p *models.Page) error{
	if err := database.DB.Create(&p).Error; err != nil{
		return err
	}
	return nil
}

func Get(id string) (models.Page, error){
	var page models.Page
	if err := database.DB.First(&page,id).Error; err != nil {
		return models.Page{}, err
	}
	return page, nil
}