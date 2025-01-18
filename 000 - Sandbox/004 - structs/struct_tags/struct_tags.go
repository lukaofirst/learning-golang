package struct_tags

import "time"

type Note struct {
	Title     string    `json:"title"` // use `json:"-"` to remove a key
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}