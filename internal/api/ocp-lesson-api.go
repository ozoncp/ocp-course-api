package api

import (
	"context"

	"github.com/rs/zerolog/log"

	"github.com/ozoncp/ocp-course-api/internal/api/model"
	"github.com/ozoncp/ocp-course-api/internal/repo"
	"github.com/ozoncp/ocp-course-api/internal/utils/commons"
	pb "github.com/ozoncp/ocp-course-api/pkg/ocp-lesson-api"
)

type ocpLessonApiServer struct {
	repo      repo.RepoModelLesson
	events    chan<- model.LessonEvent
	batchSize commons.NaturalInt
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
	s.events <- model.LessonEvent{
		Type: model.LessonCreated,
		Body: map[string]interface{}{"id": id},
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
	s.events <- model.LessonEvent{
		Type: model.LessonRemoved,
		Body: map[string]interface{}{"id": req.LessonId},
	}
	return &pb.RemoveLessonV1Response{Found: true}, nil
}

func (s *ocpLessonApiServer) UpdateLessonV1(
	ctx context.Context,
	req *pb.UpdateLessonV1Request,
) (*pb.UpdateLessonV1Response, error) {
	err := s.repo.UpdateModelLesson(req.Lesson)
	if err != nil {
		return &pb.UpdateLessonV1Response{Found: false}, err
	}
	s.events <- model.LessonEvent{
		Type: model.LessonUpdated,
		Body: map[string]interface{}{"id": req.Lesson.GetId()},
	}
	return &pb.UpdateLessonV1Response{Found: true}, nil
}

func (s *ocpLessonApiServer) MultiCreateLessonV1(
	ctx context.Context,
	req *pb.MultiCreateLessonV1Request,
) (*pb.MultiCreateLessonV1Response, error) {
	srcLen := len(req.Lessons)
	size := s.batchSize.ToInt()
	for i := 0; i < srcLen; i += size {
		end := commons.MinInt(i+size, srcLen)
		ls := req.Lessons[i:end:end]
		ds := make([]model.Lesson, 0, size)
		for _, l := range ls {
			ds = append(ds, l)
		}
		err := s.repo.AddModelLessons(ds)
		if err != nil {
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
		if i+size >= srcLen {
			break
		}
	}
	return &pb.MultiCreateLessonV1Response{}, nil
}

func NewOcpLessonApi(
	repo repo.RepoModelLesson,
	events chan<- model.LessonEvent,
	batchSize commons.NaturalInt,
) pb.OcpLessonApiServer {
	return &ocpLessonApiServer{repo: repo, events: events, batchSize: batchSize}
}
