// Code generated by MockGen. DO NOT EDIT.
// Source: grpc_server.go

// Package profile_api is a generated GoMock package.
package profile_api

import (
	context "context"
	reflect "reflect"

	models "github.com/2024_2_BetterCallFirewall/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockprofileService is a mock of profileService interface.
type MockprofileService struct {
	ctrl     *gomock.Controller
	recorder *MockprofileServiceMockRecorder
}

// MockprofileServiceMockRecorder is the mock recorder for MockprofileService.
type MockprofileServiceMockRecorder struct {
	mock *MockprofileService
}

// NewMockprofileService creates a new mock instance.
func NewMockprofileService(ctrl *gomock.Controller) *MockprofileService {
	mock := &MockprofileService{ctrl: ctrl}
	mock.recorder = &MockprofileServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockprofileService) EXPECT() *MockprofileServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockprofileService) Create(ctx context.Context, user *models.User) (uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, user)
	ret0, _ := ret[0].(uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockprofileServiceMockRecorder) Create(ctx, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockprofileService)(nil).Create), ctx, user)
}

// GetByEmail mocks base method.
func (m *MockprofileService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByEmail", ctx, email)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByEmail indicates an expected call of GetByEmail.
func (mr *MockprofileServiceMockRecorder) GetByEmail(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByEmail", reflect.TypeOf((*MockprofileService)(nil).GetByEmail), ctx, email)
}

// GetFriendsID mocks base method.
func (m *MockprofileService) GetFriendsID(ctx context.Context, userID uint32) ([]uint32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFriendsID", ctx, userID)
	ret0, _ := ret[0].([]uint32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFriendsID indicates an expected call of GetFriendsID.
func (mr *MockprofileServiceMockRecorder) GetFriendsID(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFriendsID", reflect.TypeOf((*MockprofileService)(nil).GetFriendsID), ctx, userID)
}

// GetHeader mocks base method.
func (m *MockprofileService) GetHeader(ctx context.Context, userID uint32) (*models.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader", ctx, userID)
	ret0, _ := ret[0].(*models.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeader indicates an expected call of GetHeader.
func (mr *MockprofileServiceMockRecorder) GetHeader(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockprofileService)(nil).GetHeader), ctx, userID)
}
