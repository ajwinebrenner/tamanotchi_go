package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/rs/cors"
	"gorm.io/gorm"
)

func Routes(db *gorm.DB, schema *graphql.Schema, origins []string, methods []string) http.Handler {
	router := chi.NewRouter()

	router.Use(
		cors.New(cors.Options{
			AllowedOrigins: origins,
			AllowedMethods: methods,
		}).Handler,
	)
	router.Method(http.MethodPost, "/query", &relay.Handler{Schema: schema})
	router.Get("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		fmt.Fprint(w, "<h1>World!</h1>")
	})
	return router
}
