package api

import (
	"context"
	"database/sql"
	"errors"

	"github.com/opentracing/opentracing-go"
	otl "github.com/opentracing/opentracing-go/log"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/ozoncp/ocp-course-api/internal/api/model"
	im "github.com/ozoncp/ocp-course-api/internal/metrics"
	"github.com/ozoncp/ocp-course-api/internal/repo"
	"github.com/ozoncp/ocp-course-api/internal/utils/commons"
	pb "github.com/ozoncp/ocp-course-api/pkg/ocp-lesson-api"
)

type ocpLessonApiServer struct {
	pb.UnimplementedOcpLessonApiServer

	repo      repo.RepoModelLesson
	events    chan<- model.LessonEvent
	batchSize commons.NaturalInt
}

func toPbLesson(l model.Lesson) *pb.Lesson {
	return &pb.Lesson{
		Id:       l.GetId(),
		CourseId: l.GetCourseId(),
		Number:   l.GetNumber(),
		Name:     l.GetName(),
	}
}

func (s *ocpLessonApiServer) ListLessonsV1(
	ctx context.Context,
	req *pb.ListLessonsV1Request,
) (*pb.ListLessonsV1Response, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ListLessonsV1")
	defer span.Finish()

	log.Info().Msgf("ListLessonsV1Request %v", req)
	lessons, err := s.repo.List(req.Limit, req.Offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			span.SetTag("result", "NotFound")
			return &pb.ListLessonsV1Response{Lessons: nil}, nil
		}
		span.SetTag("result", "Fail")
		span.LogFields(otl.String("error", err.Error()))
		return nil, err
	}
	result := make([]*pb.Lesson, 0, len(lessons))
	for _, l := range lessons {
		result = append(result, toPbLesson(l))
	}
	span.SetTag("result", "OK")
	span.LogFields(otl.Int("records", len(result)))
	return &pb.ListLessonsV1Response{Lessons: result}, nil
}

func (s *ocpLessonApiServer) DescribeLessonV1(
	ctx context.Context,
	req *pb.DescribeLessonV1Request,
) (*pb.DescribeLessonV1Response, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "DescribeLessonV1")
	defer span.Finish()
	span.LogFields(otl.Uint64("id", req.LessonId))

	log.Info().Msgf("DescribeLessonV1Request: %v", req)

	lesson, err := s.repo.Describe(req.LessonId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			span.SetTag("result", "NotFound")
			return nil,
				status.Errorf(codes.NotFound,
					"Lesson with ID %v wasn't found.", req.LessonId)
		}
		span.SetTag("result", "Fail")
		span.LogFields(otl.String("error", err.Error()))
		return nil, err
	}
	span.SetTag("result", "OK")
	return &pb.DescribeLessonV1Response{Lesson: toPbLesson(lesson)}, nil
}

func (s *ocpLessonApiServer) CreateLessonV1(
	ctx context.Context,
	req *pb.CreateLessonV1Request,
) (*pb.CreateLessonV1Response, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "CreateLessonV1")
	defer span.Finish()

	log.Info().Msgf("CreateLessonV1Request: %v", req)
	im.IncIncomingRequests("CreateLessonV1")

	id, err := s.repo.Add(req.Lesson)
	if err != nil {
		span.SetTag("result", "Fail")
		span.LogFields(otl.String("error", err.Error()))
		return nil, err
	}
	s.events <- model.LessonEvent{
		Type: model.LessonCreated,
		Body: map[string]interface{}{"id": id},
	}
	im.IncIncomingRequestsSuccess("CreateLessonV1")
	span.SetTag("result", "OK")
	return &pb.CreateLessonV1Response{LessonId: id}, nil
}

func (s *ocpLessonApiServer) RemoveLessonV1(
	ctx context.Context,
	req *pb.RemoveLessonV1Request,
) (*pb.RemoveLessonV1Response, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "RemoveLessonV1")
	defer span.Finish()
	span.LogFields(otl.Uint64("id", req.LessonId))

	log.Info().Msgf("RemoveLessonV1Request: %v", req)
	im.IncIncomingRequests("RemoveLessonV1")
	err := s.repo.Remove(req.LessonId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			span.SetTag("result", "NotFound")
			return &pb.RemoveLessonV1Response{Found: false}, nil
		}
		span.SetTag("result", "Fail")
		span.LogFields(otl.String("error", err.Error()))
		return nil, err
	}
	s.events <- model.LessonEvent{
		Type: model.LessonRemoved,
		Body: map[string]interface{}{"id": req.LessonId},
	}
	im.IncIncomingRequestsSuccess("RemoveLessonV1")
	span.SetTag("result", "OK")
	return &pb.RemoveLessonV1Response{Found: true}, nil
}

func (s *ocpLessonApiServer) UpdateLessonV1(
	ctx context.Context,
	req *pb.UpdateLessonV1Request,
) (*pb.UpdateLessonV1Response, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "UpdateLessonV1")
	defer span.Finish()
	span.LogFields(otl.Uint64("id", req.Lesson.Id))

	im.IncIncomingRequests("UpdateLessonV1")
	err := s.repo.Update(req.Lesson)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			span.SetTag("result", "NotFound")
			return &pb.UpdateLessonV1Response{Found: false}, nil
		}
		span.SetTag("result", "Fail")
		span.LogFields(otl.String("error", err.Error()))
		return nil, err
	}
	s.events <- model.LessonEvent{
		Type: model.LessonUpdated,
		Body: map[string]interface{}{"id": req.Lesson.GetId()},
	}
	im.IncIncomingRequestsSuccess("UpdateLessonV1")
	span.SetTag("result", "OK")
	return &pb.UpdateLessonV1Response{Found: true}, nil
}

func (s *ocpLessonApiServer) MultiCreateLessonV1(
	ctx context.Context,
	req *pb.MultiCreateLessonV1Request,
) (*pb.MultiCreateLessonV1Response, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "MultiCreateLessonV1")
	defer span.Finish()

	im.IncIncomingRequests("MultiCreateLessonV1")
	srcLen := len(req.Lessons)
	size := s.batchSize.ToInt()
	for i := 0; i < srcLen; i += size {
		childSpan := opentracing.StartSpan("batch handler", opentracing.ChildOf(span.Context()))
		end := commons.MinInt(i+size, srcLen)
		ls := req.Lessons[i:end:end]
		childSpan.LogFields(
			otl.Int("records", len(ls)),
		)
		ds := make([]model.Lesson, 0, size)
		for _, l := range ls {
			ds = append(ds, l)
		}
		err := s.repo.Adds(ds)
		if err != nil {
			childSpan.SetTag("result", "Fail")
			childSpan.LogFields(otl.String("error", err.Error()))
			childSpan.Finish()
			span.SetTag("result", "Fail")
			span.LogFields(otl.String("error", err.Error()))
			return &pb.MultiCreateLessonV1Response{
				NotSaved: req.Lessons[i:],
				Error:    err.Error(),
			}, nil
		}
		for _, l := range ls {
			s.events <- model.LessonEvent{
				Type: model.LessonCreated,
				Body: map[string]interface{}{"id": l.GetId()},
			}
		}
		childSpan.SetTag("result", "OK")
		childSpan.Finish()
		if i+size >= srcLen {
			break
		}
	}
	span.SetTag("result", "OK")
	im.IncIncomingRequestsSuccess("MultiCreateLessonV1")
	return &pb.MultiCreateLessonV1Response{}, nil
}

func NewOcpLessonApi(
	repo repo.RepoModelLesson,
	events chan<- model.LessonEvent,
	batchSize commons.NaturalInt,
) pb.OcpLessonApiServer {
	return &ocpLessonApiServer{repo: repo, events: events, batchSize: batchSize}
}
