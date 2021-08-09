package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mdhishaamakhtar/learnFiber/api/handlers"
	"github.com/mdhishaamakhtar/learnFiber/pkg/models"
	"github.com/mdhishaamakhtar/learnFiber/pkg/post"
	"github.com/mdhishaamakhtar/learnFiber/pkg/user"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func ConnectDB() (*gorm.DB, error) {
	url, ok := viper.Get("DATABASE_URL").(string)
	if !ok {
		log.Panicln(fmt.Errorf("could not find database url"))
	}
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	return db, err
}

func getPort() string {
	port, ok := viper.Get("PORT").(string)
	if !ok {
		fmt.Println("No PORT provided. Defaulting to 3000")
		return ":3000"
	}
	return fmt.Sprintf(":%v", port)
}

func main() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Panicln(fmt.Errorf("fatal error config file: %s", err))
	}

	db, err := ConnectDB()
	if err != nil {
		log.Panicln(fmt.Errorf("fatal error in connecting to db: %v", err))
	}
	log.Println("Connected to Database")

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	err = db.AutoMigrate(
		&models.User{},
		&models.Post{},
	)
	if err != nil {
		log.Panicln(fmt.Errorf("error migrating models: %v", err))
	}
	log.Println("Database Migration Done")

	// Setup Services, Repos and Handlers
	userRepo := user.NewRepo(db)
	userSvc := user.NewService(userRepo)
	handlers.MakeUserHandler(app, userSvc)

	postRepo := post.NewRepo(db)
	postSvc := post.NewService(postRepo)
	handlers.MakePostHandler(app, postSvc)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"Status": "Ok",
		})
	})

	log.Fatal(app.Listen(getPort()))
}
