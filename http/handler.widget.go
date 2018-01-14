package http

import (
	"net/http"

	"github.com/benbjohnson/tiny-ego"
	"github.com/go-chi/chi"
)

func (h *Handler) registerWidgetRoutes(r chi.Router) {
	r.Route("/widgets", func(r chi.Router) {
		r.Get("/", h.handleWidgetIndex)
		r.Get("/{id}", h.handleWidgetView)
		r.Get("/new", h.handleWidgetNew)
		r.Post("/new", h.handleWidgetCreate)
		r.Get("/{id}/edit", h.handleWidgetEdit)
		r.Patch("/{id}/edit", h.handleWidgetUpdate)
	})
}

func (h *Handler) handleWidgetIndex(w http.ResponseWriter, r *http.Request) {
	var filter tinyego.WidgetFilter
	filter.Query = FormStringPtr(r, "q")
	filter.Size = FormStringPtr(r, "size")

	widgets, err := h.WidgetService.FindWidgets(r.Context(), filter, tinyego.FindOptions{})
	if err != nil {
		Error(w, r, err)
		return
	}

	tmpl := WidgetIndexTemplate{Widgets: widgets, Filter: filter}
	tmpl.Render(w, r)
}

func (h *Handler) handleWidgetView(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleWidgetNew(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleWidgetCreate(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleWidgetEdit(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleWidgetUpdate(w http.ResponseWriter, r *http.Request) {

}
