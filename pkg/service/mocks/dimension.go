// Code generated by MockGen. DO NOT EDIT.
// Source: /home/wil/dev/sro/gameserver-service/pkg/service/dimension.go
//
// Generated by this command:
//
//	mockgen -source=/home/wil/dev/sro/gameserver-service/pkg/service/dimension.go -destination=/home/wil/dev/sro/gameserver-service/pkg/service/mocks/dimension.go
//

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	game "github.com/ShatteredRealms/gameserver-service/pkg/model/game"
	uuid "github.com/google/uuid"
	gomock "go.uber.org/mock/gomock"
)

// MockDimensionService is a mock of DimensionService interface.
type MockDimensionService struct {
	ctrl     *gomock.Controller
	recorder *MockDimensionServiceMockRecorder
	isgomock struct{}
}

// MockDimensionServiceMockRecorder is the mock recorder for MockDimensionService.
type MockDimensionServiceMockRecorder struct {
	mock *MockDimensionService
}

// NewMockDimensionService creates a new mock instance.
func NewMockDimensionService(ctrl *gomock.Controller) *MockDimensionService {
	mock := &MockDimensionService{ctrl: ctrl}
	mock.recorder = &MockDimensionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDimensionService) EXPECT() *MockDimensionServiceMockRecorder {
	return m.recorder
}

// CreateDimension mocks base method.
func (m *MockDimensionService) CreateDimension(ctx context.Context, name, version, location string, mapIds []uuid.UUID) (*game.Dimension, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDimension", ctx, name, version, location, mapIds)
	ret0, _ := ret[0].(*game.Dimension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDimension indicates an expected call of CreateDimension.
func (mr *MockDimensionServiceMockRecorder) CreateDimension(ctx, name, version, location, mapIds any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDimension", reflect.TypeOf((*MockDimensionService)(nil).CreateDimension), ctx, name, version, location, mapIds)
}

// DeleteDimension mocks base method.
func (m *MockDimensionService) DeleteDimension(ctx context.Context, dimensionId *uuid.UUID) (*game.Dimension, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDimension", ctx, dimensionId)
	ret0, _ := ret[0].(*game.Dimension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDimension indicates an expected call of DeleteDimension.
func (mr *MockDimensionServiceMockRecorder) DeleteDimension(ctx, dimensionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDimension", reflect.TypeOf((*MockDimensionService)(nil).DeleteDimension), ctx, dimensionId)
}

// EditDimension mocks base method.
func (m *MockDimensionService) EditDimension(ctx context.Context, dimension *game.Dimension) (*game.Dimension, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditDimension", ctx, dimension)
	ret0, _ := ret[0].(*game.Dimension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditDimension indicates an expected call of EditDimension.
func (mr *MockDimensionServiceMockRecorder) EditDimension(ctx, dimension any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditDimension", reflect.TypeOf((*MockDimensionService)(nil).EditDimension), ctx, dimension)
}

// GetDeletedDimensions mocks base method.
func (m *MockDimensionService) GetDeletedDimensions(ctx context.Context) (game.Dimensions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeletedDimensions", ctx)
	ret0, _ := ret[0].(game.Dimensions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeletedDimensions indicates an expected call of GetDeletedDimensions.
func (mr *MockDimensionServiceMockRecorder) GetDeletedDimensions(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeletedDimensions", reflect.TypeOf((*MockDimensionService)(nil).GetDeletedDimensions), ctx)
}

// GetDimensionById mocks base method.
func (m *MockDimensionService) GetDimensionById(ctx context.Context, dimensionId *uuid.UUID) (*game.Dimension, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimensionById", ctx, dimensionId)
	ret0, _ := ret[0].(*game.Dimension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimensionById indicates an expected call of GetDimensionById.
func (mr *MockDimensionServiceMockRecorder) GetDimensionById(ctx, dimensionId any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensionById", reflect.TypeOf((*MockDimensionService)(nil).GetDimensionById), ctx, dimensionId)
}

// GetDimensions mocks base method.
func (m *MockDimensionService) GetDimensions(ctx context.Context) (game.Dimensions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimensions", ctx)
	ret0, _ := ret[0].(game.Dimensions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimensions indicates an expected call of GetDimensions.
func (mr *MockDimensionServiceMockRecorder) GetDimensions(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensions", reflect.TypeOf((*MockDimensionService)(nil).GetDimensions), ctx)
}