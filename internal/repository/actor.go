package repository

import (
	"context"

	"github.com/tracker-tv/actor-api/internal/data"
)

type ActorRepository interface {
	GetActors() ([]data.Actor, error)
	GetActor(id int) (*data.Actor, error)
}

func (r *Repository) GetActor(id int) (*data.Actor, error) {
	actor := &data.Actor{}
	err := r.DB.QueryRow(context.Background(), "SELECT * FROM actor WHERE id = $1", id).Scan(&actor.ID, &actor.Name, &actor.CreatedAt, &actor.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return actor, nil
}

func (r *Repository) GetActors() ([]data.Actor, error) {
	rows, err := r.DB.Query(context.Background(), "SELECT * FROM actor")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var actors []data.Actor
	for rows.Next() {
		actor := data.Actor{}
		err := rows.Scan(&actor.ID, &actor.Name, &actor.CreatedAt, &actor.UpdatedAt)
		if err != nil {
			return nil, err
		}
		actors = append(actors, actor)
	}

	return actors, nil
}
