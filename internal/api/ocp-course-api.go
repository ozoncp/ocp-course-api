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
	pb "github.com/ozoncp/ocp-course-api/pkg/ocp-course-api"
)

type ocpCourseApiServer struct {
	pb.UnimplementedOcpCourseApiServer

	repo      repo.RepoModelCourse
	events    chan<- model.CourseEvent
	batchSize commons.NaturalInt
}

func toPbCourse(c model.Course) *pb.Course {
	return &pb.Course{
		Id:          c.GetId(),
		ClassroomId: c.GetClassroomId(),
		Name:        c.GetName(),
		Stream:      c.GetStream(),
	}
}

func (s *ocpCourseApiServer) ListCoursesV1(
	ctx context.Context,
	req *pb.ListCoursesV1Request,
) (*pb.ListCoursesV1Response, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "ListCoursesV1")
	defer span.Finish()

	log.Info().Msgf("ListCoursesV1Request %v", req)
	courses, err := s.repo.ListModelCourses(req.Limit, req.Offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			span.SetTag("result", "NotFound")
			return &pb.ListCoursesV1Response{Courses: nil}, nil
		}
		span.SetTag("result", "Fail")
		span.LogFields(otl.String("error", err.Error()))
		return nil, err
	}
	result := make([]*pb.Course, 0, len(courses))
	for _, c := range courses {
		result = append(result, toPbCourse(c))
	}
	span.SetTag("result", "OK")
	span.LogFields(otl.Int("records", len(result)))
	return &pb.ListCoursesV1Response{Courses: result}, nil
}

func (s *ocpCourseApiServer) DescribeCourseV1(
	ctx context.Context,
	req *pb.DescribeCourseV1Request,
) (*pb.DescribeCourseV1Response, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "DescribeCourseV1")
	defer span.Finish()
	span.LogFields(otl.Uint64("id", req.CourseId))

	log.Info().Msgf("DescribeCourseV1Request: %v", req)

	course, err := s.repo.DescribeModelCourse(req.CourseId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			span.SetTag("result", "NotFound")
			span.LogFields(otl.Uint64("id", req.CourseId))
			return nil,
				status.Errorf(codes.NotFound,
					"Course with ID %v wasn't found.", req.CourseId)
		}
		span.SetTag("result", "Fail")
		span.LogFields(otl.String("error", err.Error()))
		return nil, err
	}
	span.SetTag("result", "OK")
	return &pb.DescribeCourseV1Response{Course: toPbCourse(course)}, nil
}

func (s *ocpCourseApiServer) CreateCourseV1(
	ctx context.Context,
	req *pb.CreateCourseV1Request,
) (*pb.CreateCourseV1Response, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "CreateCourseV1")
	defer span.Finish()

	log.Info().Msgf("CreateCourseV1Request: %v", req)
	im.IncIncomingRequests("CreateCourseV1")

	id, err := s.repo.AddModelCourse(req.Course)
	if err != nil {
		span.SetTag("result", "Fail")
		span.LogFields(otl.String("error", err.Error()))
		return nil, err
	}
	s.events <- model.CourseEvent{
		Type: model.CourseCreated,
		Body: map[string]interface{}{"id": id},
	}
	im.IncIncomingRequestsSuccess("CreateCourseV1")
	span.SetTag("result", "OK")
	return &pb.CreateCourseV1Response{CourseId: id}, nil
}

func (s *ocpCourseApiServer) RemoveCourseV1(
	ctx context.Context,
	req *pb.RemoveCourseV1Request,
) (*pb.RemoveCourseV1Response, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "RemoveCourseV1")
	defer span.Finish()
	span.LogFields(otl.Uint64("id", req.CourseId))

	log.Info().Msgf("RemoveCourseV1Request: %v", req)
	im.IncIncomingRequests("RemoveCourseV1")
	err := s.repo.RemoveModelCourse(req.CourseId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			span.SetTag("result", "NotFound")
			return &pb.RemoveCourseV1Response{Found: false}, nil
		}
		span.SetTag("result", "Fail")
		span.LogFields(otl.String("error", err.Error()))
		return nil, err
	}
	s.events <- model.CourseEvent{
		Type: model.CourseRemoved,
		Body: map[string]interface{}{"id": req.CourseId},
	}
	im.IncIncomingRequestsSuccess("RemoveCourseV1")
	span.SetTag("result", "OK")
	return &pb.RemoveCourseV1Response{Found: true}, nil
}

func (s *ocpCourseApiServer) UpdateCourseV1(
	ctx context.Context,
	req *pb.UpdateCourseV1Request,
) (*pb.UpdateCourseV1Response, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "UpdateCourseV1")
	defer span.Finish()
	span.LogFields(otl.Uint64("id", req.Course.Id))

	im.IncIncomingRequests("UpdateCourseV1")
	err := s.repo.UpdateModelCourse(req.Course)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			span.SetTag("result", "NotFound")
			return &pb.UpdateCourseV1Response{Found: false}, nil
		}
		span.SetTag("result", "Fail")
		span.LogFields(otl.String("error", err.Error()))
		return nil, err
	}
	s.events <- model.CourseEvent{
		Type: model.CourseUpdated,
		Body: map[string]interface{}{"id": req.Course.GetId()},
	}
	im.IncIncomingRequestsSuccess("UpdateCourseV1")
	span.SetTag("result", "OK")
	return &pb.UpdateCourseV1Response{Found: true}, nil
}

func (s *ocpCourseApiServer) MultiCreateCourseV1(
	ctx context.Context,
	req *pb.MultiCreateCourseV1Request,
) (*pb.MultiCreateCourseV1Response, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "MultiCreateCourseV1")
	defer span.Finish()

	im.IncIncomingRequests("MultiCreateCourseV1")
	srcLen := len(req.Courses)
	size := s.batchSize.ToInt()
	for i := 0; i < srcLen; i += size {
		childSpan := opentracing.StartSpan("batch handler", opentracing.ChildOf(span.Context()))
		end := commons.MinInt(i+size, srcLen)
		cs := req.Courses[i:end:end]
		childSpan.LogFields(otl.Int("records", len(cs)))
		ds := make([]model.Course, 0, size)
		for _, c := range cs {
			ds = append(ds, c)
		}
		err := s.repo.AddModelCourses(ds)
		if err != nil {
			childSpan.SetTag("result", "Fail")
			childSpan.LogFields(otl.String("error", err.Error()))
			childSpan.Finish()
			span.SetTag("result", "Fail")
			span.LogFields(otl.String("error", err.Error()))
			return &pb.MultiCreateCourseV1Response{
				NotSaved: req.Courses[i:],
				Error:    err.Error(),
			}, nil
		}
		for _, c := range cs {
			s.events <- model.CourseEvent{
				Type: model.CourseCreated,
				Body: map[string]interface{}{"id": c.GetId()},
			}
		}
		childSpan.SetTag("result", "OK")
		childSpan.Finish()
		if i+size >= srcLen {
			break
		}
	}
	im.IncIncomingRequestsSuccess("MultiCreateCourseV1")
	span.SetTag("result", "OK")
	return &pb.MultiCreateCourseV1Response{}, nil
}

func NewOcpCourseApi(
	repo repo.RepoModelCourse,
	events chan<- model.CourseEvent,
	batchSize commons.NaturalInt,
) pb.OcpCourseApiServer {
	return &ocpCourseApiServer{repo: repo, events: events, batchSize: batchSize}
}
