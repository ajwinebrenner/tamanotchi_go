package resolvers

import (
	"context"

	"github.com/ajwinebrenner/tamanotchi_go/api/internal/models"
	"gorm.io/gorm"
)

func NewSpeciesResolver(db *gorm.DB) *SpeciesResolver {
	return &SpeciesResolver{
		NewSpeciesQuery(db),
	}
}

type SpeciesResolver struct {
	SpeciesQuery
}

type Species struct {
	species *models.Species
}

func (s *Species) Name(ctx context.Context) string {
	return s.species.Name
}
func (s *Species) Stage(ctx context.Context) int32 {
	return s.species.Stage
}
func (s *Species) MaxExp(ctx context.Context) int32 {
	return s.species.MaxExp
}
func (s *Species) FaveFood(ctx context.Context) *Food {
	fave := s.species.FaveFood
	if fave == nil {
		return nil
	}
	return &Food{food: fave}
}
func (s *Species) Evolution(ctx context.Context) *Species {
	upgrade := s.species.Evolution
	if upgrade == nil {
		return nil
	}
	return &Species{species: upgrade}
}
