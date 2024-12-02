// Code generated by MockGen. DO NOT EDIT.
// Source: /home/wil/dev/sro/gamesever-service/pkg/repository/connection.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=/home/wil/dev/sro/gamesever-service/pkg/repository/connection.go -destination=/home/wil/dev/sro/gamesever-service/pkg/mocks/connection_mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gameserver "github.com/ShatteredRealms/gameserver-service/pkg/model/gameserver"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockConnectionRepository is a mock of ConnectionRepository interface.
type MockConnectionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockConnectionRepositoryMockRecorder
	isgomock struct{}
}

// MockConnectionRepositoryMockRecorder is the mock recorder for MockConnectionRepository.
type MockConnectionRepositoryMockRecorder struct {
	mock *MockConnectionRepository
}

// NewMockConnectionRepository creates a new mock instance.
func NewMockConnectionRepository(ctrl *gomock.Controller) *MockConnectionRepository {
	mock := &MockConnectionRepository{ctrl: ctrl}
	mock.recorder = &MockConnectionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConnectionRepository) EXPECT() *MockConnectionRepositoryMockRecorder {
	return m.recorder
}

// CreatePendingConnection mocks base method.
func (m *MockConnectionRepository) CreatePendingConnection(ctx context.Context, character, serverName string) (*gameserver.PendingConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePendingConnection", ctx, character, serverName)
	ret0, _ := ret[0].(*gameserver.PendingConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePendingConnection indicates an expected call of CreatePendingConnection.
func (mr *MockConnectionRepositoryMockRecorder) CreatePendingConnection(ctx, character, serverName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePendingConnection", reflect.TypeOf((*MockConnectionRepository)(nil).CreatePendingConnection), ctx, character, serverName)
}

// DeletePendingConnection mocks base method.
func (m *MockConnectionRepository) DeletePendingConnection(ctx context.Context, id *uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePendingConnection", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePendingConnection indicates an expected call of DeletePendingConnection.
func (mr *MockConnectionRepositoryMockRecorder) DeletePendingConnection(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePendingConnection", reflect.TypeOf((*MockConnectionRepository)(nil).DeletePendingConnection), ctx, id)
}

// FindPendingConnection mocks base method.
func (m *MockConnectionRepository) FindPendingConnection(ctx context.Context, id *uuid.UUID) (*gameserver.PendingConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindPendingConnection", ctx, id)
	ret0, _ := ret[0].(*gameserver.PendingConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindPendingConnection indicates an expected call of FindPendingConnection.
func (mr *MockConnectionRepositoryMockRecorder) FindPendingConnection(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindPendingConnection", reflect.TypeOf((*MockConnectionRepository)(nil).FindPendingConnection), ctx, id)
}
