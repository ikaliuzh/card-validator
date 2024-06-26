// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ikaliuzh/card-validator/pkg/validator/expirydate (interfaces: DateProvider)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockDateProvider is a mock of DateProvider interface.
type MockDateProvider struct {
	ctrl     *gomock.Controller
	recorder *MockDateProviderMockRecorder
}

// MockDateProviderMockRecorder is the mock recorder for MockDateProvider.
type MockDateProviderMockRecorder struct {
	mock *MockDateProvider
}

// NewMockDateProvider creates a new mock instance.
func NewMockDateProvider(ctrl *gomock.Controller) *MockDateProvider {
	mock := &MockDateProvider{ctrl: ctrl}
	mock.recorder = &MockDateProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDateProvider) EXPECT() *MockDateProviderMockRecorder {
	return m.recorder
}

// CurrentMonth mocks base method.
func (m *MockDateProvider) CurrentMonth() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentMonth")
	ret0, _ := ret[0].(int)
	return ret0
}

// CurrentMonth indicates an expected call of CurrentMonth.
func (mr *MockDateProviderMockRecorder) CurrentMonth() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentMonth", reflect.TypeOf((*MockDateProvider)(nil).CurrentMonth))
}

// CurrentYear mocks base method.
func (m *MockDateProvider) CurrentYear() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentYear")
	ret0, _ := ret[0].(int)
	return ret0
}

// CurrentYear indicates an expected call of CurrentYear.
func (mr *MockDateProviderMockRecorder) CurrentYear() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentYear", reflect.TypeOf((*MockDateProvider)(nil).CurrentYear))
}
