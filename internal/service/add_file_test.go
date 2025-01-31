package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/ngikut-project-sprint/TutupLapak-File/internal/model"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/service"
	mocksImage "github.com/ngikut-project-sprint/TutupLapak-File/mocks/imageutil"
	mocksRepo "github.com/ngikut-project-sprint/TutupLapak-File/mocks/repository"
	mocksRW "github.com/ngikut-project-sprint/TutupLapak-File/mocks/rwutil"
)

func TestFileService_AddFile_Success(t *testing.T) {
	mockRepo := new(mocksRepo.FileRepository)
	mockCfg := mockCFG()
	awsCfg := mockAWSConfig(t)
	mockReadFile := new(mocksRW.Reader)
	mockImgDecoder := new(mocksImage.ImageDecoder)
	mockImgCompressor := new(mocksImage.ImageCompressor)
	service := service.NewS3FileService(
		mockRepo,
		mockCfg,
		awsCfg,
		mockReadFile.ReadAll,
		mockImgDecoder.Decode,
		mockImgCompressor.CompressImage,
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	file := model.File{
		FileURI:          "https://example.com/file.jpg",
		FileThumbnailURI: "https://example.com/thumbnail.jpg",
	}

	mockRepo.On("AddFile", ctx, mock.AnythingOfType("model.File")).Return(uint(1), nil)

	fileID, err := service.AddFile(ctx, file.FileURI, file.FileThumbnailURI)
	assert.NoError(t, err)
	assert.Equal(t, fileID, uint(1))

	mockRepo.AssertExpectations(t)
	mockReadFile.AssertExpectations(t)
	mockImgDecoder.AssertExpectations(t)
	mockImgCompressor.AssertExpectations(t)
}

func TestFileService_AddFile_TimeoutError(t *testing.T) {
	mockRepo := new(mocksRepo.FileRepository)
	mockCfg := mockCFG()
	awsCfg := mockAWSConfig(t)
	mockReadFile := new(mocksRW.Reader)
	mockImgDecoder := new(mocksImage.ImageDecoder)
	mockImgCompressor := new(mocksImage.ImageCompressor)
	service := service.NewS3FileService(
		mockRepo,
		mockCfg,
		awsCfg,
		mockReadFile.ReadAll,
		mockImgDecoder.Decode,
		mockImgCompressor.CompressImage,
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	file := model.File{
		FileURI:          "https://example.com/file.jpg",
		FileThumbnailURI: "https://example.com/thumbnail.jpg",
	}

	mockRepo.On("AddFile", ctx, mock.AnythingOfType("model.File")).Return(uint(0), context.DeadlineExceeded)

	_, err := service.AddFile(ctx, file.FileURI, file.FileThumbnailURI)

	if err == nil {
		t.Fatalf("Expected timeout error, but got no error")
	}

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Expected context.DeadlineExceeded error, but got: %v", err)
	}

	mockRepo.AssertExpectations(t)
	mockReadFile.AssertExpectations(t)
	mockImgDecoder.AssertExpectations(t)
	mockImgCompressor.AssertExpectations(t)
}

func TestFileService_AddFile_RepoError(t *testing.T) {
	mockRepo := new(mocksRepo.FileRepository)
	mockCfg := mockCFG()
	awsCfg := mockAWSConfig(t)
	mockReadFile := new(mocksRW.Reader)
	mockImgDecoder := new(mocksImage.ImageDecoder)
	mockImgCompressor := new(mocksImage.ImageCompressor)
	service := service.NewS3FileService(
		mockRepo,
		mockCfg,
		awsCfg,
		mockReadFile.ReadAll,
		mockImgDecoder.Decode,
		mockImgCompressor.CompressImage,
	)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	file := model.File{
		FileURI:          "https://example.com/file.jpg",
		FileThumbnailURI: "https://example.com/thumbnail.jpg",
	}

	mockRepo.On("AddFile", ctx, mock.AnythingOfType("model.File")).Return(uint(0), errors.New("SQL error"))

	_, err := service.AddFile(ctx, file.FileURI, file.FileThumbnailURI)
	assert.Error(t, err)

	mockRepo.AssertExpectations(t)
	mockReadFile.AssertExpectations(t)
	mockImgDecoder.AssertExpectations(t)
	mockImgCompressor.AssertExpectations(t)
}
