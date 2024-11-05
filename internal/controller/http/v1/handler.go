package v1

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/probuborka/go_final_project/internal/entity"
)

type Handler struct {
	//message service.Message
}

func New() *Handler {
	return &Handler{
		//message: services.Message,
	}
}

func (h Handler) Init() *chi.Mux {
	r := chi.NewRouter()

	//web
	r.Handle("/", http.FileServer(http.Dir(entity.WebDir)))

	// // Создать сообщение
	// r.Post("/message", h.postMessage)

	// // Получить статистику
	// r.Get("/statistics", h.getStatistics)

	return r
}
