package impl_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ozoncp/ocp-course-api/internal/api/model"
	"github.com/ozoncp/ocp-course-api/internal/repo"
	"github.com/ozoncp/ocp-course-api/internal/repo/impl"
	ocp_course_api "github.com/ozoncp/ocp-course-api/pkg/ocp-course-api"
)

func CourseEqual(t *testing.T, expected model.Course, actual model.Course) {
	t.Helper()
	assert.Equal(t, expected.GetId(), actual.GetId())
	assert.Equal(t, expected.GetClassroomId(), actual.GetClassroomId())
	assert.Equal(t, expected.GetName(), actual.GetName())
	assert.Equal(t, expected.GetStream(), actual.GetStream())
}

func PrepareTestCourse(t *testing.T) repo.RepoModelCourse {
	t.Helper()
	_, err := db.Exec("TRUNCATE TABLE courses RESTART IDENTITY;")
	assert.Nil(t, err)
	return impl.NewRepoCourse(context.Background(), db)
}

func TestRepoCourseAddAndDescripbe(t *testing.T) {
	repo := PrepareTestCourse(t)
	d := &ocp_course_api.Course{
		Id: 1, ClassroomId: 1, Name: "cpp2go", Stream: "spring'21"}

	id, err := repo.AddModelCourse(d)
	assert.Nil(t, err)
	assert.Equal(t, d.GetId(), id)

	got, err := repo.DescribeModelCourse(d.GetId())
	assert.Nil(t, err)
	CourseEqual(t, d, got)

	err = repo.RemoveModelCourse(d.GetId())
	assert.Nil(t, err)

	_, err = repo.DescribeModelCourse(d.GetId())
	assert.NotNil(t, err)
	assert.ErrorIs(t, err, sql.ErrNoRows)
}

func TestRepoCourseMulipleAddAndList(t *testing.T) {
	repo := PrepareTestCourse(t)
	ds := []model.Course{
		&ocp_course_api.Course{Id: 1, ClassroomId: 1, Name: "golang", Stream: "autumn'20"},
		&ocp_course_api.Course{Id: 2, ClassroomId: 1, Name: "go-everywhene", Stream: "winter'21"},
		&ocp_course_api.Course{Id: 3, ClassroomId: 2, Name: "cpp2go", Stream: "spring'21"},
		&ocp_course_api.Course{Id: 4, ClassroomId: 3, Name: "all2go", Stream: "summer'21"},
		&ocp_course_api.Course{Id: 5, ClassroomId: 1, Name: "1,2,3,go", Stream: "autumn'21"},
	}

	err := repo.AddModelCourses(ds)
	assert.Nil(t, err)

	got1, err := repo.ListModelCourses(3, 0)
	assert.Nil(t, err)
	CourseEqual(t, ds[0], got1[0])
	CourseEqual(t, ds[1], got1[1])
	CourseEqual(t, ds[2], got1[2])

	got2, err := repo.ListModelCourses(3, 3)
	assert.Nil(t, err)
	CourseEqual(t, ds[3], got2[0])
	CourseEqual(t, ds[4], got2[1])
}

func TestRepoCourseUpdate(t *testing.T) {
	repo := PrepareTestCourse(t)
	d1 := &ocp_course_api.Course{
		Id: 1, ClassroomId: 1, Name: "cpp2go", Stream: "spring'21"}
	d2 := &ocp_course_api.Course{
		Id: d1.GetId(), ClassroomId: 2, Name: "C++ to Go", Stream: "Summer'21"}

	_, err := repo.AddModelCourse(d1)
	assert.Nil(t, err)

	err = repo.UpdateModelCourse(d2)
	assert.Nil(t, err)

	got, err := repo.DescribeModelCourse(d1.GetId())
	assert.Nil(t, err)
	CourseEqual(t, d2, got)
}
