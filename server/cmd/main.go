package main

import (
	"log"
	"os"
	"webchat/internal/api"
	"webchat/internal/model"
	"webchat/internal/store"
	"webchat/internal/ws"

	"github.com/labstack/echo/v4"
)

func main() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "ayush123")
	os.Setenv("DB_NAME", "webchat")
	os.Setenv("DB_PORT", "5432")

	store.ConnectDatabase()
	// store.DB.AutoMigrate(&ws.OutboundMessage{})
	store.DB.AutoMigrate(&model.Message{})
	hub := ws.NewHub()

	go hub.Run()

	e := echo.New()
	api.RegisterRoutes(e, hub)
	log.Println("Server started at :8080")
	e.Logger.Fatal(e.Start(":8080"))

}
