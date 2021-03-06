// Code generated by MockGen. DO NOT EDIT.
// Source: /Users/yigalziskand/dev/go/src/github.com/ziskandyigal/stringculator/calculator/calculator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCalculator is a mock of Calculator interface.
type MockCalculator struct {
	ctrl     *gomock.Controller
	recorder *MockCalculatorMockRecorder
}

// MockCalculatorMockRecorder is the mock recorder for MockCalculator.
type MockCalculatorMockRecorder struct {
	mock *MockCalculator
}

// NewMockCalculator creates a new mock instance.
func NewMockCalculator(ctrl *gomock.Controller) *MockCalculator {
	mock := &MockCalculator{ctrl: ctrl}
	mock.recorder = &MockCalculatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCalculator) EXPECT() *MockCalculatorMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockCalculator) Add(num1, num2 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", num1, num2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Add indicates an expected call of Add.
func (mr *MockCalculatorMockRecorder) Add(num1, num2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockCalculator)(nil).Add), num1, num2)
}

// Div mocks base method.
func (m *MockCalculator) Div(num1, num2 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Div", num1, num2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Div indicates an expected call of Div.
func (mr *MockCalculatorMockRecorder) Div(num1, num2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Div", reflect.TypeOf((*MockCalculator)(nil).Div), num1, num2)
}

// Sub mocks base method.
func (m *MockCalculator) Sub(num1, num2 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sub", num1, num2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sub indicates an expected call of Sub.
func (mr *MockCalculatorMockRecorder) Sub(num1, num2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sub", reflect.TypeOf((*MockCalculator)(nil).Sub), num1, num2)
}

// Sum mocks base method.
func (m *MockCalculator) Sum(num1, num2 int) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sum", num1, num2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sum indicates an expected call of Sum.
func (mr *MockCalculatorMockRecorder) Sum(num1, num2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sum", reflect.TypeOf((*MockCalculator)(nil).Sum), num1, num2)
}
