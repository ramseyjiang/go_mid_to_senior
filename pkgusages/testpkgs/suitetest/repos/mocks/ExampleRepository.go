// Code generated by mockery v2.15.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// ExampleRepository is an autogenerated mock type for the ExampleRepository type
type ExampleRepository struct {
	mock.Mock
}

// GetExampleTaxValue provides a mock function with given fields: firstParm, secondParam
func (_m *ExampleRepository) GetExampleTaxValue(firstParm int, secondParam int) int {
	ret := _m.Called(firstParm, secondParam)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(firstParm, secondParam)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewExampleRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewExampleRepository creates a new instance of ExampleRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewExampleRepository(t mockConstructorTestingTNewExampleRepository) *ExampleRepository {
	mock := &ExampleRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
