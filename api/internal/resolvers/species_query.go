package resolvers

import (
	"context"

	"github.com/ajwinebrenner/tamanotchi_go/api/internal/models"
	"gorm.io/gorm"
)

func NewSpeciesQuery(db *gorm.DB) SpeciesQuery {
	return SpeciesQuery{db: db}
}

type SpeciesQuery struct {
	db *gorm.DB
}

func (s *SpeciesQuery) AllSpecies(ctx context.Context) ([]*Species, error) {
	allSpecies := []*models.Species{}
	err := s.db.Preload("Evolution").Preload("FaveFood").
		Find(&allSpecies).WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	speciesResolvers := make([]*Species, 0, len(allSpecies))
	for _, species := range allSpecies {
		speciesResolvers = append(speciesResolvers, &Species{species: species})
	}
	return speciesResolvers, nil
}

type speciesArgs struct {
	Name string
}

func (s *SpeciesQuery) Species(ctx context.Context, args speciesArgs) (*Species, error) {
	species := &models.Species{}
	err := s.db.Preload("Evolution").Preload("FaveFood").
		First(species, "name = ?", args.Name).WithContext(ctx).Error
	if err != nil {
		return nil, err
	}
	return &Species{species: species}, nil
}

// TODO get babies (for hatching new pet)
