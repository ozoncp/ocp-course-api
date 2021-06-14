package model

type CourseEventType = int

const (
	CourseCreated CourseEventType = iota
	CourseUpdated
	CourseRemoved
)

type CourseEvent struct {
	Type CourseEventType
	Body map[string]interface{}
}
