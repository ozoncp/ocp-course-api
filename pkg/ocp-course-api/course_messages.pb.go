// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: course_messages.proto

package ocp_course_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ClassroomId uint64 `protobuf:"varint,2,opt,name=classroom_id,json=classroomId,proto3" json:"classroom_id,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Stream      string `protobuf:"bytes,4,opt,name=stream,proto3" json:"stream,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_course_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_course_messages_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Course) GetClassroomId() uint64 {
	if x != nil {
		return x.ClassroomId
	}
	return 0
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Course) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

type ListCoursesV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListCoursesV1Request) Reset() {
	*x = ListCoursesV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCoursesV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCoursesV1Request) ProtoMessage() {}

func (x *ListCoursesV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_course_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCoursesV1Request.ProtoReflect.Descriptor instead.
func (*ListCoursesV1Request) Descriptor() ([]byte, []int) {
	return file_course_messages_proto_rawDescGZIP(), []int{1}
}

func (x *ListCoursesV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListCoursesV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListCoursesV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Courses []*Course `protobuf:"bytes,1,rep,name=courses,proto3" json:"courses,omitempty"`
}

func (x *ListCoursesV1Response) Reset() {
	*x = ListCoursesV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCoursesV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCoursesV1Response) ProtoMessage() {}

func (x *ListCoursesV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_course_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCoursesV1Response.ProtoReflect.Descriptor instead.
func (*ListCoursesV1Response) Descriptor() ([]byte, []int) {
	return file_course_messages_proto_rawDescGZIP(), []int{2}
}

func (x *ListCoursesV1Response) GetCourses() []*Course {
	if x != nil {
		return x.Courses
	}
	return nil
}

type DescribeCourseV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId uint64 `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
}

func (x *DescribeCourseV1Request) Reset() {
	*x = DescribeCourseV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeCourseV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCourseV1Request) ProtoMessage() {}

func (x *DescribeCourseV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_course_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCourseV1Request.ProtoReflect.Descriptor instead.
func (*DescribeCourseV1Request) Descriptor() ([]byte, []int) {
	return file_course_messages_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeCourseV1Request) GetCourseId() uint64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

type DescribeCourseV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course *Course `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *DescribeCourseV1Response) Reset() {
	*x = DescribeCourseV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeCourseV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeCourseV1Response) ProtoMessage() {}

func (x *DescribeCourseV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_course_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeCourseV1Response.ProtoReflect.Descriptor instead.
func (*DescribeCourseV1Response) Descriptor() ([]byte, []int) {
	return file_course_messages_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeCourseV1Response) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

type CreateCourseV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Course *Course `protobuf:"bytes,1,opt,name=course,proto3" json:"course,omitempty"`
}

func (x *CreateCourseV1Request) Reset() {
	*x = CreateCourseV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCourseV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourseV1Request) ProtoMessage() {}

func (x *CreateCourseV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_course_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourseV1Request.ProtoReflect.Descriptor instead.
func (*CreateCourseV1Request) Descriptor() ([]byte, []int) {
	return file_course_messages_proto_rawDescGZIP(), []int{5}
}

func (x *CreateCourseV1Request) GetCourse() *Course {
	if x != nil {
		return x.Course
	}
	return nil
}

type CreateCourseV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId uint64 `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
}

func (x *CreateCourseV1Response) Reset() {
	*x = CreateCourseV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCourseV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCourseV1Response) ProtoMessage() {}

func (x *CreateCourseV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_course_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCourseV1Response.ProtoReflect.Descriptor instead.
func (*CreateCourseV1Response) Descriptor() ([]byte, []int) {
	return file_course_messages_proto_rawDescGZIP(), []int{6}
}

func (x *CreateCourseV1Response) GetCourseId() uint64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

type RemoveCourseV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CourseId uint64 `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
}

func (x *RemoveCourseV1Request) Reset() {
	*x = RemoveCourseV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveCourseV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCourseV1Request) ProtoMessage() {}

func (x *RemoveCourseV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_course_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCourseV1Request.ProtoReflect.Descriptor instead.
func (*RemoveCourseV1Request) Descriptor() ([]byte, []int) {
	return file_course_messages_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveCourseV1Request) GetCourseId() uint64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

type RemoveCourseV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found bool `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
}

func (x *RemoveCourseV1Response) Reset() {
	*x = RemoveCourseV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_course_messages_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveCourseV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveCourseV1Response) ProtoMessage() {}

func (x *RemoveCourseV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_course_messages_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveCourseV1Response.ProtoReflect.Descriptor instead.
func (*RemoveCourseV1Response) Descriptor() ([]byte, []int) {
	return file_course_messages_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveCourseV1Response) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

var File_course_messages_proto protoreflect.FileDescriptor

var file_course_messages_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8b, 0x01, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x0c, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32,
	0x02, 0x20, 0x00, 0x52, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a,
	0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x06, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x50,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x32, 0x05, 0x18, 0xe8, 0x07, 0x20,
	0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x22, 0x49, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x56,
	0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x63, 0x70,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x22, 0x3f, 0x0a, 0x17, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02,
	0x20, 0x00, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x18,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x51, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x38, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0xa2, 0x01,
	0x02, 0x08, 0x01, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x22, 0x3e, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20,
	0x00, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x22, 0x3d, 0x0a, 0x15, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00,
	0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x16, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a, 0x6f, 0x6e, 0x63, 0x70, 0x2f,
	0x6f, 0x63, 0x70, 0x2d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d, 0x61, 0x70,
	0x69, 0x3b, 0x6f, 0x63, 0x70, 0x5f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_course_messages_proto_rawDescOnce sync.Once
	file_course_messages_proto_rawDescData = file_course_messages_proto_rawDesc
)

func file_course_messages_proto_rawDescGZIP() []byte {
	file_course_messages_proto_rawDescOnce.Do(func() {
		file_course_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_course_messages_proto_rawDescData)
	})
	return file_course_messages_proto_rawDescData
}

var file_course_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_course_messages_proto_goTypes = []interface{}{
	(*Course)(nil),                   // 0: ocp.course.api.Course
	(*ListCoursesV1Request)(nil),     // 1: ocp.course.api.ListCoursesV1Request
	(*ListCoursesV1Response)(nil),    // 2: ocp.course.api.ListCoursesV1Response
	(*DescribeCourseV1Request)(nil),  // 3: ocp.course.api.DescribeCourseV1Request
	(*DescribeCourseV1Response)(nil), // 4: ocp.course.api.DescribeCourseV1Response
	(*CreateCourseV1Request)(nil),    // 5: ocp.course.api.CreateCourseV1Request
	(*CreateCourseV1Response)(nil),   // 6: ocp.course.api.CreateCourseV1Response
	(*RemoveCourseV1Request)(nil),    // 7: ocp.course.api.RemoveCourseV1Request
	(*RemoveCourseV1Response)(nil),   // 8: ocp.course.api.RemoveCourseV1Response
}
var file_course_messages_proto_depIdxs = []int32{
	0, // 0: ocp.course.api.ListCoursesV1Response.courses:type_name -> ocp.course.api.Course
	0, // 1: ocp.course.api.DescribeCourseV1Response.course:type_name -> ocp.course.api.Course
	0, // 2: ocp.course.api.CreateCourseV1Request.course:type_name -> ocp.course.api.Course
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_course_messages_proto_init() }
func file_course_messages_proto_init() {
	if File_course_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_course_messages_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_messages_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCoursesV1Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_messages_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCoursesV1Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_messages_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeCourseV1Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_messages_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeCourseV1Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_messages_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCourseV1Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_messages_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCourseV1Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_messages_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveCourseV1Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_course_messages_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveCourseV1Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_course_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_course_messages_proto_goTypes,
		DependencyIndexes: file_course_messages_proto_depIdxs,
		MessageInfos:      file_course_messages_proto_msgTypes,
	}.Build()
	File_course_messages_proto = out.File
	file_course_messages_proto_rawDesc = nil
	file_course_messages_proto_goTypes = nil
	file_course_messages_proto_depIdxs = nil
}
