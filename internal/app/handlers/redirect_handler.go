package handlers

import (
	"mishin-shortener/internal/errors/deleted"
	"net/http"
	"strings"
)

// Обработчик, осуществляющий переброску(редирект) по сокращенному урлу.
func (h *ShortanerHandler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	toLocation, err := h.DB.Get(r.Context(), strings.Trim(r.RequestURI, "/"))

	if _, ok := err.(*deleted.DeletedError); ok { // если удаленнный
		w.WriteHeader(http.StatusGone)
	} else if err != nil { // если любая другая ошибка
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	} else { // если все хорошо
		w.Header().Set("Location", toLocation)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
}
