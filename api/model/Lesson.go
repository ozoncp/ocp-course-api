package model

import "fmt"

type Lesson struct {
	Id       uint64
	CourseId uint64
	Number   uint
	Name     string
}

func (v Lesson) String() string {
	return fmt.Sprintf("Lesson {Id: %v; CourseId: %v; Number: %v, Name: %v}",
		v.Id, v.CourseId, v.Number, v.Name)
}
