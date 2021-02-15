package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/huaxk/hackernews/graph"
	"github.com/huaxk/hackernews/graph/generated"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/huaxk/hackernews/internal/auth"
	"github.com/huaxk/hackernews/internal/models"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres password=postgres dbname=hackernews_development sslmode=disable TimeZone=Asia/Shanghai"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.User{}, &models.Link{})

	router := chi.NewRouter()
	router.Use(auth.Middleware(db))

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
