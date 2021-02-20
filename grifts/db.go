package grifts

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/huaxk/hackernews/models"
	"github.com/huaxk/hackernews/repo/gorm"
	. "github.com/markbates/grift/grift"
	"github.com/pioz/faker"
)

var _ = Namespace("db", func() {
	envy.Load("database.env")
	url, err := envy.MustGet("DEV_DATABASE_URL")
	if err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(url)
	if err != nil {
		log.Fatal(err)
	}

	Desc("migrate", "Migrates the databases")
	Set("migrate", func(c *Context) error {

		err = db.AutoMigrate(models.DbModels...)
		if err == nil {
			log.Println("db:migrate: created database schema")
		}
		return err
	})

	Desc("seed", "Seeds faker data")
	Set("seed", func(c *Context) error {
		log.Println(faker.Username())
		return nil
	})
})
