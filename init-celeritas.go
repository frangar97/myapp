package main

import (
	"log"
	"os"

	"github.com/frangar97/celeritas"
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

	cel.InfoLog.Println("Debug esta seteado a", cel.Debug)

	app := &application{App: cel}

	return app
}
