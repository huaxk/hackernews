package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"github.com/huaxk/hackernews/graph/handler"
	"github.com/huaxk/hackernews/internal/auth"
	"github.com/huaxk/hackernews/repo/gorm"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := gorm.Open("host=localhost port=5432 user=postgres password=postgres dbname=hackernews_development sslmode=disable TimeZone=Asia/Shanghai")
	if err != nil {
		log.Fatal(err)
	}
	// db.AutoMigrate(&models.User{}, &models.Link{})

	router := chi.NewRouter()
	router.Use(auth.Middleware(db))

	router.Handle("/", handler.NewPlaygroundHandler("/query"))
	router.Handle("/query", handler.NewHandler(gorm.NewRepository(db)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
