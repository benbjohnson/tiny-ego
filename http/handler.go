package http

import (
	"net/http"

	"github.com/benbjohnson/tiny-ego"
	"github.com/benbjohnson/tiny-ego/http/template"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Handler struct {
	mux *chi.Mux

	WidgetService tinyego.WidgetService
}

func NewHandler() *Handler {
	h := &Handler{mux: chi.NewMux()}
	h.mux.Use(middleware.RedirectSlashes)
	h.mux.Use(middleware.RealIP)
	h.mux.Use(middleware.Recoverer)
	h.mux.Use(middleware.DefaultCompress)
	h.mux.Use(OverrideMethod)

	h.mux.Group(func(r chi.Router) {
		r.Get("/", h.handleIndex)
		h.registerWidgetRoutes(r)
	})

	return h
}

func (h *Handler) handleIndex(w http.ResponseWriter, r *http.Request) {
	var tmpl template.IndexTemplate
	tmpl.Render(r.Context(), w)
}

// OverrideMethod overrides the request method if "_method" form field exists.
func OverrideMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			switch method := r.FormValue("_method"); method {
			case http.MethodPatch, http.MethodPut, http.MethodDelete:
				r.Method = method
			}
		}
		next.ServeHTTP(w, r)
	})
}

func WriteHeader(w http.ResponseWriter, err error) {
	if err == nil {
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
}

func Error(w http.ResponseWriter, r *http.Request, err error) {
	WriteHeader(w, err)
	tmpl := template.ErrorTemplate{Err: err}
	tmpl.Render(r.Context(), w)
}
