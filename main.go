package main

import (
	"github.com/frangar97/celeritas"
	"github.com/frangar97/myapp/data"
	"github.com/frangar97/myapp/handlers"
)

type application struct {
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
	Models   data.Models
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
