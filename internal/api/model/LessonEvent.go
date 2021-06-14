package model

type LessonEventType = int

const (
	LessonCreated LessonEventType = iota
	LessonUpdated
	LessonRemoved
)

type LessonEvent struct {
	Type LessonEventType
	Body map[string]interface{}
}
