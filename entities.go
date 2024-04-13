package banner

import (
	"encoding/json"
	"time"
)

type Banner struct {
	Id        int             `json:"id" db:"id"`
	Content   json.RawMessage `json:"content" db:"content"`
	Tag       []int           `json:"tag" db:"tag_id"`
	Feature   int             `json:"feature" db:"feature_id"`
	IsActive  bool            `json:"is_active" db:"is_active"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
	UpdatedAt time.Time       `json:"updated_at" db:"updated_at"`
}

type UserBanner struct {
	Tag     int
	Feature int
	Content json.RawMessage `json:"content" db:"content"`
}

type Tag struct {
	ID int `json:"id"`
}

type Feature struct {
	ID int `json:"id"`
}
