// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repo/repo_generic.go

// Package mock_repo is a generated GoMock package.
package mock_repo

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	repo "github.com/ozoncp/ocp-course-api/internal/repo"
)

// MockRepoTValue is a mock of RepoTValue interface.
type MockRepoTValue struct {
	ctrl     *gomock.Controller
	recorder *MockRepoTValueMockRecorder
}

// MockRepoTValueMockRecorder is the mock recorder for MockRepoTValue.
type MockRepoTValueMockRecorder struct {
	mock *MockRepoTValue
}

// NewMockRepoTValue creates a new mock instance.
func NewMockRepoTValue(ctrl *gomock.Controller) *MockRepoTValue {
	mock := &MockRepoTValue{ctrl: ctrl}
	mock.recorder = &MockRepoTValueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepoTValue) EXPECT() *MockRepoTValueMockRecorder {
	return m.recorder
}

// AddTValue mocks base method.
func (m *MockRepoTValue) AddTValue(v repo.TValue) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTValue", v)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTValue indicates an expected call of AddTValue.
func (mr *MockRepoTValueMockRecorder) AddTValue(v interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTValue", reflect.TypeOf((*MockRepoTValue)(nil).AddTValue), v)
}

// AddTValues mocks base method.
func (m *MockRepoTValue) AddTValues(vs []repo.TValue) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTValues", vs)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTValues indicates an expected call of AddTValues.
func (mr *MockRepoTValueMockRecorder) AddTValues(vs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTValues", reflect.TypeOf((*MockRepoTValue)(nil).AddTValues), vs)
}

// DescribeTValue mocks base method.
func (m *MockRepoTValue) DescribeTValue(id uint64) (repo.TValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTValue", id)
	ret0, _ := ret[0].(repo.TValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTValue indicates an expected call of DescribeTValue.
func (mr *MockRepoTValueMockRecorder) DescribeTValue(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTValue", reflect.TypeOf((*MockRepoTValue)(nil).DescribeTValue), id)
}

// ListTValues mocks base method.
func (m *MockRepoTValue) ListTValues(limit, offset uint64) ([]repo.TValue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTValues", limit, offset)
	ret0, _ := ret[0].([]repo.TValue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTValues indicates an expected call of ListTValues.
func (mr *MockRepoTValueMockRecorder) ListTValues(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTValues", reflect.TypeOf((*MockRepoTValue)(nil).ListTValues), limit, offset)
}

// RemoveTValue mocks base method.
func (m *MockRepoTValue) RemoveTValue(id uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTValue", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTValue indicates an expected call of RemoveTValue.
func (mr *MockRepoTValueMockRecorder) RemoveTValue(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTValue", reflect.TypeOf((*MockRepoTValue)(nil).RemoveTValue), id)
}
