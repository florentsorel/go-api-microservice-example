package handler

import (
	"context"
	"net/http"
	"time"
)

func (h *Handler) GetActors(w http.ResponseWriter, r *http.Request) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	actors, err := h.Service.ActorService.GetActors()
	if err != nil {
		h.Logger.Error("error getting actors", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, actors)
}

func (h *Handler) GetActor(w http.ResponseWriter, r *http.Request) {
	id, err := readIDPath(r)

	if err != nil {
		http.Error(w, "Url Param 'id' must be an integer", http.StatusBadRequest)
		return
	}

	actor, err := h.Service.ActorService.GetActor(id)

	writeJSON(w, actor)
}
