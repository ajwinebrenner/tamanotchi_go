package models

type Species struct {
	Name          string `gorm:"PrimaryKey"`
	Stage         int
	MaxExp        int
	FaveFoodName  string
	FaveFood      *Food `gorm:"foreignKey:FaveFoodName"`
	EvolutionName string
	Evolution     *Species `gorm:"foreignKey:EvolutionName"`
}
