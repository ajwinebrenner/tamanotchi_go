package resolvers

import (
	"context"

	"github.com/ajwinebrenner/tamanotchi_go/api/internal/models"
	"gorm.io/gorm"
)

func NewFoodQuery(db *gorm.DB) FoodQuery {
	return FoodQuery{db: db}
}

type FoodQuery struct {
	db *gorm.DB
}

func (f *FoodQuery) Foods(ctx context.Context) ([]*Food, error) {
	var foods []*models.Food
	err := f.db.Find(&foods).WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	foodResolvers := make([]*Food, 0, len(foods))
	for _, food := range foods {
		foodResolvers = append(foodResolvers, &Food{food: food})
	}
	return foodResolvers, nil
}
