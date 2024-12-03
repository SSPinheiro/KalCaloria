package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Comida struct {
	gorm.Model
	Nome    string
	Tipo    string
	Caloria int64
}

type ComidaResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Nome      string    `json:"nome"`
	Tipo      string    `json:"tipo"`
	Caloria   int64     `json:"caloria"`
}
