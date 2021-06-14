package api

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-course-api/internal/api/model"
	"github.com/ozoncp/ocp-course-api/internal/repo"
	pb "github.com/ozoncp/ocp-course-api/pkg/ocp-lesson-api"
)

type ocpLessonApiServer struct {
	repo repo.RepoModelLesson
	events chan<- model.LessonEvent
	pb.UnimplementedOcpLessonApiServer
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
	log.Info().Msgf("ListLessonsV1Request %v", req)
	lessons, err := s.repo.ListModelLessons(req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	result := make([]*pb.Lesson, 0, len(lessons))
	for _, l := range lessons {
		result = append(result, toPbLesson(l))
	}
	return &pb.ListLessonsV1Response{Lessons: result}, nil
}

func (s *ocpLessonApiServer) DescribeLessonV1(
	ctx context.Context,
	req *pb.DescribeLessonV1Request,
) (*pb.DescribeLessonV1Response, error) {
	log.Info().Msgf("DescribeLessonV1Request: %v", req)

	lesson, err := s.repo.DescribeModelLesson(req.LessonId)
	if err != nil {
		return nil, err
	}
	return &pb.DescribeLessonV1Response{Lesson: toPbLesson(lesson)}, nil
}

func (s *ocpLessonApiServer) CreateLessonV1(
	ctx context.Context,
	req *pb.CreateLessonV1Request,
) (*pb.CreateLessonV1Response, error) {
	log.Info().Msgf("CreateLessonV1Request: %v", req)

	id, err := s.repo.AddModelLesson(req.Lesson)
	if err != nil {
		return nil, err
	}
	return &pb.CreateLessonV1Response{LessonId: id}, nil
}

func (s *ocpLessonApiServer) RemoveLessonV1(
	ctx context.Context,
	req *pb.RemoveLessonV1Request,
) (*pb.RemoveLessonV1Response, error) {
	log.Info().Msgf("RemoveLessonV1Request: %v", req)
	err := s.repo.RemoveModelLesson(req.LessonId)
	if err != nil {
		return nil, err
	}
	return &pb.RemoveLessonV1Response{Found: true}, nil
}

func NewOcpLessonApi(repo repo.RepoModelLesson, events chan<- model.LessonEvent) pb.OcpLessonApiServer {
	return &ocpLessonApiServer{repo: repo, events: events}
}
