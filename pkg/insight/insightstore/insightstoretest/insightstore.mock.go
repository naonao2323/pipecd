// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/pipe-cd/pipecd/pkg/insight/insightstore (interfaces: Store)

// Package insightstoretest is a generated GoMock package.
package insightstoretest

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	insight "github.com/pipe-cd/pipecd/pkg/insight"
	model "github.com/pipe-cd/pipecd/pkg/model"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetDeployments mocks base method.
func (m *MockStore) GetDeployments(arg0 context.Context, arg1 string, arg2, arg3 int64, arg4 *model.InsightDeploymentVersion) ([]*model.InsightDeployment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDeployments", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*model.InsightDeployment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDeployments indicates an expected call of GetDeployments.
func (mr *MockStoreMockRecorder) GetDeployments(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDeployments", reflect.TypeOf((*MockStore)(nil).GetDeployments), arg0, arg1, arg2, arg3, arg4)
}

// LoadApplicationCounts mocks base method.
func (m *MockStore) LoadApplicationCounts(arg0 context.Context, arg1 string) (*insight.ApplicationCounts, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadApplicationCounts", arg0, arg1)
	ret0, _ := ret[0].(*insight.ApplicationCounts)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadApplicationCounts indicates an expected call of LoadApplicationCounts.
func (mr *MockStoreMockRecorder) LoadApplicationCounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadApplicationCounts", reflect.TypeOf((*MockStore)(nil).LoadApplicationCounts), arg0, arg1)
}

// LoadMilestone mocks base method.
func (m *MockStore) LoadMilestone(arg0 context.Context) (*insight.Milestone, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadMilestone", arg0)
	ret0, _ := ret[0].(*insight.Milestone)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadMilestone indicates an expected call of LoadMilestone.
func (mr *MockStoreMockRecorder) LoadMilestone(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadMilestone", reflect.TypeOf((*MockStore)(nil).LoadMilestone), arg0)
}

// PutApplicationCounts mocks base method.
func (m *MockStore) PutApplicationCounts(arg0 context.Context, arg1 string, arg2 insight.ApplicationCounts) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutApplicationCounts", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutApplicationCounts indicates an expected call of PutApplicationCounts.
func (mr *MockStoreMockRecorder) PutApplicationCounts(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutApplicationCounts", reflect.TypeOf((*MockStore)(nil).PutApplicationCounts), arg0, arg1, arg2)
}

// PutDeployments mocks base method.
func (m *MockStore) PutDeployments(arg0 context.Context, arg1 string, arg2 []*model.InsightDeployment, arg3 *model.InsightDeploymentVersion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutDeployments", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutDeployments indicates an expected call of PutDeployments.
func (mr *MockStoreMockRecorder) PutDeployments(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutDeployments", reflect.TypeOf((*MockStore)(nil).PutDeployments), arg0, arg1, arg2, arg3)
}

// PutMilestone mocks base method.
func (m *MockStore) PutMilestone(arg0 context.Context, arg1 *insight.Milestone) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutMilestone", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// PutMilestone indicates an expected call of PutMilestone.
func (mr *MockStoreMockRecorder) PutMilestone(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutMilestone", reflect.TypeOf((*MockStore)(nil).PutMilestone), arg0, arg1)
}