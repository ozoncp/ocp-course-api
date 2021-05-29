package main

import (
	"errors"
	"fmt"

	"github.com/ozoncp/ocp-course-api/api/model"
	"github.com/ozoncp/ocp-course-api/internal/flusher"
)

type fakeRepoCourse struct{}
type fakeRepoLesson struct{}

func (this *fakeRepoCourse) AddModelCourse(v model.Course) (uint64, error) {
	return 0, errors.New("not implemented")
}

func (this *fakeRepoCourse) AddModelCourses(v []model.Course) error {
	return errors.New("not implemented")
}

func (this *fakeRepoLesson) AddModelLesson(v model.Lesson) (uint64, error) {
	return 0, errors.New("not implemented")
}

func (this *fakeRepoLesson) AddModelLessons(v []model.Lesson) error {
	return errors.New("not implemented")
}

func main() {
	fmt.Println("Project: ocp-course-api")
	fmt.Println("Author: Aleksei Shashev")
	fmt.Println("Site: https://github.com/ozoncp/ocp-course-api")

	f := flusher.NewFlusher(&fakeRepoCourse{}, &fakeRepoLesson{}, 32)
	fmt.Println(f.FlushModelCourse([]model.Course{{Id: 0, ClassroomId: 0, Name: "c0", Stream: "s0"}}))
	fmt.Println(f.FlushModelLesson([]model.Lesson{{Id: 0, CourseId: 0, Number: 0, Name: "l0"}}))
}
