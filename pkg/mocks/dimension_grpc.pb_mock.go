// Code generated by MockGen. DO NOT EDIT.
// Source: /home/wil/dev/sro/dimension-service/pkg/pb/dimension_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -package=mocks -source=/home/wil/dev/sro/dimension-service/pkg/pb/dimension_grpc.pb.go -destination=/home/wil/dev/sro/dimension-service/pkg/mocks/dimension_grpc.pb_mock.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	pb "github.com/ShatteredRealms/dimension-service/pkg/pb"
	pb0 "github.com/ShatteredRealms/go-common-service/pkg/pb"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockDimensionServiceClient is a mock of DimensionServiceClient interface.
type MockDimensionServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockDimensionServiceClientMockRecorder
	isgomock struct{}
}

// MockDimensionServiceClientMockRecorder is the mock recorder for MockDimensionServiceClient.
type MockDimensionServiceClientMockRecorder struct {
	mock *MockDimensionServiceClient
}

// NewMockDimensionServiceClient creates a new mock instance.
func NewMockDimensionServiceClient(ctrl *gomock.Controller) *MockDimensionServiceClient {
	mock := &MockDimensionServiceClient{ctrl: ctrl}
	mock.recorder = &MockDimensionServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDimensionServiceClient) EXPECT() *MockDimensionServiceClientMockRecorder {
	return m.recorder
}

// CreateDimension mocks base method.
func (m *MockDimensionServiceClient) CreateDimension(ctx context.Context, in *pb.CreateDimensionRequest, opts ...grpc.CallOption) (*pb.Dimension, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateDimension", varargs...)
	ret0, _ := ret[0].(*pb.Dimension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDimension indicates an expected call of CreateDimension.
func (mr *MockDimensionServiceClientMockRecorder) CreateDimension(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDimension", reflect.TypeOf((*MockDimensionServiceClient)(nil).CreateDimension), varargs...)
}

// DeleteDimension mocks base method.
func (m *MockDimensionServiceClient) DeleteDimension(ctx context.Context, in *pb0.TargetId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteDimension", varargs...)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDimension indicates an expected call of DeleteDimension.
func (mr *MockDimensionServiceClientMockRecorder) DeleteDimension(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDimension", reflect.TypeOf((*MockDimensionServiceClient)(nil).DeleteDimension), varargs...)
}

// DuplicateDimension mocks base method.
func (m *MockDimensionServiceClient) DuplicateDimension(ctx context.Context, in *pb.DuplicateDimensionRequest, opts ...grpc.CallOption) (*pb.Dimension, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DuplicateDimension", varargs...)
	ret0, _ := ret[0].(*pb.Dimension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DuplicateDimension indicates an expected call of DuplicateDimension.
func (mr *MockDimensionServiceClientMockRecorder) DuplicateDimension(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DuplicateDimension", reflect.TypeOf((*MockDimensionServiceClient)(nil).DuplicateDimension), varargs...)
}

// EditDimension mocks base method.
func (m *MockDimensionServiceClient) EditDimension(ctx context.Context, in *pb.EditDimensionRequest, opts ...grpc.CallOption) (*pb.Dimension, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EditDimension", varargs...)
	ret0, _ := ret[0].(*pb.Dimension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditDimension indicates an expected call of EditDimension.
func (mr *MockDimensionServiceClientMockRecorder) EditDimension(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditDimension", reflect.TypeOf((*MockDimensionServiceClient)(nil).EditDimension), varargs...)
}

// GetDimension mocks base method.
func (m *MockDimensionServiceClient) GetDimension(ctx context.Context, in *pb0.TargetId, opts ...grpc.CallOption) (*pb.Dimension, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDimension", varargs...)
	ret0, _ := ret[0].(*pb.Dimension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimension indicates an expected call of GetDimension.
func (mr *MockDimensionServiceClientMockRecorder) GetDimension(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimension", reflect.TypeOf((*MockDimensionServiceClient)(nil).GetDimension), varargs...)
}

// GetDimensions mocks base method.
func (m *MockDimensionServiceClient) GetDimensions(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*pb.Dimensions, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetDimensions", varargs...)
	ret0, _ := ret[0].(*pb.Dimensions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimensions indicates an expected call of GetDimensions.
func (mr *MockDimensionServiceClientMockRecorder) GetDimensions(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensions", reflect.TypeOf((*MockDimensionServiceClient)(nil).GetDimensions), varargs...)
}

// MockDimensionServiceServer is a mock of DimensionServiceServer interface.
type MockDimensionServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockDimensionServiceServerMockRecorder
	isgomock struct{}
}

// MockDimensionServiceServerMockRecorder is the mock recorder for MockDimensionServiceServer.
type MockDimensionServiceServerMockRecorder struct {
	mock *MockDimensionServiceServer
}

// NewMockDimensionServiceServer creates a new mock instance.
func NewMockDimensionServiceServer(ctrl *gomock.Controller) *MockDimensionServiceServer {
	mock := &MockDimensionServiceServer{ctrl: ctrl}
	mock.recorder = &MockDimensionServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDimensionServiceServer) EXPECT() *MockDimensionServiceServerMockRecorder {
	return m.recorder
}

// CreateDimension mocks base method.
func (m *MockDimensionServiceServer) CreateDimension(arg0 context.Context, arg1 *pb.CreateDimensionRequest) (*pb.Dimension, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateDimension", arg0, arg1)
	ret0, _ := ret[0].(*pb.Dimension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateDimension indicates an expected call of CreateDimension.
func (mr *MockDimensionServiceServerMockRecorder) CreateDimension(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateDimension", reflect.TypeOf((*MockDimensionServiceServer)(nil).CreateDimension), arg0, arg1)
}

// DeleteDimension mocks base method.
func (m *MockDimensionServiceServer) DeleteDimension(arg0 context.Context, arg1 *pb0.TargetId) (*emptypb.Empty, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDimension", arg0, arg1)
	ret0, _ := ret[0].(*emptypb.Empty)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteDimension indicates an expected call of DeleteDimension.
func (mr *MockDimensionServiceServerMockRecorder) DeleteDimension(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDimension", reflect.TypeOf((*MockDimensionServiceServer)(nil).DeleteDimension), arg0, arg1)
}

// DuplicateDimension mocks base method.
func (m *MockDimensionServiceServer) DuplicateDimension(arg0 context.Context, arg1 *pb.DuplicateDimensionRequest) (*pb.Dimension, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DuplicateDimension", arg0, arg1)
	ret0, _ := ret[0].(*pb.Dimension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DuplicateDimension indicates an expected call of DuplicateDimension.
func (mr *MockDimensionServiceServerMockRecorder) DuplicateDimension(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DuplicateDimension", reflect.TypeOf((*MockDimensionServiceServer)(nil).DuplicateDimension), arg0, arg1)
}

// EditDimension mocks base method.
func (m *MockDimensionServiceServer) EditDimension(arg0 context.Context, arg1 *pb.EditDimensionRequest) (*pb.Dimension, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditDimension", arg0, arg1)
	ret0, _ := ret[0].(*pb.Dimension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EditDimension indicates an expected call of EditDimension.
func (mr *MockDimensionServiceServerMockRecorder) EditDimension(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditDimension", reflect.TypeOf((*MockDimensionServiceServer)(nil).EditDimension), arg0, arg1)
}

// GetDimension mocks base method.
func (m *MockDimensionServiceServer) GetDimension(arg0 context.Context, arg1 *pb0.TargetId) (*pb.Dimension, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimension", arg0, arg1)
	ret0, _ := ret[0].(*pb.Dimension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimension indicates an expected call of GetDimension.
func (mr *MockDimensionServiceServerMockRecorder) GetDimension(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimension", reflect.TypeOf((*MockDimensionServiceServer)(nil).GetDimension), arg0, arg1)
}

// GetDimensions mocks base method.
func (m *MockDimensionServiceServer) GetDimensions(arg0 context.Context, arg1 *emptypb.Empty) (*pb.Dimensions, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDimensions", arg0, arg1)
	ret0, _ := ret[0].(*pb.Dimensions)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDimensions indicates an expected call of GetDimensions.
func (mr *MockDimensionServiceServerMockRecorder) GetDimensions(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDimensions", reflect.TypeOf((*MockDimensionServiceServer)(nil).GetDimensions), arg0, arg1)
}

// mustEmbedUnimplementedDimensionServiceServer mocks base method.
func (m *MockDimensionServiceServer) mustEmbedUnimplementedDimensionServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedDimensionServiceServer")
}

// mustEmbedUnimplementedDimensionServiceServer indicates an expected call of mustEmbedUnimplementedDimensionServiceServer.
func (mr *MockDimensionServiceServerMockRecorder) mustEmbedUnimplementedDimensionServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedDimensionServiceServer", reflect.TypeOf((*MockDimensionServiceServer)(nil).mustEmbedUnimplementedDimensionServiceServer))
}

// MockUnsafeDimensionServiceServer is a mock of UnsafeDimensionServiceServer interface.
type MockUnsafeDimensionServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeDimensionServiceServerMockRecorder
	isgomock struct{}
}

// MockUnsafeDimensionServiceServerMockRecorder is the mock recorder for MockUnsafeDimensionServiceServer.
type MockUnsafeDimensionServiceServerMockRecorder struct {
	mock *MockUnsafeDimensionServiceServer
}

// NewMockUnsafeDimensionServiceServer creates a new mock instance.
func NewMockUnsafeDimensionServiceServer(ctrl *gomock.Controller) *MockUnsafeDimensionServiceServer {
	mock := &MockUnsafeDimensionServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeDimensionServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeDimensionServiceServer) EXPECT() *MockUnsafeDimensionServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedDimensionServiceServer mocks base method.
func (m *MockUnsafeDimensionServiceServer) mustEmbedUnimplementedDimensionServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedDimensionServiceServer")
}

// mustEmbedUnimplementedDimensionServiceServer indicates an expected call of mustEmbedUnimplementedDimensionServiceServer.
func (mr *MockUnsafeDimensionServiceServerMockRecorder) mustEmbedUnimplementedDimensionServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedDimensionServiceServer", reflect.TypeOf((*MockUnsafeDimensionServiceServer)(nil).mustEmbedUnimplementedDimensionServiceServer))
}
