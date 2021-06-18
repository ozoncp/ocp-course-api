// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: lesson_service.proto

package ocp_lesson_api

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Lesson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CourseId uint64 `protobuf:"varint,2,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	Number   uint32 `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Lesson) Reset() {
	*x = Lesson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lesson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lesson) ProtoMessage() {}

func (x *Lesson) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lesson.ProtoReflect.Descriptor instead.
func (*Lesson) Descriptor() ([]byte, []int) {
	return file_lesson_service_proto_rawDescGZIP(), []int{0}
}

func (x *Lesson) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Lesson) GetCourseId() uint64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *Lesson) GetNumber() uint32 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Lesson) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListLessonsV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint64 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListLessonsV1Request) Reset() {
	*x = ListLessonsV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLessonsV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLessonsV1Request) ProtoMessage() {}

func (x *ListLessonsV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLessonsV1Request.ProtoReflect.Descriptor instead.
func (*ListLessonsV1Request) Descriptor() ([]byte, []int) {
	return file_lesson_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListLessonsV1Request) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListLessonsV1Request) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ListLessonsV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lessons []*Lesson `protobuf:"bytes,1,rep,name=lessons,proto3" json:"lessons,omitempty"`
}

func (x *ListLessonsV1Response) Reset() {
	*x = ListLessonsV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLessonsV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLessonsV1Response) ProtoMessage() {}

func (x *ListLessonsV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLessonsV1Response.ProtoReflect.Descriptor instead.
func (*ListLessonsV1Response) Descriptor() ([]byte, []int) {
	return file_lesson_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListLessonsV1Response) GetLessons() []*Lesson {
	if x != nil {
		return x.Lessons
	}
	return nil
}

type DescribeLessonV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LessonId uint64 `protobuf:"varint,1,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
}

func (x *DescribeLessonV1Request) Reset() {
	*x = DescribeLessonV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeLessonV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeLessonV1Request) ProtoMessage() {}

func (x *DescribeLessonV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeLessonV1Request.ProtoReflect.Descriptor instead.
func (*DescribeLessonV1Request) Descriptor() ([]byte, []int) {
	return file_lesson_service_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeLessonV1Request) GetLessonId() uint64 {
	if x != nil {
		return x.LessonId
	}
	return 0
}

type DescribeLessonV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lesson *Lesson `protobuf:"bytes,1,opt,name=lesson,proto3" json:"lesson,omitempty"`
}

func (x *DescribeLessonV1Response) Reset() {
	*x = DescribeLessonV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescribeLessonV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeLessonV1Response) ProtoMessage() {}

func (x *DescribeLessonV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeLessonV1Response.ProtoReflect.Descriptor instead.
func (*DescribeLessonV1Response) Descriptor() ([]byte, []int) {
	return file_lesson_service_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeLessonV1Response) GetLesson() *Lesson {
	if x != nil {
		return x.Lesson
	}
	return nil
}

type CreateLessonV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lesson *Lesson `protobuf:"bytes,1,opt,name=lesson,proto3" json:"lesson,omitempty"`
}

func (x *CreateLessonV1Request) Reset() {
	*x = CreateLessonV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLessonV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLessonV1Request) ProtoMessage() {}

func (x *CreateLessonV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLessonV1Request.ProtoReflect.Descriptor instead.
func (*CreateLessonV1Request) Descriptor() ([]byte, []int) {
	return file_lesson_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateLessonV1Request) GetLesson() *Lesson {
	if x != nil {
		return x.Lesson
	}
	return nil
}

type CreateLessonV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LessonId uint64 `protobuf:"varint,1,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
}

func (x *CreateLessonV1Response) Reset() {
	*x = CreateLessonV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLessonV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLessonV1Response) ProtoMessage() {}

func (x *CreateLessonV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLessonV1Response.ProtoReflect.Descriptor instead.
func (*CreateLessonV1Response) Descriptor() ([]byte, []int) {
	return file_lesson_service_proto_rawDescGZIP(), []int{6}
}

func (x *CreateLessonV1Response) GetLessonId() uint64 {
	if x != nil {
		return x.LessonId
	}
	return 0
}

type RemoveLessonV1Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LessonId uint64 `protobuf:"varint,1,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
}

func (x *RemoveLessonV1Request) Reset() {
	*x = RemoveLessonV1Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveLessonV1Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveLessonV1Request) ProtoMessage() {}

func (x *RemoveLessonV1Request) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveLessonV1Request.ProtoReflect.Descriptor instead.
func (*RemoveLessonV1Request) Descriptor() ([]byte, []int) {
	return file_lesson_service_proto_rawDescGZIP(), []int{7}
}

func (x *RemoveLessonV1Request) GetLessonId() uint64 {
	if x != nil {
		return x.LessonId
	}
	return 0
}

type RemoveLessonV1Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Found bool `protobuf:"varint,1,opt,name=found,proto3" json:"found,omitempty"`
}

func (x *RemoveLessonV1Response) Reset() {
	*x = RemoveLessonV1Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lesson_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveLessonV1Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveLessonV1Response) ProtoMessage() {}

func (x *RemoveLessonV1Response) ProtoReflect() protoreflect.Message {
	mi := &file_lesson_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveLessonV1Response.ProtoReflect.Descriptor instead.
func (*RemoveLessonV1Response) Descriptor() ([]byte, []int) {
	return file_lesson_service_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveLessonV1Response) GetFound() bool {
	if x != nil {
		return x.Found
	}
	return false
}

var File_lesson_service_proto protoreflect.FileDescriptor

var file_lesson_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x6f, 0x63, 0x70, 0x2e, 0x6c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01,
	0x0a, 0x06, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x24, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x50, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x73, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0a, 0xfa, 0x42,
	0x07, 0x32, 0x05, 0x18, 0xe8, 0x07, 0x20, 0x00, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x49, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x4c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x30, 0x0a, 0x07, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x73, 0x22, 0x3f, 0x0a, 0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a,
	0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x49, 0x64, 0x22, 0x4a, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2e, 0x0a, 0x06, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x22,
	0x51, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x56,
	0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x6c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x42, 0x08, 0xfa, 0x42, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x06, 0x6c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x22, 0x3e, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x09,
	0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x49, 0x64, 0x22, 0x3d, 0x0a, 0x15, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x09, 0x6c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x2e, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x32, 0x8b, 0x04, 0x0a, 0x0c, 0x4f, 0x63, 0x70, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x41,
	0x70, 0x69, 0x12, 0x71, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x73, 0x56, 0x31, 0x12, 0x24, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73,
	0x56, 0x31, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x6f, 0x63, 0x70, 0x2e,
	0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x73, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x73, 0x12, 0x86, 0x01, 0x0a, 0x10, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x12, 0x27, 0x2e, 0x6f, 0x63, 0x70,
	0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x7c,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x56, 0x31,
	0x12, 0x25, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x56, 0x31,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x65, 0x73, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x0b, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x73, 0x3a, 0x06, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x12, 0x80, 0x01, 0x0a,
	0x0e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x12,
	0x25, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x6f, 0x63, 0x70, 0x2e, 0x6c, 0x65, 0x73,
	0x73, 0x6f, 0x6e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x4c, 0x65,
	0x73, 0x73, 0x6f, 0x6e, 0x56, 0x31, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x2a, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x42,
	0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x7a,
	0x6f, 0x6e, 0x63, 0x70, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x63, 0x70, 0x2d, 0x6c, 0x65, 0x73, 0x73,
	0x6f, 0x6e, 0x2d, 0x61, 0x70, 0x69, 0x3b, 0x6f, 0x63, 0x70, 0x5f, 0x6c, 0x65, 0x73, 0x73, 0x6f,
	0x6e, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lesson_service_proto_rawDescOnce sync.Once
	file_lesson_service_proto_rawDescData = file_lesson_service_proto_rawDesc
)

func file_lesson_service_proto_rawDescGZIP() []byte {
	file_lesson_service_proto_rawDescOnce.Do(func() {
		file_lesson_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_lesson_service_proto_rawDescData)
	})
	return file_lesson_service_proto_rawDescData
}

var file_lesson_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_lesson_service_proto_goTypes = []interface{}{
	(*Lesson)(nil),                   // 0: ocp.lesson.api.Lesson
	(*ListLessonsV1Request)(nil),     // 1: ocp.lesson.api.ListLessonsV1Request
	(*ListLessonsV1Response)(nil),    // 2: ocp.lesson.api.ListLessonsV1Response
	(*DescribeLessonV1Request)(nil),  // 3: ocp.lesson.api.DescribeLessonV1Request
	(*DescribeLessonV1Response)(nil), // 4: ocp.lesson.api.DescribeLessonV1Response
	(*CreateLessonV1Request)(nil),    // 5: ocp.lesson.api.CreateLessonV1Request
	(*CreateLessonV1Response)(nil),   // 6: ocp.lesson.api.CreateLessonV1Response
	(*RemoveLessonV1Request)(nil),    // 7: ocp.lesson.api.RemoveLessonV1Request
	(*RemoveLessonV1Response)(nil),   // 8: ocp.lesson.api.RemoveLessonV1Response
}
var file_lesson_service_proto_depIdxs = []int32{
	0, // 0: ocp.lesson.api.ListLessonsV1Response.lessons:type_name -> ocp.lesson.api.Lesson
	0, // 1: ocp.lesson.api.DescribeLessonV1Response.lesson:type_name -> ocp.lesson.api.Lesson
	0, // 2: ocp.lesson.api.CreateLessonV1Request.lesson:type_name -> ocp.lesson.api.Lesson
	1, // 3: ocp.lesson.api.OcpLessonApi.ListLessonsV1:input_type -> ocp.lesson.api.ListLessonsV1Request
	3, // 4: ocp.lesson.api.OcpLessonApi.DescribeLessonV1:input_type -> ocp.lesson.api.DescribeLessonV1Request
	5, // 5: ocp.lesson.api.OcpLessonApi.CreateLessonV1:input_type -> ocp.lesson.api.CreateLessonV1Request
	7, // 6: ocp.lesson.api.OcpLessonApi.RemoveLessonV1:input_type -> ocp.lesson.api.RemoveLessonV1Request
	2, // 7: ocp.lesson.api.OcpLessonApi.ListLessonsV1:output_type -> ocp.lesson.api.ListLessonsV1Response
	4, // 8: ocp.lesson.api.OcpLessonApi.DescribeLessonV1:output_type -> ocp.lesson.api.DescribeLessonV1Response
	6, // 9: ocp.lesson.api.OcpLessonApi.CreateLessonV1:output_type -> ocp.lesson.api.CreateLessonV1Response
	8, // 10: ocp.lesson.api.OcpLessonApi.RemoveLessonV1:output_type -> ocp.lesson.api.RemoveLessonV1Response
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_lesson_service_proto_init() }
func file_lesson_service_proto_init() {
	if File_lesson_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lesson_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lesson); i {
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
		file_lesson_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLessonsV1Request); i {
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
		file_lesson_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLessonsV1Response); i {
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
		file_lesson_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeLessonV1Request); i {
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
		file_lesson_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescribeLessonV1Response); i {
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
		file_lesson_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLessonV1Request); i {
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
		file_lesson_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLessonV1Response); i {
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
		file_lesson_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveLessonV1Request); i {
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
		file_lesson_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveLessonV1Response); i {
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
			RawDescriptor: file_lesson_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lesson_service_proto_goTypes,
		DependencyIndexes: file_lesson_service_proto_depIdxs,
		MessageInfos:      file_lesson_service_proto_msgTypes,
	}.Build()
	File_lesson_service_proto = out.File
	file_lesson_service_proto_rawDesc = nil
	file_lesson_service_proto_goTypes = nil
	file_lesson_service_proto_depIdxs = nil
}
