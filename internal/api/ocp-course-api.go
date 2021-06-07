package api

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/ozoncp/ocp-course-api/pkg/ocp-course-api"
)

var (
	c1 = pb.Course{
		Id:          1,
		ClassroomId: 1,
		Name:        "Ozon C++ to Go",
		Stream:      "Spring'21",
	}

	c2 = pb.Course{
		Id:          2,
		ClassroomId: 1,
		Name:        "Ozon every language to Go",
		Stream:      "Summer'21",
	}
)

type ocpCourseApiServer struct {
	pb.UnimplementedOcpCourseApiServer
}

func (*ocpCourseApiServer) ListCoursesV1(
	ctx context.Context,
	req *pb.ListCoursesV1Request,
) (*pb.ListCoursesV1Response, error) {

	log.Info().Msgf("ListCoursesV1Request %v", req)
	return &pb.ListCoursesV1Response{Courses: []*pb.Course{&c1, &c2}}, nil
}

func (*ocpCourseApiServer) DescribeCourseV1(
	ctx context.Context,
	req *pb.DescribeCourseV1Request,
) (*pb.DescribeCourseV1Response, error) {
	log.Info().Msgf("DescribeCourseV1Request: %v", req)

	switch req.GetCourseId() {
	case 1:
		return &pb.DescribeCourseV1Response{Course: &c1}, nil
	case 2:
		return &pb.DescribeCourseV1Response{Course: &c2}, nil
	}

	return nil,
		status.Errorf(
			codes.NotFound,
			"Course with id %v doesn't exist", req.GetCourseId())
}

func (*ocpCourseApiServer) CreateCourseV1(
	ctx context.Context,
	req *pb.CreateCourseV1Request,
) (*pb.CreateCourseV1Response, error) {
	log.Info().Msgf("CreateCourseV1Request: %v", req)

	if err := req.GetCourse().Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Wrong request: %v", err)
	}

	if req.GetCourse().Id == 1 || req.GetCourse().Id == 2 {
		return nil, status.Errorf(codes.AlreadyExists, codes.Unimplemented.String())
	}
	return &pb.CreateCourseV1Response{CourseId: req.GetCourse().Id}, nil
}

func (*ocpCourseApiServer) RemoveCourseV1(
	ctx context.Context,
	req *pb.RemoveCourseV1Request,
) (*pb.RemoveCourseV1Response, error) {
	log.Info().Msgf("RemoveCourseV1Request: %v", req)
	return nil, status.Errorf(codes.Unimplemented, codes.Unimplemented.String())
}

func NewOcpCourseApi() pb.OcpCourseApiServer {
	return &ocpCourseApiServer{}
}
