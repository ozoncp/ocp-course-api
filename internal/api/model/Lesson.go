package model

type Lesson interface {
	GetId() uint64
	GetCourseId() uint64
	GetNumber() uint32
	GetName() string
}
