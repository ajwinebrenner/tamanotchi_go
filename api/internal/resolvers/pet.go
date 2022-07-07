package resolvers

import (
	"context"

	"github.com/ajwinebrenner/tamanotchi_go/api/internal/models"
	"gorm.io/gorm"
)

func NewPetResolver(db *gorm.DB) *PetResolver {
	return &PetResolver{
		NewPetQuery(db),
		NewPetMutation(db),
	}
}

type PetResolver struct {
	PetQuery
	PetMutation
}

type Pet struct {
	pet *models.Pet
}

func (p *Pet) ID(ctx context.Context) int32 {
	return p.pet.ID
}

func (p *Pet) Name(ctx context.Context) string {
	return p.pet.Name
}

func (p *Pet) Species(ctx context.Context) *Species {
	return &Species{species: p.pet.Species}
}

func (p *Pet) Happiness(ctx context.Context) int32 {
	return p.pet.Happiness
}

func (p *Pet) Energy(ctx context.Context) int32 {
	return p.pet.Energy
}

func (p *Pet) Mood(ctx context.Context) string {
	return p.pet.Mood.String()
}

func (p *Pet) Exp(ctx context.Context) int32 {
	return p.pet.Exp
}

func (p *Pet) Money(ctx context.Context) int32 {
	return p.pet.Money
}

func (p *Pet) House(ctx context.Context) *House {
	return &House{house: p.pet.House}
}
