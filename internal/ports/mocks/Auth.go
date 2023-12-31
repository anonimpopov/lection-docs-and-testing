// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	models "ru.mts.teta.tests_and_docs/internal/domain/models"
)

// Auth is an autogenerated mock type for the Auth type
type Auth struct {
	mock.Mock
}

// Info provides a mock function with given fields: ctx, login
func (_m *Auth) Info(ctx context.Context, login string) (*models.User, error) {
	ret := _m.Called(ctx, login)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(ctx, login)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, user, password
func (_m *Auth) Login(ctx context.Context, user string, password string) (models.TokenPair, error) {
	ret := _m.Called(ctx, user, password)

	var r0 models.TokenPair
	if rf, ok := ret.Get(0).(func(context.Context, string, string) models.TokenPair); ok {
		r0 = rf(ctx, user, password)
	} else {
		r0 = ret.Get(0).(models.TokenPair)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, user, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Validate provides a mock function with given fields: ctx, tokens
func (_m *Auth) Validate(ctx context.Context, tokens models.TokenPair) (string, error) {
	ret := _m.Called(ctx, tokens)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, models.TokenPair) string); ok {
		r0 = rf(ctx, tokens)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.TokenPair) error); ok {
		r1 = rf(ctx, tokens)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuth interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuth creates a new instance of Auth. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuth(t mockConstructorTestingTNewAuth) *Auth {
	mock := &Auth{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
