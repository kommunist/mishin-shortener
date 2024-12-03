package redirect

import (
	"mishin-shortener/internal/errors/deleted"
	"net/http"
	"strings"
)

// Обработчик, осуществляющий переброску(редирект) по сокращенному урлу.
func (h *Handler) Call(w http.ResponseWriter, r *http.Request) {
	to, err := h.Perform(r.Context(), strings.Trim(r.RequestURI, "/"))

	if err == nil {
		w.Header().Set("Location", to)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}

	if _, ok := err.(*deleted.DeletedError); ok { // если удаленнный
		w.WriteHeader(http.StatusGone)
	} else { // если любая другая ошибка
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}
