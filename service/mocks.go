// Code generated by MockGen. DO NOT EDIT.
// Source: service/config.go
//
// Generated by this command:
//
//	mockgen -source=service/config.go -destination=service/mocks.go -package=service ConfigService
//

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"

	model "github.com/SawitProRecruitment/JuniorBackendEngineering/model"
	gomock "go.uber.org/mock/gomock"
)

// MockConfigService is a mock of ConfigService interface.
type MockConfigService struct {
	ctrl     *gomock.Controller
	recorder *MockConfigServiceMockRecorder
}

// MockConfigServiceMockRecorder is the mock recorder for MockConfigService.
type MockConfigServiceMockRecorder struct {
	mock *MockConfigService
}

// NewMockConfigService creates a new mock instance.
func NewMockConfigService(ctrl *gomock.Controller) *MockConfigService {
	mock := &MockConfigService{ctrl: ctrl}
	mock.recorder = &MockConfigServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConfigService) EXPECT() *MockConfigServiceMockRecorder {
	return m.recorder
}

// GetFieldDimensions mocks base method.
func (m *MockConfigService) GetFieldDimensions() (model.Field, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFieldDimensions")
	ret0, _ := ret[0].(model.Field)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFieldDimensions indicates an expected call of GetFieldDimensions.
func (mr *MockConfigServiceMockRecorder) GetFieldDimensions() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFieldDimensions", reflect.TypeOf((*MockConfigService)(nil).GetFieldDimensions))
}

// GetTrees mocks base method.
func (m *MockConfigService) GetTrees(NumberOfTrees int) ([]model.Tree, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTrees", NumberOfTrees)
	ret0, _ := ret[0].([]model.Tree)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTrees indicates an expected call of GetTrees.
func (mr *MockConfigServiceMockRecorder) GetTrees(NumberOfTrees any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTrees", reflect.TypeOf((*MockConfigService)(nil).GetTrees), NumberOfTrees)
}

// MockScanner is a mock of Scanner interface.
type MockScanner struct {
	ctrl     *gomock.Controller
	recorder *MockScannerMockRecorder
}

// MockScannerMockRecorder is the mock recorder for MockScanner.
type MockScannerMockRecorder struct {
	mock *MockScanner
}

// NewMockScanner creates a new mock instance.
func NewMockScanner(ctrl *gomock.Controller) *MockScanner {
	mock := &MockScanner{ctrl: ctrl}
	mock.recorder = &MockScannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScanner) EXPECT() *MockScannerMockRecorder {
	return m.recorder
}

// Scan mocks base method.
func (m *MockScanner) Scan() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockScannerMockRecorder) Scan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockScanner)(nil).Scan))
}

// Text mocks base method.
func (m *MockScanner) Text() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Text")
	ret0, _ := ret[0].(string)
	return ret0
}

// Text indicates an expected call of Text.
func (mr *MockScannerMockRecorder) Text() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Text", reflect.TypeOf((*MockScanner)(nil).Text))
}
