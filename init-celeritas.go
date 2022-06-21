package main

import (
	"log"
	"os"

	"github.com/frangar97/celeritas"
	"github.com/frangar97/myapp/handlers"
)

func initApplication() *application {
	path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	//Iniciamos celeritas
	cel := &celeritas.Celeritas{}
	err = cel.New(path)

	if err != nil {
		log.Fatal(err)
	}

	cel.AppName = "MyApp"

	myHandlers := &handlers.Handlers{
		App: cel,
	}

	app := &application{App: cel, Handlers: myHandlers}

	app.App.Routes = app.routes()

	return app
}
