package models

import "github.com/ajwinebrenner/tamanotchi_go/api/internal/constants/variety"

type Food struct {
	Name      string `gorm:"PrimaryKey"`
	Price     int
	Energy    int
	Happiness int
	Variety   variety.Variety
}
