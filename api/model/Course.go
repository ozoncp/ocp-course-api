package model

import "fmt"

type Course struct {
	Id          uint64
	ClassroomId uint64
	Name        string
	Stream      string
}

func (v Course) String() string {
	return fmt.Sprintf("Course {Id: %v; ClassroomId: %v; Name: %v; Stream: %v}",
		v.Id, v.ClassroomId, v.Name, v.Stream)
}
