package utils

import (
	"log"
	"os"

	. "github.com/ozoncp/ocp-course-api/api/model"
	. "github.com/ozoncp/ocp-course-api/internal/utils/slice"
)

func SplitToBulksInt(xs []int, batchSize int) [][]int {
	return SliceGenerateWindowsInt(xs, batchSize, batchSize)
}

func SplitToBulksCourse(xs []Course, batchSize int) [][]Course {
	return SliceGenerateWindowsModelCourse(xs, batchSize, batchSize)
}

func SplitToBulksLesson(xs []Lesson, batchSize int) [][]Lesson {
	return SliceGenerateWindowsModelLesson(xs, batchSize, batchSize)
}

func SliceToMapCourse(xs []Course) map[uint64]Course {
	return SliceToMapModelCourseUint64(xs, func(c Course) uint64 { return c.Id })
}

func SliceToMapLesson(xs []Lesson) map[uint64]Lesson {
	return SliceToMapModelLessonUint64(xs, func(c Lesson) uint64 { return c.Id })
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
