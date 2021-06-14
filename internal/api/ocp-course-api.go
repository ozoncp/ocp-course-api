package api

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-course-api/internal/api/model"
	"github.com/ozoncp/ocp-course-api/internal/repo"
	pb "github.com/ozoncp/ocp-course-api/pkg/ocp-course-api"
)

type ocpCourseApiServer struct {
	repo repo.RepoModelCourse
	events chan<- model.CourseEvent
	pb.UnimplementedOcpCourseApiServer
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

	log.Info().Msgf("ListCoursesV1Request %v", req)
	courses, err := s.repo.ListModelCourses(req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	result := make([]*pb.Course, 0, len(courses))
	for _, c := range courses {
		result = append(result, toPbCourse(c))
	}
	return &pb.ListCoursesV1Response{Courses: result}, nil
}

func (s *ocpCourseApiServer) DescribeCourseV1(
	ctx context.Context,
	req *pb.DescribeCourseV1Request,
) (*pb.DescribeCourseV1Response, error) {
	log.Info().Msgf("DescribeCourseV1Request: %v", req)

	course, err := s.repo.DescribeModelCourse(req.CourseId)
	if err != nil {
		return nil, err
	}
	return &pb.DescribeCourseV1Response{Course: toPbCourse(course)}, nil
}

func (s *ocpCourseApiServer) CreateCourseV1(
	ctx context.Context,
	req *pb.CreateCourseV1Request,
) (*pb.CreateCourseV1Response, error) {
	log.Info().Msgf("CreateCourseV1Request: %v", req)

	id, err := s.repo.AddModelCourse(req.Course)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCourseV1Response{CourseId: id}, nil
}

func (s *ocpCourseApiServer) RemoveCourseV1(
	ctx context.Context,
	req *pb.RemoveCourseV1Request,
) (*pb.RemoveCourseV1Response, error) {
	log.Info().Msgf("RemoveCourseV1Request: %v", req)
	err := s.repo.RemoveModelCourse(req.CourseId)
	if err != nil {
		return &pb.RemoveCourseV1Response{Found: false}, err
	}
	return &pb.RemoveCourseV1Response{Found: true}, nil
}

func NewOcpCourseApi(repo repo.RepoModelCourse,
	events chan<- model.CourseEvent,
) pb.OcpCourseApiServer {
	return &ocpCourseApiServer{repo: repo, events: events}
}
