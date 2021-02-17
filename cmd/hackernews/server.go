package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/huaxk/hackernews/gqlgen"
	"github.com/huaxk/hackernews/internal/auth"
	"github.com/huaxk/hackernews/models"
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

	router.Handle("/", gqlgen.NewPlaygroundHandler("/query"))
	router.Handle("/query", gqlgen.NewHandler(db))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
