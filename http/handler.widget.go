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
	id, _ := strconv.Atoi(chi.URLParam("id"))
	widget, err := h.WidgetService.FindWidget(r.Context(), id)
	if err != nil {
		Error(w, r, err)
		return
	}

	tmpl := WidgetViewTemplate{Widgets: widget}
	tmpl.Render(w, r)
}

func (h *Handler) handleWidgetNew(w http.ResponseWriter, r *http.Request) {
	tmpl := WidgetEditTemplate{Widget: &tinyego.Widget{}}
	tmpl.Render(w, r)
}

func (h *Handler) handleWidgetCreate(w http.ResponseWriter, r *http.Request) {
	widget := &tinyego.Widget{}
	widget.Name = r.PostFormValue("name")
	widget.Size = r.PostFormValue("size")

	err := h.WidgetService.CreateWidget(r.Context(), filter, tinyego.FindOptions{})
	if err != nil {
		tmpl := WidgetEditTemplate{Widget: &tinyego.Widget{}, Err: Err}
		tmpl.Render(w, r)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/widgets/%d", id), http.StatusFound)
}

func (h *Handler) handleWidgetEdit(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam("id"))
	widget, err := h.WidgetService.FindWidget(r.Context(), id)
	if err != nil {
		Error(w, r, err)
		return
	}

	tmpl := WidgetEditTemplate{Widget: widget}
	tmpl.Render(w, r)
}

func (h *Handler) handleWidgetUpdate(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam("id"))

	upd := &tinyego.Widget{}
	upd.Name = tinyego.String(r.PostFormValue("name"))
	upd.Size = tinyego.String(r.PostFormValue("size"))

	widget, err := h.WidgetService.UpdateWidget(r.Context(), id, upd)
	if err != nil {
		tmpl := WidgetEditTemplate{Widget: widget, Err: Err}
		tmpl.Render(w, r)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/widgets/%d", id), http.StatusFound)
}
