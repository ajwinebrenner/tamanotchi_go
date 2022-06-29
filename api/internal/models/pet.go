package models

import "github.com/ajwinebrenner/tamanotchi_go/api/internal/constants/mood"

type Pet struct {
	ID          int `gorm:"PrimaryKey"`
	SpeciesName string
	Species     *Species `gorm:"foreignKey:SpeciesName"`
	Happiness   int
	Energy      int
	Mood        mood.Mood
	Exp         int
	Money       int
	HouseName   string
	House       *House `gorm:"foreignKey:HouseName"`
	// TODO: Filth int
}
