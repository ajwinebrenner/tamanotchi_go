package models

type Food struct {
	Name        string `gorm:"PrimaryKey"`
	Price       int
	Energy      int
	Happiness   int
	isUnhealthy bool
	isHealing   bool
}
