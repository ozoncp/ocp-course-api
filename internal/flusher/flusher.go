package flusher

import (
	"github.com/ozoncp/ocp-course-api/internal/utils/commons"
)

type flusher struct {
	FlusherModelCourse
	FlusherModelLesson
}

func NewFlusher(repoCourse repoModelCourse,
	repoLesson repoModelLesson,
	batchSize commons.NaturalInt) flusher {
	return flusher{
		NewFlusherModelCourse(repoCourse, batchSize),
		NewFlusherModelLesson(repoLesson, batchSize)}
}
