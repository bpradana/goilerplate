package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/bpradana/goilerplate/cmd/db"
	"github.com/bpradana/goilerplate/pkg/users"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("[app] [server] error loading .env file, err: ", err.Error())
	}

	// Connect to database
	db, err := db.ConnectDB()
	if err != nil {
		log.Println("[app] [server] error connecting to database, err: ", err.Error())
	}

	// Create echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	userRepo := users.NewRepository(db)
	userUsecase := users.NewUsecase(userRepo)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
