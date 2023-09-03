package main

import (
	"log"
	"os"

	"github.com/frangar97/celeritas"
	"github.com/frangar97/myapp/data"
	"github.com/frangar97/myapp/handlers"
)

func initApplication() *application {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	cel := &celeritas.Celeritas{}
	err = cel.New(path)
	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "myapp"

	handlers := &handlers.Handlers{
		App: cel,
	}

	app := &application{
		App:      cel,
		Handlers: handlers,
	}

	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)

	return app
}
