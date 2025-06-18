package main

import (
	"context"
	"fmt"

	"newaccess/internal/dto"
	"newaccess/internal/service"
)

// App struct
type App struct {
	ctx         context.Context
	userService service.UserService
}

// NewApp creates a new App application struct
func NewApp(userService service.UserService) *App {
	return &App{
		userService: userService,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Exporte métodos para o frontend
func (a *App) PinExists(pin string) (*dto.QueryPinReponse, error) {
	return a.userService.PinExists(context.Background(), pin)
}
