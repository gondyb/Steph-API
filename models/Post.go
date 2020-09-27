package models

import (
	"github.com/google/uuid"
)

// Post repreents a post for a specific class
type Post struct {
	Base
	Title       string
	Body        string
	UserID      uuid.UUID
	ClassroomID uuid.UUID
}
