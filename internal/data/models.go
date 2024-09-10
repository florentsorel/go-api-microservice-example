package data

import (
	"database/sql"
	"time"
)

type Actor struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}
