package banner

import (
	"encoding/json"
	"errors"
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

type UpdateBanner struct {
	Content   *json.RawMessage
	Tag       *[]int
	Feature   *int
	IsActive  *bool
	UpdatedAt time.Time
}

func (i UpdateBanner) Validate() error {
	if i.Content == nil && i.Tag == nil && i.Feature == nil && i.IsActive == nil {
		return errors.New("update banner has no values")
	}

	return nil
}

type Tag struct {
	ID int `json:"id"`
}

type Feature struct {
	ID int `json:"id"`
}
