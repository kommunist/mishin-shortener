package handlers

import (
	"net/http"
)

func (h *ShortanerHandler) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	toLocation, err := h.DB.Get(r.Context(), r.RequestURI)

	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	} else {
		w.Header().Set("Location", toLocation)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return
	}
}
