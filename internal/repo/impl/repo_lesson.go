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
	lessonTableName = "lessons"
)

type repoLesson struct {
	ctx context.Context
	db  *sqlx.DB
}

type rowLesson struct {
	id       uint64
	courseId uint64
	number   uint32
	name     string
}

func (r *rowLesson) GetId() uint64       { return r.id }
func (r *rowLesson) GetCourseId() uint64 { return r.courseId }
func (r *rowLesson) GetNumber() uint32   { return r.number }
func (r *rowLesson) GetName() string     { return r.name }

func NewRepoLesson(ctx context.Context, db *sqlx.DB) repo.RepoModelLesson {
	return &repoLesson{ctx: ctx, db: db}
}

func (r *repoLesson) DescribeModelLesson(id uint64) (model.Lesson, error) {
	query := sq.Select("id", "course_id", "number", "name").
		From(lessonTableName).
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	var lesson rowLesson
	if err := query.QueryRowContext(r.ctx).Scan(
		&lesson.id,
		&lesson.courseId,
		&lesson.number,
		&lesson.name,
	); err != nil {
		return nil, err
	}
	return &lesson, nil
}
func (r *repoLesson) ListModelLessons(limit uint64, offset uint64) ([]model.Lesson, error) {
	query := sq.Select("id", "course_id", "number", "name").
		From(lessonTableName).
		Limit(limit).
		Offset(offset).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	rows, err := query.QueryContext(r.ctx)
	if err != nil {
		return nil, err
	}

	var lessons []model.Lesson
	for rows.Next() {
		var lesson rowLesson
		if err := rows.Scan(
			&lesson.id,
			&lesson.courseId,
			&lesson.number,
			&lesson.name,
		); err != nil {
			continue
		}
		lessons = append(lessons, &lesson)
	}
	return lessons, nil
}

func (r *repoLesson) AddModelLesson(v model.Lesson) (uint64, error) {
	query := sq.Insert(lessonTableName).
		Columns("id", "course_id", "number", "name").
		Values(v.GetId(), v.GetCourseId(), v.GetNumber(), v.GetName()).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	if _, err := query.ExecContext(r.ctx); err != nil {
		return 0, err
	}
	return v.GetId(), nil
}

func (r *repoLesson) AddModelLessons(vs []model.Lesson) error {
	query := sq.Insert(lessonTableName).
		Columns("id", "course_id", "number", "name").
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	for _, v := range vs {
		query = query.
			Values(v.GetId(), v.GetCourseId(), v.GetNumber(), v.GetName())
	}

	_, err := query.ExecContext(r.ctx)
	return err
}

func (r *repoLesson) RemoveModelLesson(id uint64) error {
	query := sq.Delete(lessonTableName).
		Where(sq.Eq{"id": id}).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	result, err := query.ExecContext(r.ctx)
	if count, err := result.RowsAffected(); err == nil && count == 0 {
		return sql.ErrNoRows
	}
	return err
}

func (r *repoLesson) UpdateModelLesson(v model.Lesson) error {
	query := sq.Update(lessonTableName).
		Set("course_id", v.GetCourseId()).
		Set("number", v.GetNumber()).
		Set("name", v.GetName()).
		RunWith(r.db).
		PlaceholderFormat(sq.Dollar)

	if _, err := query.ExecContext(r.ctx); err != nil {
		return err
	}
	return nil
}
