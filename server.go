package main

import (
	"books-backend/handler"
	"books-backend/repository"
	"books-backend/router"
	"books-backend/service"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupTimeZone() error {
	location, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		return err
	}
	time.Local = location
	return nil
}

func fiberConfig() fiber.Config {
	return fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Books",
	}
}

func corsConfig() cors.Config {
	return cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}
}
func setupFiber() error {
	app := fiber.New(fiberConfig())
	app.Use(cors.New(corsConfig()))
	redis := setupRedis()
	db, err := setupDatabase()
	if err != nil {
		return err
	}

	reviewRepository := repository.NewReviewRepository(db)

	reviewService := service.NewReviewController(reviewRepository)

	reviewHandler := handler.NewReviewHandler(reviewService, redis)

	router.New(app, reviewHandler)
	err = app.Listen(":" + os.Getenv("PORT"))
	return err
}

func setupDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_ADDRESS"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func setupRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("localhost:6379", os.Getenv("REDIS_ADDRESS"), os.Getenv("REDIS_PORT")),
	})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	err = setupTimeZone()
	if err != nil {
		panic(err)
	}

	err = setupFiber()
	if err != nil {
		panic(err)
	}
}
