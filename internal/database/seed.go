package database

import (
	"karhub-beer-api/internal/models"
)

func SeedBeers() {
	var beers = []models.Beer{
		{ID: 1, Name: "Weissbier", MinTemp: -1, MaxTemp: 3},
		{ID: 2, Name: "Pilsens", MinTemp: -2, MaxTemp: 4},
		{ID: 3, Name: "Weizenbier", MinTemp: -4, MaxTemp: 6},
		{ID: 4, Name: "Red ale", MinTemp: -5, MaxTemp: 5},
		{ID: 5, Name: "India pale ale", MinTemp: -6, MaxTemp: 7},
		{ID: 6, Name: "IPA", MinTemp: -7, MaxTemp: 10},
		{ID: 7, Name: "Dunkel", MinTemp: -8, MaxTemp: 2},
		{ID: 8, Name: "Imperial Stouts", MinTemp: -10, MaxTemp: 13},
		{ID: 9, Name: "Brown ale", MinTemp: 0, MaxTemp: 14},
	}

	for _, beer := range beers {
		var existing models.Beer
		result := DB.First(&existing, beer.ID)
		if result.Error != nil {
			DB.Create(&beer)
		}
	}
}
