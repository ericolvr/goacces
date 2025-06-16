package main

import (
	"embed"
	"log"

	"newaccess/internal/config"
	"newaccess/internal/handlers"
	"newaccess/internal/repository"
	"newaccess/internal/routes"
	"newaccess/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	db, err := config.NewSQLiteDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco SQLite: %v", err)
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	handler := handlers.NewUserHandler(service)

	router := gin.Default()
	routes.UserRoutes(router, handler)

	// Inicia o servidor Gin em uma goroutine
	go func() {
		if err := router.Run(":8080"); err != nil {
			log.Fatalf("Erro ao iniciar servidor Gin: %v", err)
		}
	}()

	app := NewApp()

	err = wails.Run(&options.App{
		Title:  "newaccess",
		Width:  780,
		Height: 480,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Println("Error:", err.Error())
	}
}
