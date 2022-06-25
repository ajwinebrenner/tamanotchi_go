package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ajwinebrenner/tamanotchi_go/api/internal/config"
	"github.com/ajwinebrenner/tamanotchi_go/api/internal/database"
	"github.com/graph-gophers/graphql-go"
)

func Start(c *config.Config) {
	db, err := database.ReadyDB(c.DBName)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	//graphql schema

	//routes
	handlers := Routes(db, &graphql.Schema{}, c.CORSorigins, c.CORSmethods)

	//listen and serve
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", c.Host, c.Port),
		Handler: handlers,
	}
	server.ListenAndServe()
}
