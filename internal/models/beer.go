package models

type Beer struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	Name    string  `json:"name"`
	MinTemp float64 `json:"minTemp"`
	MaxTemp float64 `json:"maxTemp"`
}
