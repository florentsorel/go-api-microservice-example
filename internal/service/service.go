package service

import (
	"github.com/jackc/pgx/v5"
	"github.com/tracker-tv/actor-api/internal/repository"
)

type Service struct {
	ActorService ActorServiceInterface
}

func New(db *pgx.Conn) *Service {
	repo := repository.NewRepository(db)
	return &Service{
		ActorService: &actorService{r: &repo},
	}
}
