package flusher

type flusher struct {
	FlusherModelCourse
	FlusherModelLesson
}

func NewFlusher(repoCourse repoModelCourse,
	repoLesson repoModelLesson,
	batchSize int) flusher {
	return flusher{
		NewFlusherModelCourse(repoCourse, batchSize),
		NewFlusherModelLesson(repoLesson, batchSize)}
}
