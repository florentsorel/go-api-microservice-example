package actors

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func TestGetActor_Success(t *testing.T) {
	db, err := pgx.Connect(context.Background(), "postgres://test-user:test-password@localhost:5432/tracker-tv?sslmode=disable")
	assert.NoError(t, err)

	sut := &PgRepository{DB: db}

	createdAt, err := time.Parse(time.RFC3339, "2020-01-01T00:00:00Z")

	actor, err := sut.GetActor(1)

	assert.NoError(t, err)
	assert.NotNil(t, actor)
	assert.Equal(t, 1, actor.ID)
	assert.Equal(t, "Bryan Cranston", actor.Name)
	assert.Equal(t, createdAt, actor.CreatedAt)
	assert.Equal(t, sql.NullTime{}, actor.UpdatedAt)
}
