package impl_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ozoncp/ocp-course-api/internal/api/model"
	"github.com/ozoncp/ocp-course-api/internal/repo"
	"github.com/ozoncp/ocp-course-api/internal/repo/impl"
	ocp_lesson_api "github.com/ozoncp/ocp-course-api/pkg/ocp-lesson-api"
)

func LessonEqual(t *testing.T, expected model.Lesson, actual model.Lesson) {
	t.Helper()
	assert.Equal(t, expected.GetId(), actual.GetId())
	assert.Equal(t, expected.GetCourseId(), actual.GetCourseId())
	assert.Equal(t, expected.GetNumber(), actual.GetNumber())
	assert.Equal(t, expected.GetName(), actual.GetName())
}

func PrepareTestLesson(t *testing.T) repo.RepoModelLesson {
	t.Helper()
	_, err := db.Exec("TRUNCATE TABLE lessons RESTART IDENTITY;")
	assert.Nil(t, err)
	return impl.NewRepoLesson(context.Background(), db)
}

func TestRepoLessonAddAndDescripbe(t *testing.T) {
	repo := PrepareTestLesson(t)
	d := &ocp_lesson_api.Lesson{
		Id: 1, CourseId: 1, Number: 1, Name: "intro"}

	id, err := repo.AddModelLesson(d)
	assert.Nil(t, err)
	assert.Equal(t, d.GetId(), id)

	got, err := repo.DescribeModelLesson(d.GetId())
	assert.Nil(t, err)
	LessonEqual(t, d, got)

	err = repo.RemoveModelLesson(d.GetId())
	assert.Nil(t, err)

	_, err = repo.DescribeModelLesson(d.GetId())
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, sql.ErrNoRows)
}

func TestRepoLessonMulipleAddAndList(t *testing.T) {
	repo := PrepareTestLesson(t)
	ds := []model.Lesson{
		&ocp_lesson_api.Lesson{Id: 1, CourseId: 1, Number: 1, Name: "introduction"},
		&ocp_lesson_api.Lesson{Id: 2, CourseId: 1, Number: 2, Name: "basic syntax"},
		&ocp_lesson_api.Lesson{Id: 3, CourseId: 1, Number: 3, Name: "loops"},
		&ocp_lesson_api.Lesson{Id: 4, CourseId: 2, Number: 1, Name: "welcome"},
		&ocp_lesson_api.Lesson{Id: 5, CourseId: 2, Number: 2, Name: "basic"},
	}

	err := repo.AddModelLessons(ds)
	assert.Nil(t, err)

	got1, err := repo.ListModelLessons(3, 0)
	assert.Nil(t, err)
	LessonEqual(t, ds[0], got1[0])
	LessonEqual(t, ds[1], got1[1])
	LessonEqual(t, ds[2], got1[2])

	got2, err := repo.ListModelLessons(3, 3)
	assert.Nil(t, err)
	LessonEqual(t, ds[3], got2[0])
	LessonEqual(t, ds[4], got2[1])
}

func TestRepoLessonUpdate(t *testing.T) {
	repo := PrepareTestLesson(t)
	d1 := &ocp_lesson_api.Lesson{
		Id: 1, CourseId: 1, Number: 1, Name: "intro"}
	d2 := &ocp_lesson_api.Lesson{
		Id: d1.GetId(), CourseId: 2, Number: 5, Name: "welcome"}

	_, err := repo.AddModelLesson(d1)
	assert.Nil(t, err)

	err = repo.UpdateModelLesson(d2)
	assert.Nil(t, err)

	got, err := repo.DescribeModelLesson(d1.GetId())
	assert.Nil(t, err)
	LessonEqual(t, d2, got)
}
