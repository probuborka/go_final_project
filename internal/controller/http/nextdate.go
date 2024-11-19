package http

import (
	"net/http"
	"time"

	"github.com/probuborka/go_final_project/internal/entity"
	"github.com/probuborka/go_final_project/pkg/logger"
)

func (h handler) getNextDate(w http.ResponseWriter, r *http.Request) {

	now := r.FormValue("now")
	nowDate, err := time.Parse(entity.Format, now)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.Error(err)
		return
	}

	startDate := r.FormValue("date")

	repeat := r.FormValue("repeat")

	nextDate, err := h.task.NextDate(nowDate, startDate, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.Error(err)
		return
	}

	_, err = w.Write([]byte(nextDate))
	if err != nil {
		logger.Error(err)
		return
	}
}
