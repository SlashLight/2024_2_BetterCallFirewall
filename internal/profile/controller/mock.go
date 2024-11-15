// Code generated by MockGen. DO NOT EDIT.
// Source: controller.go

// Package controller is a generated GoMock package.
package controller

import (
	context "context"
	http "net/http"
	reflect "reflect"

	models "github.com/2024_2_BetterCallFirewall/internal/models"
	gomock "github.com/golang/mock/gomock"
)

// MockResponder is a mock of Responder interface.
type MockResponder struct {
	ctrl     *gomock.Controller
	recorder *MockResponderMockRecorder
}

// MockResponderMockRecorder is the mock recorder for MockResponder.
type MockResponderMockRecorder struct {
	mock *MockResponder
}

// NewMockResponder creates a new mock instance.
func NewMockResponder(ctrl *gomock.Controller) *MockResponder {
	mock := &MockResponder{ctrl: ctrl}
	mock.recorder = &MockResponderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResponder) EXPECT() *MockResponderMockRecorder {
	return m.recorder
}

// ErrorBadRequest mocks base method.
func (m *MockResponder) ErrorBadRequest(w http.ResponseWriter, err error, requestID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ErrorBadRequest", w, err, requestID)
}

// ErrorBadRequest indicates an expected call of ErrorBadRequest.
func (mr *MockResponderMockRecorder) ErrorBadRequest(w, err, requestID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrorBadRequest", reflect.TypeOf((*MockResponder)(nil).ErrorBadRequest), w, err, requestID)
}

// ErrorInternal mocks base method.
func (m *MockResponder) ErrorInternal(w http.ResponseWriter, err error, requestID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ErrorInternal", w, err, requestID)
}

// ErrorInternal indicates an expected call of ErrorInternal.
func (mr *MockResponderMockRecorder) ErrorInternal(w, err, requestID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ErrorInternal", reflect.TypeOf((*MockResponder)(nil).ErrorInternal), w, err, requestID)
}

// LogError mocks base method.
func (m *MockResponder) LogError(err error, requestID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LogError", err, requestID)
}

// LogError indicates an expected call of LogError.
func (mr *MockResponderMockRecorder) LogError(err, requestID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogError", reflect.TypeOf((*MockResponder)(nil).LogError), err, requestID)
}

// OutputJSON mocks base method.
func (m *MockResponder) OutputJSON(w http.ResponseWriter, data any, requestID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OutputJSON", w, data, requestID)
}

// OutputJSON indicates an expected call of OutputJSON.
func (mr *MockResponderMockRecorder) OutputJSON(w, data, requestID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutputJSON", reflect.TypeOf((*MockResponder)(nil).OutputJSON), w, data, requestID)
}

// OutputNoMoreContentJSON mocks base method.
func (m *MockResponder) OutputNoMoreContentJSON(w http.ResponseWriter, requestID string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OutputNoMoreContentJSON", w, requestID)
}

// OutputNoMoreContentJSON indicates an expected call of OutputNoMoreContentJSON.
func (mr *MockResponderMockRecorder) OutputNoMoreContentJSON(w, requestID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OutputNoMoreContentJSON", reflect.TypeOf((*MockResponder)(nil).OutputNoMoreContentJSON), w, requestID)
}

// MockProfileUsecase is a mock of ProfileUsecase interface.
type MockProfileUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockProfileUsecaseMockRecorder
}

// MockProfileUsecaseMockRecorder is the mock recorder for MockProfileUsecase.
type MockProfileUsecaseMockRecorder struct {
	mock *MockProfileUsecase
}

// NewMockProfileUsecase creates a new mock instance.
func NewMockProfileUsecase(ctrl *gomock.Controller) *MockProfileUsecase {
	mock := &MockProfileUsecase{ctrl: ctrl}
	mock.recorder = &MockProfileUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProfileUsecase) EXPECT() *MockProfileUsecaseMockRecorder {
	return m.recorder
}

// AcceptFriendReq mocks base method.
func (m *MockProfileUsecase) AcceptFriendReq(who, whose uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptFriendReq", who, whose)
	ret0, _ := ret[0].(error)
	return ret0
}

// AcceptFriendReq indicates an expected call of AcceptFriendReq.
func (mr *MockProfileUsecaseMockRecorder) AcceptFriendReq(who, whose interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptFriendReq", reflect.TypeOf((*MockProfileUsecase)(nil).AcceptFriendReq), who, whose)
}

// DeleteProfile mocks base method.
func (m *MockProfileUsecase) DeleteProfile(arg0 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProfile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProfile indicates an expected call of DeleteProfile.
func (mr *MockProfileUsecaseMockRecorder) DeleteProfile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProfile", reflect.TypeOf((*MockProfileUsecase)(nil).DeleteProfile), arg0)
}

// GetAll mocks base method.
func (m *MockProfileUsecase) GetAll(ctx context.Context, self, lastId uint32) ([]*models.ShortProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, self, lastId)
	ret0, _ := ret[0].([]*models.ShortProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockProfileUsecaseMockRecorder) GetAll(ctx, self, lastId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockProfileUsecase)(nil).GetAll), ctx, self, lastId)
}

// GetAllFriends mocks base method.
func (m *MockProfileUsecase) GetAllFriends(ctx context.Context, id, lastId uint32) ([]*models.ShortProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFriends", ctx, id, lastId)
	ret0, _ := ret[0].([]*models.ShortProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFriends indicates an expected call of GetAllFriends.
func (mr *MockProfileUsecaseMockRecorder) GetAllFriends(ctx, id, lastId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFriends", reflect.TypeOf((*MockProfileUsecase)(nil).GetAllFriends), ctx, id, lastId)
}

// GetAllSubs mocks base method.
func (m *MockProfileUsecase) GetAllSubs(ctx context.Context, id, lastId uint32) ([]*models.ShortProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSubs", ctx, id, lastId)
	ret0, _ := ret[0].([]*models.ShortProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSubs indicates an expected call of GetAllSubs.
func (mr *MockProfileUsecaseMockRecorder) GetAllSubs(ctx, id, lastId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSubs", reflect.TypeOf((*MockProfileUsecase)(nil).GetAllSubs), ctx, id, lastId)
}

// GetAllSubscriptions mocks base method.
func (m *MockProfileUsecase) GetAllSubscriptions(ctx context.Context, id, lastId uint32) ([]*models.ShortProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllSubscriptions", ctx, id, lastId)
	ret0, _ := ret[0].([]*models.ShortProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllSubscriptions indicates an expected call of GetAllSubscriptions.
func (mr *MockProfileUsecaseMockRecorder) GetAllSubscriptions(ctx, id, lastId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllSubscriptions", reflect.TypeOf((*MockProfileUsecase)(nil).GetAllSubscriptions), ctx, id, lastId)
}

// GetCommunitySubs mocks base method.
func (m *MockProfileUsecase) GetCommunitySubs(ctx context.Context, communityID, lastID uint32) ([]*models.ShortProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommunitySubs", ctx, communityID, lastID)
	ret0, _ := ret[0].([]*models.ShortProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommunitySubs indicates an expected call of GetCommunitySubs.
func (mr *MockProfileUsecaseMockRecorder) GetCommunitySubs(ctx, communityID, lastID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommunitySubs", reflect.TypeOf((*MockProfileUsecase)(nil).GetCommunitySubs), ctx, communityID, lastID)
}

// GetHeader mocks base method.
func (m *MockProfileUsecase) GetHeader(ctx context.Context, userID uint32) (models.Header, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHeader", ctx, userID)
	ret0, _ := ret[0].(models.Header)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHeader indicates an expected call of GetHeader.
func (mr *MockProfileUsecaseMockRecorder) GetHeader(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHeader", reflect.TypeOf((*MockProfileUsecase)(nil).GetHeader), ctx, userID)
}

// GetProfileById mocks base method.
func (m *MockProfileUsecase) GetProfileById(arg0 context.Context, arg1 uint32) (*models.FullProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileById", arg0, arg1)
	ret0, _ := ret[0].(*models.FullProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileById indicates an expected call of GetProfileById.
func (mr *MockProfileUsecaseMockRecorder) GetProfileById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileById", reflect.TypeOf((*MockProfileUsecase)(nil).GetProfileById), arg0, arg1)
}

// RemoveFromFriends mocks base method.
func (m *MockProfileUsecase) RemoveFromFriends(who, whose uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromFriends", who, whose)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromFriends indicates an expected call of RemoveFromFriends.
func (mr *MockProfileUsecaseMockRecorder) RemoveFromFriends(who, whose interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromFriends", reflect.TypeOf((*MockProfileUsecase)(nil).RemoveFromFriends), who, whose)
}

// SendFriendReq mocks base method.
func (m *MockProfileUsecase) SendFriendReq(receiver, sender uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendFriendReq", receiver, sender)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendFriendReq indicates an expected call of SendFriendReq.
func (mr *MockProfileUsecaseMockRecorder) SendFriendReq(receiver, sender interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendFriendReq", reflect.TypeOf((*MockProfileUsecase)(nil).SendFriendReq), receiver, sender)
}

// Unsubscribe mocks base method.
func (m *MockProfileUsecase) Unsubscribe(who, whose uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", who, whose)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockProfileUsecaseMockRecorder) Unsubscribe(who, whose interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockProfileUsecase)(nil).Unsubscribe), who, whose)
}

// UpdateProfile mocks base method.
func (m *MockProfileUsecase) UpdateProfile(arg0 context.Context, arg1 *models.FullProfile) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProfile", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProfile indicates an expected call of UpdateProfile.
func (mr *MockProfileUsecaseMockRecorder) UpdateProfile(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProfile", reflect.TypeOf((*MockProfileUsecase)(nil).UpdateProfile), arg0, arg1)
}
