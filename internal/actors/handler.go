package actors

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

type ActorHandler struct {
	Logger       *slog.Logger
	ActorService ActorService
}

func NewActorHandler(logger *slog.Logger, actorService ActorService) *ActorHandler {
	return &ActorHandler{Logger: logger, ActorService: actorService}
}

func (h ActorHandler) GetActor(w http.ResponseWriter, r *http.Request) {
	id, err := readIDPath(r)

	actors, err := h.ActorService.FindActorById(id)
	if err != nil {
		h.Logger.Error("error getting actor", "err", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSON(w, actors)
}

func writeJSON(w http.ResponseWriter, v any) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write(js)
}

func readIDPath(r *http.Request) (int, error) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return 0, err
	}
	return id, nil
}
