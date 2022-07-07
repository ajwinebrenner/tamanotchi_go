package database

import (
	"github.com/ajwinebrenner/tamanotchi_go/api/internal/models"
	"gorm.io/gorm"
)

const STARTING_HOUSE = "hut"

func insertHouses(db *gorm.DB) error {
	houses := []*models.House{
		{Name: "mansion", Price: 30, Size: 3, UpgradeName: ""},
		{Name: "house", Price: 20, Size: 2, UpgradeName: "mansion"},
		{Name: STARTING_HOUSE, Price: 10, Size: 1, UpgradeName: "house"},
	}
	return db.Create(&houses).Error
}
