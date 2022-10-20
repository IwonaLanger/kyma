// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// DirectorClient is an autogenerated mock type for the DirectorClient type
type DirectorClient struct {
	mock.Mock
}

// AssignRuntimeToFormation provides a mock function with given fields: runtimeId, formationName
func (_m *DirectorClient) AssignRuntimeToFormation(runtimeId string, formationName string) error {
	ret := _m.Called(runtimeId, formationName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(runtimeId, formationName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetConnectionToken provides a mock function with given fields: runtimeID
func (_m *DirectorClient) GetConnectionToken(runtimeID string) (string, string, error) {
	ret := _m.Called(runtimeID)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(runtimeID)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(runtimeID)
	} else {
		r1 = ret.Get(1).(string)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(runtimeID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// RegisterFormation provides a mock function with given fields: formationName
func (_m *DirectorClient) RegisterFormation(formationName string) error {
	ret := _m.Called(formationName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(formationName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterRuntime provides a mock function with given fields: runtimeName
func (_m *DirectorClient) RegisterRuntime(runtimeName string) (string, error) {
	ret := _m.Called(runtimeName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(runtimeName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(runtimeName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnregisterFormation provides a mock function with given fields: formationName
func (_m *DirectorClient) UnregisterFormation(formationName string) error {
	ret := _m.Called(formationName)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(formationName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UnregisterRuntime provides a mock function with given fields: id
func (_m *DirectorClient) UnregisterRuntime(id string) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewDirectorClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewDirectorClient creates a new instance of DirectorClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDirectorClient(t mockConstructorTestingTNewDirectorClient) *DirectorClient {
	mock := &DirectorClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}