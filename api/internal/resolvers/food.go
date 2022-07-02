package resolvers

import (
	"context"

	"github.com/ajwinebrenner/tamanotchi_go/api/internal/models"
	"gorm.io/gorm"
)

func NewFoodResolver(db *gorm.DB) *FoodResolver {
	return &FoodResolver{
		NewFoodQuery(db),
	}
}

type FoodResolver struct {
	FoodQuery
}

type Food struct {
	food *models.Food
}

func (f *Food) Name(ctx context.Context) string {
	return f.food.Name
}
func (f *Food) Price(ctx context.Context) int32 {
	return f.food.Price
}
func (f *Food) Energy(ctx context.Context) int32 {
	return f.food.Energy
}
func (f *Food) Happiness(ctx context.Context) int32 {
	return f.food.Happiness
}
func (f *Food) Variety(ctx context.Context) int32 {
	return int32(f.food.Variety)
}
