package grifts

import (
	"log"

	"github.com/bxcodec/faker/v3"
	"github.com/gobuffalo/envy"
	"github.com/huaxk/hackernews/models"
	"github.com/huaxk/hackernews/repo/gorm"
	. "github.com/markbates/grift/grift"
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
		user := models.User{
			Name:     faker.Username(),
			Password: faker.Password(),
		}
		db.Save(&user)

		link := models.Link{
			Title:   faker.Sentence(),
			Address: faker.Word(),
			UserID:  user.ID,
		}
		db.Save(&link)
		return nil
	})
})

// var password = "12345"
// var user = userFactory()

// func userFactory() *models.User {
// 	hashedPassword, err := models.HashPassword(password)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	user := &models.User{
// 		Password: hashedPassword,
// 	}
// 	faker.Build(user)
// 	return user
// }

// func linkFactory() *models.Link {
// 	link := &models.Link{
// 		User: *user,
// 	}
// 	faker.Build(link)
// 	return link
// }
