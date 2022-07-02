package resolvers

import "gorm.io/gorm"

func NewRootResolver(db *gorm.DB) *RootResolver {
	return &RootResolver{
		FoodResolver: NewFoodResolver(db),
	}
}

type RootResolver struct {
	*FoodResolver
}
