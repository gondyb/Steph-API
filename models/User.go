package models

// User represents a user of the application, that may be either a teacher or a student
type User struct {
	Base
	Name  string
	Role  string
	Email string
	Posts []Post
}
