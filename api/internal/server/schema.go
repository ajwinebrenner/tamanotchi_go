package server

import (
	"embed"
	"log"

	"github.com/ajwinebrenner/tamanotchi_go/api/internal/resolvers"
	"github.com/graph-gophers/graphql-go"
	"gorm.io/gorm"
)

//go:embed graphql/schema.graphql
var f embed.FS

func CreateGraphQLSchema(db *gorm.DB) (*graphql.Schema, error) {
	bstr, err := f.ReadFile("graphql/schema.graphql")
	if err != nil {
		log.Fatalf("error reading schema %v", err)
	}
	schemaStr := string(bstr)
	opts := []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	resolver := resolvers.NewRootResolver(db)
	return graphql.ParseSchema(schemaStr, resolver, opts...)
}
