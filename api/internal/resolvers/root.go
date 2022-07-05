package resolvers

import "gorm.io/gorm"

func NewRootResolver(db *gorm.DB) *RootResolver {
	return &RootResolver{
		FoodResolver:    NewFoodResolver(db),
		HouseResolver:   NewHouseResolver(db),
		SpeciesResolver: NewSpeciesResolver(db),
		PetResolver:     NewPetResolver(db),
	}
}

type RootResolver struct {
	*FoodResolver
	*HouseResolver
	*SpeciesResolver
	*PetResolver
}
