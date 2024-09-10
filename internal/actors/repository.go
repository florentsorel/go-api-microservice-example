package actors

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/tracker-tv/actor-api/internal/data"
)

type ActorRepository interface {
	GetActor(id int) (*data.Actor, error)
}

type PgRepository struct {
	DB *pgx.Conn
}

func NewActorRepository(db *pgx.Conn) ActorRepository {
	return &PgRepository{DB: db}
}

func (r *PgRepository) GetActor(id int) (*data.Actor, error) {
	actor := &data.Actor{}
	err := r.DB.QueryRow(context.Background(), "SELECT * FROM actor WHERE id = $1", id).Scan(&actor.ID, &actor.Name, &actor.CreatedAt, &actor.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return actor, nil
}
