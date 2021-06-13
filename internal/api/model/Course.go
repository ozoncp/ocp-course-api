package model

type Course interface {
	GetId() uint64
	GetClassroomId() uint64
	GetName() string
	GetStream() string
}
