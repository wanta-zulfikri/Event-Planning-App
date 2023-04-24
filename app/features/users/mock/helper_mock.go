package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

type Helper interface {
	HashedPassword(password string) ([]byte, error)
}

type MockHelper struct {
	ctrl     *gomock.Controller
	recorder *MockHelperMockRecorder
}

func (m *MockHelper) HashedPassword(password string) ([]byte, error) {
	ret := m.ctrl.Call(m, "HashedPassword", password)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func NewMockHelper(ctrl *gomock.Controller) *MockHelper {
	mock := &MockHelper{ctrl: ctrl}
	mock.recorder = &MockHelperMockRecorder{mock}
	return mock
}

type MockHelperMockRecorder struct {
	mock *MockHelper
}

func (mr *MockHelperMockRecorder) HashedPassword(password interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashedPassword", reflect.TypeOf((*Helper)(nil)).Elem(), password)
}
