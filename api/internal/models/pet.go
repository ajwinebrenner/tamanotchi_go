package models

import "github.com/ajwinebrenner/tamanotchi_go/api/internal/constants/mood"

type Pet struct {
	ID          int32 `gorm:"PrimaryKey"`
	SpeciesName string
	Species     *Species `gorm:"foreignKey:SpeciesName"`
	Happiness   int32
	Energy      int32
	Mood        mood.Mood
	Exp         int32
	Money       int32
	HouseName   string
	House       *House `gorm:"foreignKey:HouseName"`
	// TODO: Filth int
}
