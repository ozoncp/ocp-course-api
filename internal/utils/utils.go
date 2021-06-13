package utils

import (
	"log"
	"os"

	"github.com/ozoncp/ocp-course-api/internal/api/model"
	slice "github.com/ozoncp/ocp-course-api/internal/utils/slice"
)

func SplitToBulksInt(xs []int, batchSize int) [][]int {
	return slice.GenerateWindowsInt(xs, batchSize, batchSize)
}

func SplitToBulksCourse(xs []model.Course, batchSize int) [][]model.Course {
	return slice.GenerateWindowsModelCourse(xs, batchSize, batchSize)
}

func SplitToBulksLesson(xs []model.Lesson, batchSize int) [][]model.Lesson {
	return slice.GenerateWindowsModelLesson(xs, batchSize, batchSize)
}

func ToMapCourse(xs []model.Course) map[uint64]model.Course {
	return slice.ToMapModelCourseUint64(xs, func(c model.Course) uint64 { return c.GetId() })
}

func ToMapLesson(xs []model.Lesson) map[uint64]model.Lesson {
	return slice.ToMapModelLessonUint64(xs, func(c model.Lesson) uint64 { return c.GetId() })
}

func RepeatedlyRead(file string, count int) {
	for i := 0; i < count; i++ {
		func() {
			if fh, err := os.OpenFile(file, os.O_RDONLY|os.O_EXCL, 0666); err != nil {
				log.Fatalf("Can't open the file %v. %v\n", file, err)
			} else {
				defer func() {
					if err := fh.Close(); err != nil {
						log.Fatalf("Can't close the file %v. %v", fh, err)
					}
					log.Printf("The file %v (%v) closed.\n", fh.Name(), fh.Fd())
				}()
				info, err := fh.Stat()
				if err != nil {
					log.Fatalf("Can't get a file info. %v", err)
					return
				}
				buf := make([]byte, info.Size())
				if r, err := fh.Read(buf); err != nil {
					log.Fatalf("Can't read! %v\n", err)
					return
				} else {
					log.Printf("read %v bytes\n", r)
				}
				os.Stdout.Write(buf)
			}
		}()
	}
}
