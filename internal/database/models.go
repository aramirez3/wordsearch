// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"time"

	"github.com/google/uuid"
)

type Grid struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Grid      string
}