// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/pkg/store/store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	identity "github.com/aws/amazon-ecs-cli-v2/internal/pkg/aws/identity"
	route53 "github.com/aws/aws-sdk-go/service/route53"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockidentityService is a mock of identityService interface
type MockidentityService struct {
	ctrl     *gomock.Controller
	recorder *MockidentityServiceMockRecorder
}

// MockidentityServiceMockRecorder is the mock recorder for MockidentityService
type MockidentityServiceMockRecorder struct {
	mock *MockidentityService
}

// NewMockidentityService creates a new mock instance
func NewMockidentityService(ctrl *gomock.Controller) *MockidentityService {
	mock := &MockidentityService{ctrl: ctrl}
	mock.recorder = &MockidentityServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockidentityService) EXPECT() *MockidentityServiceMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockidentityService) Get() (identity.Caller, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(identity.Caller)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockidentityServiceMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockidentityService)(nil).Get))
}

// MockhostedZonesByNameLister is a mock of hostedZonesByNameLister interface
type MockhostedZonesByNameLister struct {
	ctrl     *gomock.Controller
	recorder *MockhostedZonesByNameListerMockRecorder
}

// MockhostedZonesByNameListerMockRecorder is the mock recorder for MockhostedZonesByNameLister
type MockhostedZonesByNameListerMockRecorder struct {
	mock *MockhostedZonesByNameLister
}

// NewMockhostedZonesByNameLister creates a new mock instance
func NewMockhostedZonesByNameLister(ctrl *gomock.Controller) *MockhostedZonesByNameLister {
	mock := &MockhostedZonesByNameLister{ctrl: ctrl}
	mock.recorder = &MockhostedZonesByNameListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockhostedZonesByNameLister) EXPECT() *MockhostedZonesByNameListerMockRecorder {
	return m.recorder
}

// ListHostedZonesByName mocks base method
func (m *MockhostedZonesByNameLister) ListHostedZonesByName(in *route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListHostedZonesByName", in)
	ret0, _ := ret[0].(*route53.ListHostedZonesByNameOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListHostedZonesByName indicates an expected call of ListHostedZonesByName
func (mr *MockhostedZonesByNameListerMockRecorder) ListHostedZonesByName(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListHostedZonesByName", reflect.TypeOf((*MockhostedZonesByNameLister)(nil).ListHostedZonesByName), in)
}