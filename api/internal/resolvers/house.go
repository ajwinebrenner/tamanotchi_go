package resolvers

import (
	"context"

	"github.com/ajwinebrenner/tamanotchi_go/api/internal/models"
	"gorm.io/gorm"
)

func NewHouseResolver(db *gorm.DB) *HouseResolver {
	return &HouseResolver{
		NewHouseQuery(db),
	}
}

type HouseResolver struct {
	HouseQuery
}

type House struct {
	house *models.House
}

func (f *House) Name(ctx context.Context) string {
	return f.house.Name
}
func (f *House) Price(ctx context.Context) int32 {
	return f.house.Price
}
func (f *House) Size(ctx context.Context) int32 {
	return f.house.Size
}
func (f *House) Upgrade(ctx context.Context) *House {
	upgrade := f.house.Upgrade
	if upgrade == nil {
		return nil
	}
	return &House{house: upgrade}
}
