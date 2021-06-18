package impl

import (
	"context"
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"

	"github.com/ozoncp/ocp-course-api/internal/api/model"
	"github.com/ozoncp/ocp-course-api/internal/repo"
)

const (
	courseTableName = "courses"
)

type repoCourse struct {
	ctx context.Context
	db  *sqlx.DB
}

type rowCourse struct {
	id          uint64
	classroomId uint64
	name        string
	stream      string
}

func (r *rowCourse) GetId() uint64          { return r.id }
func (r *rowCourse) GetClassroomId() uint64 { return r.classroomId }
func (r *rowCourse) GetName() string        { return r.name }
func (r *rowCourse) GetStream() string      { return r.stream }

func NewRepoCourse(ctx context.Context, db *sqlx.DB) repo.RepoModelCourse {
	return &repoCourse{ctx: ctx, db: db}
}

func (r *repoCourse) DescribeModelCourse(id uint64) (model.Course, error) {
	query := sq.Select("id", "classroom_id", "name", "stream").
		From(courseTableName).
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	var course rowCourse
	if err := query.QueryRowContext(r.ctx).Scan(
		&course.id,
		&course.classroomId,
		&course.name,
		&course.stream,
	); err != nil {
		return nil, err
	}
	return &course, nil
}
func (r *repoCourse) ListModelCourses(limit uint64, offset uint64) ([]model.Course, error) {
	query := sq.Select("id", "classroom_id", "name", "stream").
		From(courseTableName).
		Limit(limit).
		Offset(offset).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	rows, err := query.QueryContext(r.ctx)
	if err != nil {
		return nil, err
	}

	var courses []model.Course
	for rows.Next() {
		var course rowCourse
		if err := rows.Scan(
			&course.id,
			&course.classroomId,
			&course.name,
			&course.stream,
		); err != nil {
			continue
		}
		courses = append(courses, &course)
	}
	return courses, nil
}

func (r *repoCourse) AddModelCourse(v model.Course) (uint64, error) {
	query := sq.Insert(courseTableName).
		Columns("id", "classroom_id", "name", "stream").
		Values(v.GetId(), v.GetClassroomId(), v.GetName(), v.GetStream()).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	if _, err := query.ExecContext(r.ctx); err != nil {
		return 0, err
	}
	return v.GetId(), nil
}

func (r *repoCourse) AddModelCourses(vs []model.Course) error {
	query := sq.Insert(courseTableName).
		Columns("id", "classroom_id", "name", "stream").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	for _, v := range vs {
		query = query.
			Values(v.GetId(), v.GetClassroomId(), v.GetName(), v.GetStream())
	}

	_, err := query.ExecContext(r.ctx)
	return err
}

func (r *repoCourse) RemoveModelCourse(id uint64) error {
	query := sq.Delete(courseTableName).
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	result, err := query.ExecContext(r.ctx)
	if count, err := result.RowsAffected(); err == nil && count == 0 {
		return sql.ErrNoRows
	}
	return err
}

func (r *repoCourse) UpdateModelCourse(v model.Course) error {
	query := sq.Update(courseTableName).
		Set("classroom_id", v.GetClassroomId()).
		Set("name", v.GetName()).
		Set("stream", v.GetStream()).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	if _, err := query.ExecContext(r.ctx); err != nil {
		return err
	}
	return nil
}
