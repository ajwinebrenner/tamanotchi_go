package database

import (
	"github.com/ajwinebrenner/tamanotchi_go/api/internal/models"
	"gorm.io/gorm"
)

func insertSpecies(db *gorm.DB) error {
	species := []*models.Species{
		{Name: "impeckable", Stage: 3, MaxExp: 40, FaveFoodName: "fish", EvolutionName: ""},
		{Name: "nugget", Stage: 2, MaxExp: 25, FaveFoodName: "pizza", EvolutionName: "impeckable"},
		{Name: "chickpea", Stage: 1, MaxExp: 10, FaveFoodName: "carrot", EvolutionName: "nugget"},
		{Name: "firecracker", Stage: 3, MaxExp: 45, FaveFoodName: "pizza", EvolutionName: ""},
		{Name: "big eatie", Stage: 2, MaxExp: 30, FaveFoodName: "apple", EvolutionName: "firecracker"},
		{Name: "rexy", Stage: 1, MaxExp: 15, FaveFoodName: "fish", EvolutionName: "big eatie"},
	}
	return db.Create(&species).Error
}
