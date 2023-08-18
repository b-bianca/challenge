// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/b-bianca/melichallenge/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// UserUseCase is an autogenerated mock type for the UserUseCase type
type UserUseCase struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, u
func (_m *UserUseCase) CreateUser(ctx context.Context, u *entity.User) (*entity.User, error) {
	ret := _m.Called(ctx, u)

	var r0 *entity.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) (*entity.User, error)); ok {
		return rf(ctx, u)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) *entity.User); ok {
		r0 = rf(ctx, u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.User) error); ok {
		r1 = rf(ctx, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// OptoutUser provides a mock function with given fields: ctx, u
func (_m *UserUseCase) OptoutUser(ctx context.Context, u *entity.User) error {
	ret := _m.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.User) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserUseCase creates a new instance of UserUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUseCase {
	mock := &UserUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}