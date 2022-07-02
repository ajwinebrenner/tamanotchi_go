package resolvers

import (
	"context"

	"github.com/ajwinebrenner/tamanotchi_go/api/internal/models"
	"gorm.io/gorm"
)

func NewHouseQuery(db *gorm.DB) HouseQuery {
	return HouseQuery{db: db}
}

type HouseQuery struct {
	db *gorm.DB
}

func (h *HouseQuery) Houses(ctx context.Context) ([]*House, error) {
	houses := []*models.House{}
	err := h.db.Preload("Upgrade").Find(&houses).WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	houseResolvers := make([]*House, 0, len(houses))
	for _, house := range houses {
		houseResolvers = append(houseResolvers, &House{house: house})
	}
	return houseResolvers, nil
}

type houseArgs struct {
	Name string
}

func (h *HouseQuery) House(ctx context.Context, args houseArgs) (*House, error) {
	house := &models.House{}
	err := h.db.Preload("Upgrade").First(house, "name = ?", args.Name).WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	return &House{house: house}, nil
}
