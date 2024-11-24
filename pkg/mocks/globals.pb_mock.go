// Code generated by MockGen. DO NOT EDIT.
// Source: /home/wil/dev/sro/dimension-service/pkg/pb/globals.pb.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=/home/wil/dev/sro/dimension-service/pkg/pb/globals.pb.go -destination=/home/wil/dev/sro/dimension-service/pkg/mocks/globals.pb_mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockisUserTarget_Target is a mock of isUserTarget_Target interface.
type MockisUserTarget_Target struct {
	ctrl     *gomock.Controller
	recorder *MockisUserTarget_TargetMockRecorder
	isgomock struct{}
}

// MockisUserTarget_TargetMockRecorder is the mock recorder for MockisUserTarget_Target.
type MockisUserTarget_TargetMockRecorder struct {
	mock *MockisUserTarget_Target
}

// NewMockisUserTarget_Target creates a new mock instance.
func NewMockisUserTarget_Target(ctrl *gomock.Controller) *MockisUserTarget_Target {
	mock := &MockisUserTarget_Target{ctrl: ctrl}
	mock.recorder = &MockisUserTarget_TargetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisUserTarget_Target) EXPECT() *MockisUserTarget_TargetMockRecorder {
	return m.recorder
}

// isUserTarget_Target mocks base method.
func (m *MockisUserTarget_Target) isUserTarget_Target() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isUserTarget_Target")
}

// isUserTarget_Target indicates an expected call of isUserTarget_Target.
func (mr *MockisUserTarget_TargetMockRecorder) isUserTarget_Target() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isUserTarget_Target", reflect.TypeOf((*MockisUserTarget_Target)(nil).isUserTarget_Target))
}

// MockisCharacterTarget_Target is a mock of isCharacterTarget_Target interface.
type MockisCharacterTarget_Target struct {
	ctrl     *gomock.Controller
	recorder *MockisCharacterTarget_TargetMockRecorder
	isgomock struct{}
}

// MockisCharacterTarget_TargetMockRecorder is the mock recorder for MockisCharacterTarget_Target.
type MockisCharacterTarget_TargetMockRecorder struct {
	mock *MockisCharacterTarget_Target
}

// NewMockisCharacterTarget_Target creates a new mock instance.
func NewMockisCharacterTarget_Target(ctrl *gomock.Controller) *MockisCharacterTarget_Target {
	mock := &MockisCharacterTarget_Target{ctrl: ctrl}
	mock.recorder = &MockisCharacterTarget_TargetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisCharacterTarget_Target) EXPECT() *MockisCharacterTarget_TargetMockRecorder {
	return m.recorder
}

// isCharacterTarget_Target mocks base method.
func (m *MockisCharacterTarget_Target) isCharacterTarget_Target() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isCharacterTarget_Target")
}

// isCharacterTarget_Target indicates an expected call of isCharacterTarget_Target.
func (mr *MockisCharacterTarget_TargetMockRecorder) isCharacterTarget_Target() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isCharacterTarget_Target", reflect.TypeOf((*MockisCharacterTarget_Target)(nil).isCharacterTarget_Target))
}

// MockisDimensionTarget_Target is a mock of isDimensionTarget_Target interface.
type MockisDimensionTarget_Target struct {
	ctrl     *gomock.Controller
	recorder *MockisDimensionTarget_TargetMockRecorder
	isgomock struct{}
}

// MockisDimensionTarget_TargetMockRecorder is the mock recorder for MockisDimensionTarget_Target.
type MockisDimensionTarget_TargetMockRecorder struct {
	mock *MockisDimensionTarget_Target
}

// NewMockisDimensionTarget_Target creates a new mock instance.
func NewMockisDimensionTarget_Target(ctrl *gomock.Controller) *MockisDimensionTarget_Target {
	mock := &MockisDimensionTarget_Target{ctrl: ctrl}
	mock.recorder = &MockisDimensionTarget_TargetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisDimensionTarget_Target) EXPECT() *MockisDimensionTarget_TargetMockRecorder {
	return m.recorder
}

// isDimensionTarget_Target mocks base method.
func (m *MockisDimensionTarget_Target) isDimensionTarget_Target() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isDimensionTarget_Target")
}

// isDimensionTarget_Target indicates an expected call of isDimensionTarget_Target.
func (mr *MockisDimensionTarget_TargetMockRecorder) isDimensionTarget_Target() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isDimensionTarget_Target", reflect.TypeOf((*MockisDimensionTarget_Target)(nil).isDimensionTarget_Target))
}

// MockisMapTarget_Target is a mock of isMapTarget_Target interface.
type MockisMapTarget_Target struct {
	ctrl     *gomock.Controller
	recorder *MockisMapTarget_TargetMockRecorder
	isgomock struct{}
}

// MockisMapTarget_TargetMockRecorder is the mock recorder for MockisMapTarget_Target.
type MockisMapTarget_TargetMockRecorder struct {
	mock *MockisMapTarget_Target
}

// NewMockisMapTarget_Target creates a new mock instance.
func NewMockisMapTarget_Target(ctrl *gomock.Controller) *MockisMapTarget_Target {
	mock := &MockisMapTarget_Target{ctrl: ctrl}
	mock.recorder = &MockisMapTarget_TargetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockisMapTarget_Target) EXPECT() *MockisMapTarget_TargetMockRecorder {
	return m.recorder
}

// isMapTarget_Target mocks base method.
func (m *MockisMapTarget_Target) isMapTarget_Target() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "isMapTarget_Target")
}

// isMapTarget_Target indicates an expected call of isMapTarget_Target.
func (mr *MockisMapTarget_TargetMockRecorder) isMapTarget_Target() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isMapTarget_Target", reflect.TypeOf((*MockisMapTarget_Target)(nil).isMapTarget_Target))
}
