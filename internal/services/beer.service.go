package services

import (
	"errors"
	"karhub-beer-api/internal/database"
	"karhub-beer-api/internal/models"
)

func GetAllBeers() []models.Beer {
	var beers []models.Beer
	database.DB.Find(&beers)
	return beers
}

func SaveBeer(b models.Beer) models.Beer {
	database.DB.Create(&b)
	return b
}

func UpdateBeer(id int, updated models.Beer) (models.Beer, error) {
	var beer models.Beer
	result := database.DB.First(&beer, id)
	if result.Error != nil {
		return models.Beer{}, errors.New("beer not found")
	}
	updated.ID = beer.ID
	database.DB.Model(&beer).Updates(updated)
	return updated, nil
}

func DeleteBeer(id int) error {
	var beer models.Beer
	result := database.DB.First(&beer, id)
	if result.Error != nil {
		return errors.New("beer not found")
	}
	database.DB.Delete(&beer)
	return nil
}
