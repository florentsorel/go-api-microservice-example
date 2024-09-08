package service

import (
	"github.com/tracker-tv/actor-api/internal/repository"
)

type Service struct {
	ActorService ActorService
}

func New(r repository.Repository) *Service {
	actorRepository := repository.NewRepository(r.DB)
	return &Service{
		ActorService: ActorService{r: &actorRepository},
	}
}
