package hello

import (
	"bytes"
	"fmt"
)

// FIXME
type User struct {
	CreatedAt time.Time `json:"created_at" url:"created_at,key"` // when user was created
	ID        string    `json:"id" url:"id,key"`                 // unique identifier of user
	Name      string    `json:"name" url:"name,key"`             // unique name of user
	UpdatedAt time.Time `json:"updated_at" url:"updated_at,key"` // when user was updated
}

