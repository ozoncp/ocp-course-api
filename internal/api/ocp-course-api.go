package api

import (
	"context"

	pb "github.com/ozoncp/ocp-course-api/pkg/ocp-course-api"
)

type ocpCourseApiServer struct {
	pb.UnimplementedOcpCourseApiServer
}

func (ocpCourseApiServer) ListCoursesV1(context.Context, *pb.ListCoursesV1Request) (*pb.ListCoursesV1Response, error) {
	c1 := pb.Course{
		Id:          1,
		ClassroomId: 1,
		Name:        "Ozon C++ to Go",
		Stream:      "Spring'21",
	}
	return &pb.ListCoursesV1Response{Courses: []*pb.Course{&c1}}, nil
}

func NewOcpCourseApi() pb.OcpCourseApiServer {
	return &ocpCourseApiServer{}
}
