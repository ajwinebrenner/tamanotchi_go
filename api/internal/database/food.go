package database

import (
	"github.com/ajwinebrenner/tamanotchi_go/api/internal/constants/variety"
	"github.com/ajwinebrenner/tamanotchi_go/api/internal/models"
	"gorm.io/gorm"
)

func InsertFood(db *gorm.DB) error {
	food := []*models.Food{
		{Name: "medicine", Price: 8, Energy: 0, Happiness: 0, Variety: variety.HEALING},
		{Name: "carrot", Price: 3, Energy: 2, Happiness: 0, Variety: variety.HEALTHY},
		{Name: "candy", Price: 6, Energy: 0, Happiness: 2, Variety: variety.UNHEALTHY},
		{Name: "pizza", Price: 7, Energy: 3, Happiness: 1, Variety: variety.UNHEALTHY},
		{Name: "fish", Price: 5, Energy: 4, Happiness: 0, Variety: variety.HEALTHY},
		{Name: "apple", Price: 5, Energy: 1, Happiness: 1, Variety: variety.HEALTHY},
	}
	return db.Create(&food).Error
}
