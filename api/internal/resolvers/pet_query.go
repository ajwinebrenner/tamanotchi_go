package resolvers

import (
	"context"

	"github.com/ajwinebrenner/tamanotchi_go/api/internal/models"
	"gorm.io/gorm"
)

func NewPetQuery(db *gorm.DB) PetQuery {
	return PetQuery{db: db}
}

type PetQuery struct {
	db *gorm.DB
}

func (p *PetQuery) Pets(ctx context.Context) ([]*Pet, error) {
	pets := []*models.Pet{}
	err := p.db.Preload("Species").Preload("House").Find(&pets).WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	petResolvers := make([]*Pet, 0, len(pets))
	for _, pet := range pets {
		petResolvers = append(petResolvers, &Pet{pet: pet})
	}
	return petResolvers, nil
}

type petArgs struct {
	ID int32
}

func (p *PetQuery) Pet(ctx context.Context, args petArgs) (*Pet, error) {
	pet := &models.Pet{}
	err := p.db.Preload("Species").Preload("House").First(pet, "id = ?", args.ID).WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	return &Pet{pet: pet}, nil
}
