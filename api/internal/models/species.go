package models

type Species struct {
	Name          string `gorm:"PrimaryKey"`
	Stage         int32
	MaxExp        int32
	FaveFoodName  string
	FaveFood      *Food `gorm:"foreignKey:FaveFoodName"`
	EvolutionName string
	Evolution     *Species `gorm:"foreignKey:EvolutionName"`
}
