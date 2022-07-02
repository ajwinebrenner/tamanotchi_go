package models

import "github.com/ajwinebrenner/tamanotchi_go/api/internal/constants/variety"

type Food struct {
	Name      string `gorm:"PrimaryKey"`
	Price     int32
	Energy    int32
	Happiness int32
	Variety   variety.Variety
}
