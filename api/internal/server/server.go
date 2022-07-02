package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ajwinebrenner/tamanotchi_go/api/internal/config"
	"github.com/ajwinebrenner/tamanotchi_go/api/internal/database"
)

func Start(c *config.Config) {
	db, err := database.ReadyDB(c.DBName)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}
	//graphql schema
	schema, err := CreateGraphQLSchema(db)
	if err != nil {
		log.Fatalf("Failed to create schema %v", err)
	}
	//routes
	handlers := Routes(db, schema, c.CORSorigins, c.CORSmethods)
	//listen and serve
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", c.Host, c.Port),
		Handler: handlers,
	}
	server.ListenAndServe()
}
