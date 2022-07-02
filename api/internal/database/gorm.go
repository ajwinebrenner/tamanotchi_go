package database

import (
	"fmt"

	"github.com/ajwinebrenner/tamanotchi_go/api/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ReadyDB(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error getting connection to %v : %v", dbName, err)
	}

	m := db.Migrator()
	modelStructs := []interface{}{
		models.Food{},
		models.Species{},
		models.House{},
		models.Pet{},
	}

	for _, model := range modelStructs {
		err = m.AutoMigrate(model)
		if err != nil {
			return nil, fmt.Errorf("error migrating schema: %v", err)
		}
	}
	var count int64
	db.Model(models.Species{}).Count(&count)
	if count == 0 {
		err = insertFoods(db)
		if err != nil {
			return nil, fmt.Errorf("error populating foods table: %v", err)
		}
		err = insertHouses(db)
		if err != nil {
			return nil, fmt.Errorf("error populating houses table: %v", err)
		}
		err = insertSpecies(db)
		if err != nil {
			return nil, fmt.Errorf("error populating species table: %v", err)
		}
	}
	return db, nil
}
