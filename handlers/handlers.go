package handlers

import (
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/frangar97/celeritas"
)

type Handlers struct {
	App *celeritas.Celeritas
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	err := h.App.Render.Page(w, r, "home", nil, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}

func (h *Handlers) SessionTest(w http.ResponseWriter, r *http.Request) {
	myData := "bar"

	h.App.Session.Put(r.Context(), "foo", myData)

	myValue := h.App.Session.GetString(r.Context(), "foo")
	vars := make(jet.VarMap)
	vars.Set("foo", myValue)

	err := h.App.Render.Page(w, r, "sessions", vars, nil)
	if err != nil {
		h.App.ErrorLog.Println("error rendering:", err)
	}
}
