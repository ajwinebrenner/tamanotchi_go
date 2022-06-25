package database

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ReadyDB(dbName string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error getting connection to %v : %v", dbName, err)
	}

	m := db.Migrator()
	models := []interface{}{
		//models
	}

	for _, model := range models {
		if !m.HasTable(model) {
			err = m.AutoMigrate(model)
			if err != nil {
				return nil, fmt.Errorf("error migrating schema: %v", err)
			}
		}
	}

	return db, nil
}