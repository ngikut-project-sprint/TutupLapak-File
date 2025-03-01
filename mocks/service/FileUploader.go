// Code generated by mockery v2.51.1. DO NOT EDIT.

package mocks

import (
	context "context"

	manager "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	mock "github.com/stretchr/testify/mock"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

// FileUploader is an autogenerated mock type for the FileUploader type
type FileUploader struct {
	mock.Mock
}

type FileUploader_Expecter struct {
	mock *mock.Mock
}

func (_m *FileUploader) EXPECT() *FileUploader_Expecter {
	return &FileUploader_Expecter{mock: &_m.Mock}
}

// Upload provides a mock function with given fields: ctx, input, opts
func (_m *FileUploader) Upload(ctx context.Context, input *s3.PutObjectInput, opts ...func(*manager.Uploader)) (*manager.UploadOutput, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, input)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Upload")
	}

	var r0 *manager.UploadOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *s3.PutObjectInput, ...func(*manager.Uploader)) (*manager.UploadOutput, error)); ok {
		return rf(ctx, input, opts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *s3.PutObjectInput, ...func(*manager.Uploader)) *manager.UploadOutput); ok {
		r0 = rf(ctx, input, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*manager.UploadOutput)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *s3.PutObjectInput, ...func(*manager.Uploader)) error); ok {
		r1 = rf(ctx, input, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FileUploader_Upload_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Upload'
type FileUploader_Upload_Call struct {
	*mock.Call
}

// Upload is a helper method to define mock.On call
//   - ctx context.Context
//   - input *s3.PutObjectInput
//   - opts ...func(*manager.Uploader)
func (_e *FileUploader_Expecter) Upload(ctx interface{}, input interface{}, opts ...interface{}) *FileUploader_Upload_Call {
	return &FileUploader_Upload_Call{Call: _e.mock.On("Upload",
		append([]interface{}{ctx, input}, opts...)...)}
}

func (_c *FileUploader_Upload_Call) Run(run func(ctx context.Context, input *s3.PutObjectInput, opts ...func(*manager.Uploader))) *FileUploader_Upload_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]func(*manager.Uploader), len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(func(*manager.Uploader))
			}
		}
		run(args[0].(context.Context), args[1].(*s3.PutObjectInput), variadicArgs...)
	})
	return _c
}

func (_c *FileUploader_Upload_Call) Return(_a0 *manager.UploadOutput, _a1 error) *FileUploader_Upload_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *FileUploader_Upload_Call) RunAndReturn(run func(context.Context, *s3.PutObjectInput, ...func(*manager.Uploader)) (*manager.UploadOutput, error)) *FileUploader_Upload_Call {
	_c.Call.Return(run)
	return _c
}

// NewFileUploader creates a new instance of FileUploader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFileUploader(t interface {
	mock.TestingT
	Cleanup(func())
}) *FileUploader {
	mock := &FileUploader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
