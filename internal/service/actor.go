package service

import (
	"context"
	"time"

	"github.com/tracker-tv/actor-api/internal/data"
	"github.com/tracker-tv/actor-api/internal/repository"
)

type ActorService struct {
	r repository.ActorRepository
}

func (s *ActorService) GetActors() ([]data.Actor, error) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	actors, err := s.r.GetActors()
	if err != nil {
		return nil, err
	}

	return actors, nil
}

func (s *ActorService) GetActor(id int) (*data.Actor, error) {
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	actor, err := s.r.GetActor(id)
	if err != nil {
		return nil, err
	}

	return actor, nil
}
