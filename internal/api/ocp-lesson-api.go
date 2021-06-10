package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozoncp/ocp-course-api/pkg/ocp-lesson-api"
)

var (
	l1 = pb.Lesson{
		Id:       131,
		CourseId: 1,
		Number:   1,
		Name:     "Packages, variables, and functions.",
	}

	l2 = pb.Lesson{
		Id:       132,
		CourseId: 1,
		Number:   2,
		Name:     "Flow control statements: for, if, else, switch and defer",
	}

	l3 = pb.Lesson{
		Id:       133,
		CourseId: 2,
		Number:   1,
		Name:     "Welcome to every language to Go",
	}
)

type ocpLessonApiServer struct {
	pb.UnimplementedOcpLessonApiServer
}

func (*ocpLessonApiServer) ListLessonsV1(
	ctx context.Context,
	req *pb.ListLessonsV1Request,
) (*pb.ListLessonsV1Response, error) {
	log.Info().Msgf("ListLessonsV1Request %v", req)
	switch req.GetCourseId() {
	case 0:
		return &pb.ListLessonsV1Response{Lessons: []*pb.Lesson{&l1, &l2, &l3}}, nil
	case 1:
		return &pb.ListLessonsV1Response{Lessons: []*pb.Lesson{&l1, &l2}}, nil
	case 2:
		return &pb.ListLessonsV1Response{Lessons: []*pb.Lesson{&l3}}, nil
	}
	return &pb.ListLessonsV1Response{Lessons: []*pb.Lesson{}}, nil
}

func (*ocpLessonApiServer) DescribeLessonV1(
	ctx context.Context,
	req *pb.DescribeLessonV1Request,
) (*pb.DescribeLessonV1Response, error) {
	log.Info().Msgf("DescribeLessonV1Request: %v", req)

	switch req.GetLessonId() {
	case l1.Id:
		return &pb.DescribeLessonV1Response{Lesson: &l1}, nil
	case l2.Id:
		return &pb.DescribeLessonV1Response{Lesson: &l2}, nil
	case l3.Id:
		return &pb.DescribeLessonV1Response{Lesson: &l3}, nil
	}

	return nil,
		status.Errorf(
			codes.NotFound,
			"Lesson with id %v doesn't exist", req.GetLessonId())
}

func (*ocpLessonApiServer) CreateLessonV1(
	ctx context.Context,
	req *pb.CreateLessonV1Request,
) (*pb.CreateLessonV1Response, error) {
	log.Info().Msgf("CreateLessonV1Request: %v", req)

	if err := req.GetLesson().Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Wrong request: %v", err)
	}

	if req.GetLesson().Id == l1.Id || req.GetLesson().Id == l2.Id || req.GetLesson().Id == l3.Id {
		return nil, status.Errorf(codes.AlreadyExists, codes.Unimplemented.String())
	}
	return &pb.CreateLessonV1Response{LessonId: req.GetLesson().Id}, nil
}

func (*ocpLessonApiServer) RemoveLessonV1(
	ctx context.Context,
	req *pb.RemoveLessonV1Request,
) (*pb.RemoveLessonV1Response, error) {
	log.Info().Msgf("RemoveLessonV1Request: %v", req)
	return nil, status.Errorf(codes.Unimplemented, codes.Unimplemented.String())
}

func NewOcpLessonApi() pb.OcpLessonApiServer {
	return &ocpLessonApiServer{}
}
