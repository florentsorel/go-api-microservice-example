package actors

import (
	"github.com/tracker-tv/actor-api/internal/data"
)

type ActorService interface {
	FindActorById(id int) (*data.Actor, error)
}

type ActorServiceImpl struct {
	Repo ActorRepository
}

func NewActorService(repo ActorRepository) ActorService {
	return &ActorServiceImpl{Repo: repo}
}

func (s *ActorServiceImpl) FindActorById(id int) (*data.Actor, error) {
	actor, err := s.Repo.GetActor(id)
	if err != nil {
		return nil, err
	}

	return actor, nil
}
