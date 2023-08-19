// Code generated by mockery v2.32.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "github.com/b-bianca/melichallenge/notify-api/internal/domain/entity"
	mock "github.com/stretchr/testify/mock"
)

// NotifyUseCase is an autogenerated mock type for the NotifyUseCase type
type NotifyUseCase struct {
	mock.Mock
}

// CreateNotify provides a mock function with given fields: ctx, u
func (_m *NotifyUseCase) CreateNotify(ctx context.Context, u *entity.Notification) (*entity.Notification, error) {
	ret := _m.Called(ctx, u)

	var r0 *entity.Notification
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Notification) (*entity.Notification, error)); ok {
		return rf(ctx, u)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Notification) *entity.Notification); ok {
		r0 = rf(ctx, u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Notification)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Notification) error); ok {
		r1 = rf(ctx, u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchMessage provides a mock function with given fields: ctx
func (_m *NotifyUseCase) FetchMessage(ctx context.Context) (*entity.MessageList, error) {
	ret := _m.Called(ctx)

	var r0 *entity.MessageList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*entity.MessageList, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *entity.MessageList); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.MessageList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchNotify provides a mock function with given fields: ctx
func (_m *NotifyUseCase) FetchNotify(ctx context.Context) (*entity.NotificationList, error) {
	ret := _m.Called(ctx)

	var r0 *entity.NotificationList
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*entity.NotificationList, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *entity.NotificationList); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.NotificationList)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SendMessage provides a mock function with given fields: ctx, m
func (_m *NotifyUseCase) SendMessage(ctx context.Context, m *entity.Message) (*entity.Message, error) {
	ret := _m.Called(ctx, m)

	var r0 *entity.Message
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Message) (*entity.Message, error)); ok {
		return rf(ctx, m)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *entity.Message) *entity.Message); ok {
		r0 = rf(ctx, m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Message)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *entity.Message) error); ok {
		r1 = rf(ctx, m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewNotifyUseCase creates a new instance of NotifyUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewNotifyUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *NotifyUseCase {
	mock := &NotifyUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
