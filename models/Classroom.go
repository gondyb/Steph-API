package models

// Classroom represents the different classrooms available for a specific school
type Classroom struct {
	Base
	Name  string
	Posts []Post
}
