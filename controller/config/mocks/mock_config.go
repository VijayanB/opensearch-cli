// Code generated by MockGen. DO NOT EDIT.
// Source: opensearch-cli/controller/config (interfaces: Controller)

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "opensearch-cli/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockController is a mock of Controller interface
type MockController struct {
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
}

// MockControllerMockRecorder is the mock recorder for MockController
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockController) Read() (entity.Config, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read")
	ret0, _ := ret[0].(entity.Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockControllerMockRecorder) Read() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockController)(nil).Read))
}

// Write mocks base method
func (m *MockController) Write(arg0 entity.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockControllerMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockController)(nil).Write), arg0)
}
