package controller_test

import (
	"bytes"
	"context"
	"errors"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/ngikut-project-sprint/TutupLapak-File/internal/controller"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/model"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/utils/validation"
	mocksSvc "github.com/ngikut-project-sprint/TutupLapak-File/mocks/service"
	mocksValidation "github.com/ngikut-project-sprint/TutupLapak-File/mocks/validation"
)

func TestFileController_Post_Success(t *testing.T) {
	e := echo.New()

	file := model.File{
		FileURI:          "https://example.com/file.jpg",
		FileThumbnailURI: "https://example.com/thumbnail.jpg",
	}

	maxFileSize := int64(102400)

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	part, _ := writer.CreateFormFile("file", "test.jpg")
	_, _ = part.Write([]byte("mock file data"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/", &b)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockValidator := new(mocksValidation.FileValidator)
	mockValidator.On("ValidateFile", mock.Anything, maxFileSize).Return(nil)

	mockService := new(mocksSvc.FileService)
	mockService.On("UploadFile", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		ch := args.Get(3).(chan model.Completion)
		ch <- model.Completion{FileURL: file.FileURI, Error: nil}
	}).Once()

	mockService.On("GenerateThumbnail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		ch := args.Get(3).(chan model.Completion)
		ch <- model.Completion{FileURL: file.FileThumbnailURI, Error: nil}
	}).Once()

	mockService.On("AddFile", mock.Anything, file.FileURI, file.FileThumbnailURI).Return(uint(123), nil).Once()

	fc := controller.NewFileController(
		mockService,
		maxFileSize,
		mockValidator.ValidateFile,
		2*time.Second,
	)

	if assert.NoError(t, fc.Post(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.True(t, strings.Contains(rec.Body.String(), "123"))
		assert.True(t, strings.Contains(rec.Body.String(), file.FileURI))
		assert.True(t, strings.Contains(rec.Body.String(), file.FileThumbnailURI))
	}
}

func TestFileController_Post_InvalidFileType(t *testing.T) {
	e := echo.New()

	maxFileSize := int64(102400)

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	part, _ := writer.CreateFormFile("file", "test.jpg")
	_, _ = part.Write([]byte("mock file data"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/", &b)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockValidator := new(mocksValidation.FileValidator)
	mockValidator.On("ValidateFile", mock.Anything, maxFileSize).Return(validation.InvalidFileType)

	mockService := new(mocksSvc.FileService)

	fc := controller.NewFileController(
		mockService,
		maxFileSize,
		mockValidator.ValidateFile,
		2*time.Second,
	)

	if assert.NoError(t, fc.Post(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.True(t, strings.Contains(rec.Body.String(), validation.InvalidFileType.Error()))
	}
}

func TestFileController_Post_InvalidFileSize(t *testing.T) {
	e := echo.New()

	maxFileSize := int64(102400)

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	part, _ := writer.CreateFormFile("file", "test.jpg")
	_, _ = part.Write([]byte("mock file data"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/", &b)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockValidator := new(mocksValidation.FileValidator)
	mockValidator.On("ValidateFile", mock.Anything, maxFileSize).Return(validation.InvalidFileSize)

	mockService := new(mocksSvc.FileService)

	fc := controller.NewFileController(
		mockService,
		maxFileSize,
		mockValidator.ValidateFile,
		2*time.Second,
	)

	if assert.NoError(t, fc.Post(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.True(t, strings.Contains(rec.Body.String(), validation.InvalidFileSize.Error()))
	}
}

func TestFileController_Post_UploadFileError(t *testing.T) {
	e := echo.New()

	file := model.File{
		FileURI:          "https://example.com/file.jpg",
		FileThumbnailURI: "https://example.com/thumbnail.jpg",
	}

	maxFileSize := int64(102400)

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	part, _ := writer.CreateFormFile("file", "test.jpg")
	_, _ = part.Write([]byte("mock file data"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/", &b)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockValidator := new(mocksValidation.FileValidator)
	mockValidator.On("ValidateFile", mock.Anything, maxFileSize).Return(nil)

	mockService := new(mocksSvc.FileService)
	mockService.On("UploadFile", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		ch := args.Get(3).(chan model.Completion)
		ch <- model.Completion{FileURL: "", Error: errors.New("failed uploading file")}
	}).Once()

	mockService.On("GenerateThumbnail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		ch := args.Get(3).(chan model.Completion)
		ch <- model.Completion{FileURL: file.FileThumbnailURI, Error: nil}
	}).Once()

	fc := controller.NewFileController(
		mockService,
		maxFileSize,
		mockValidator.ValidateFile,
		2*time.Second,
	)

	if assert.NoError(t, fc.Post(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.True(t, strings.Contains(rec.Body.String(), "Server error"))
	}
}

func TestFileController_Post_GenerateThumbnailError(t *testing.T) {
	e := echo.New()

	file := model.File{
		FileURI:          "https://example.com/file.jpg",
		FileThumbnailURI: "https://example.com/thumbnail.jpg",
	}

	maxFileSize := int64(102400)

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	part, _ := writer.CreateFormFile("file", "test.jpg")
	_, _ = part.Write([]byte("mock file data"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/", &b)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockValidator := new(mocksValidation.FileValidator)
	mockValidator.On("ValidateFile", mock.Anything, maxFileSize).Return(nil)

	mockService := new(mocksSvc.FileService)
	mockService.On("UploadFile", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		ch := args.Get(3).(chan model.Completion)
		ch <- model.Completion{FileURL: file.FileURI, Error: nil}
	}).Once()

	mockService.On("GenerateThumbnail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		ch := args.Get(3).(chan model.Completion)
		ch <- model.Completion{FileURL: "", Error: errors.New("failed uploading thumbnail")}
	}).Once()

	fc := controller.NewFileController(
		mockService,
		maxFileSize,
		mockValidator.ValidateFile,
		2*time.Second,
	)

	if assert.NoError(t, fc.Post(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.True(t, strings.Contains(rec.Body.String(), "Server error"))
	}
}

func TestFileController_Post_AddFileError(t *testing.T) {
	e := echo.New()

	file := model.File{
		FileURI:          "https://example.com/file.jpg",
		FileThumbnailURI: "https://example.com/thumbnail.jpg",
	}

	maxFileSize := int64(102400)

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	part, _ := writer.CreateFormFile("file", "test.jpg")
	_, _ = part.Write([]byte("mock file data"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/", &b)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockValidator := new(mocksValidation.FileValidator)
	mockValidator.On("ValidateFile", mock.Anything, maxFileSize).Return(nil)

	mockService := new(mocksSvc.FileService)
	mockService.On("UploadFile", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		ch := args.Get(3).(chan model.Completion)
		ch <- model.Completion{FileURL: file.FileURI, Error: nil}
	}).Once()

	mockService.On("GenerateThumbnail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		ch := args.Get(3).(chan model.Completion)
		ch <- model.Completion{FileURL: file.FileThumbnailURI, Error: nil}
	}).Once()

	mockService.On("AddFile", mock.Anything, file.FileURI, file.FileThumbnailURI).Return(uint(0), errors.New("failed storing file to database")).Once()

	fc := controller.NewFileController(
		mockService,
		maxFileSize,
		mockValidator.ValidateFile,
		2*time.Second,
	)

	if assert.NoError(t, fc.Post(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.True(t, strings.Contains(rec.Body.String(), "Server error"))
	}
}

func TestFileController_Post_TimeoutError(t *testing.T) {
	e := echo.New()

	maxFileSize := int64(102400)

	var b bytes.Buffer
	writer := multipart.NewWriter(&b)
	part, _ := writer.CreateFormFile("file", "test.jpg")
	_, _ = part.Write([]byte("mock file data"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/", &b)
	req.Header.Set(echo.HeaderContentType, writer.FormDataContentType())
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockValidator := new(mocksValidation.FileValidator)
	mockValidator.On("ValidateFile", mock.Anything, maxFileSize).Return(nil)

	mockService := new(mocksSvc.FileService)
	mockService.On("UploadFile", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		ch := args.Get(3).(chan model.Completion)
		ch <- model.Completion{FileURL: "", Error: context.DeadlineExceeded}
	}).Once()

	mockService.On("GenerateThumbnail", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		ch := args.Get(3).(chan model.Completion)
		ch <- model.Completion{FileURL: "", Error: context.DeadlineExceeded}
	}).Once()

	fc := controller.NewFileController(
		mockService,
		maxFileSize,
		mockValidator.ValidateFile,
		time.Millisecond,
	)

	if assert.NoError(t, fc.Post(c)) {
		assert.Equal(t, http.StatusInternalServerError, rec.Code)
		assert.True(t, strings.Contains(rec.Body.String(), "Server error"))
	}
}
