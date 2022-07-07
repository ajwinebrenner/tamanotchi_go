package resolvers

import (
	"context"
	"errors"
	"fmt"

	"github.com/ajwinebrenner/tamanotchi_go/api/internal/constants/mood"
	"github.com/ajwinebrenner/tamanotchi_go/api/internal/database"
	"github.com/ajwinebrenner/tamanotchi_go/api/internal/models"
	"gorm.io/gorm"
)

func NewPetMutation(db *gorm.DB) PetMutation {
	return PetMutation{db: db}
}

type PetMutation struct {
	db *gorm.DB
}

type createPetArgs struct {
	Name    string
	Species string
}

func (p *PetQuery) CreatePet(ctx context.Context, args createPetArgs) (*Pet, error) {
	pet := models.Pet{
		Name:      args.Name,
		Happiness: 1,
		Energy:    10,
		Mood:      mood.IDLE,
		HouseName: database.STARTING_HOUSE,
	}
	err := p.db.Transaction(func(tx *gorm.DB) error {
		species := models.Species{}
		err := tx.Where("name = ?", args.Species).First(&species).WithContext(ctx).Error
		if err != nil {
			return err
		}
		if species.Stage != 1 {
			return errors.New("incorrect species for new pet")
		}
		pet.Species = &species
		return tx.Create(&pet).WithContext(ctx).Error
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create new pet: %v", err)
	}
	return &Pet{pet: &pet}, nil
}
