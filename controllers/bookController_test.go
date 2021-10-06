package controllers

import (
	"github.com/aysf/gojwt/config"
	"github.com/aysf/gojwt/models"
)

func InsertMockBookData() error {
	mock_book1 := models.Book{
		Title:     "seven habit",
		Author:    "Stephen R. Covey",
		Publisher: "Free Press",
	}
	mock_book2 := models.Book{
		Title:     "Rich Dad Poor Dad",
		Author:    "Robert Kiyosaki",
		Publisher: "Gramed",
	}

	if err := config.DB.Save(&mock_book1).Error; err != nil {
		return err
	}
	if err := config.DB.Save(&mock_book2).Error; err != nil {
		return err
	}

	return nil
}
