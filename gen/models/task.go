package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	UUID      uuid.UUID `pg:"uuid,pk,type:uuid"`
	Status    Status    `pg:"status"`
	CreatedAt time.Time `pg:"created_at" pg:"default:now()"`
}
