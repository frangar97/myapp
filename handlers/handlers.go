package handlers

import (
	"net/http"

	"github.com/frangar97/celeritas"
)

type Handlers struct {
	App *celeritas.Celeritas
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	h.App.Render.Page(w, r, "home", nil, nil)
}
