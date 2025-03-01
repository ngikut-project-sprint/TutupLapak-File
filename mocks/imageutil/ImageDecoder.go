// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import (
	image "image"

	io "io"

	mock "github.com/stretchr/testify/mock"
)

// ImageDecoder is an autogenerated mock type for the ImageDecoder type
type ImageDecoder struct {
	mock.Mock
}

type ImageDecoder_Expecter struct {
	mock *mock.Mock
}

func (_m *ImageDecoder) EXPECT() *ImageDecoder_Expecter {
	return &ImageDecoder_Expecter{mock: &_m.Mock}
}

// Decode provides a mock function with given fields: r
func (_m *ImageDecoder) Decode(r io.Reader) (image.Image, error) {
	ret := _m.Called(r)

	if len(ret) == 0 {
		panic("no return value specified for Decode")
	}

	var r0 image.Image
	var r1 error
	if rf, ok := ret.Get(0).(func(io.Reader) (image.Image, error)); ok {
		return rf(r)
	}
	if rf, ok := ret.Get(0).(func(io.Reader) image.Image); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(image.Image)
		}
	}

	if rf, ok := ret.Get(1).(func(io.Reader) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImageDecoder_Decode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Decode'
type ImageDecoder_Decode_Call struct {
	*mock.Call
}

// Decode is a helper method to define mock.On call
//   - r io.Reader
func (_e *ImageDecoder_Expecter) Decode(r interface{}) *ImageDecoder_Decode_Call {
	return &ImageDecoder_Decode_Call{Call: _e.mock.On("Decode", r)}
}

func (_c *ImageDecoder_Decode_Call) Run(run func(r io.Reader)) *ImageDecoder_Decode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(io.Reader))
	})
	return _c
}

func (_c *ImageDecoder_Decode_Call) Return(_a0 image.Image, _a1 error) *ImageDecoder_Decode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ImageDecoder_Decode_Call) RunAndReturn(run func(io.Reader) (image.Image, error)) *ImageDecoder_Decode_Call {
	_c.Call.Return(run)
	return _c
}

// NewImageDecoder creates a new instance of ImageDecoder. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewImageDecoder(t interface {
	mock.TestingT
	Cleanup(func())
}) *ImageDecoder {
	mock := &ImageDecoder{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
