package service_test

import (
	"context"
	"errors"
	"mime/multipart"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/ngikut-project-sprint/TutupLapak-File/internal/model"
	mocksImage "github.com/ngikut-project-sprint/TutupLapak-File/mocks/imageutil"
	mocksRepo "github.com/ngikut-project-sprint/TutupLapak-File/mocks/repository"
	mocksRW "github.com/ngikut-project-sprint/TutupLapak-File/mocks/rwutil"
	mocksService "github.com/ngikut-project-sprint/TutupLapak-File/mocks/service"
)

func TestFileService_UploadFile_Success(t *testing.T) {
	mockRepo := new(mocksRepo.FileRepository)
	mockCfg := mockCFG()
	mockImgDecoder := new(mocksImage.ImageDecoder)
	mockImgCompressor := new(mocksImage.ImageCompressor)

	fileHeader := &multipart.FileHeader{
		Filename: "testfile.jpg",
	}

	mockFile := new(mocksRW.FileData)
	mockFile.On("Close").Return(nil)

	mockFileOpener := new(mocksRW.FileOpener)
	mockFileOpener.On("Open").Return(mockFile, nil)

	fileBytes := []byte("fake file content")
	mockReadFile := new(mocksRW.Reader)
	mockReadFile.On("ReadAll", mock.Anything).Return(fileBytes, nil)

	mockUploader := new(mocksService.FileUploader)
	mockUploader.On("Upload", mock.Anything, mock.Anything).Return(nil, nil)

	service := mockS3FileService(
		mockRepo,
		mockCfg,
		mockUploader,
		mockReadFile.ReadAll,
		mockImgDecoder.Decode,
		mockImgCompressor.CompressImage,
	)

	completion := make(chan model.Completion)

	go service.UploadFile(context.Background(), mockFileOpener, fileHeader.Filename, completion)

	result := <-completion

	assert.NoError(t, result.Error)
	assert.Equal(t, "https://bucket.s3.ap-southeast-1.amazonaws.com/team/name/images/testfile.jpg", result.FileURL)

	mockRepo.AssertExpectations(t)
	mockImgDecoder.AssertExpectations(t)
	mockImgCompressor.AssertExpectations(t)
	mockFile.AssertExpectations(t)
	mockFileOpener.AssertExpectations(t)
	mockReadFile.AssertExpectations(t)
	mockUploader.AssertExpectations(t)
}

func TestFileService_UploadFile_FileOpenError(t *testing.T) {
	mockRepo := new(mocksRepo.FileRepository)
	mockCfg := mockCFG()
	mockImgDecoder := new(mocksImage.ImageDecoder)
	mockImgCompressor := new(mocksImage.ImageCompressor)

	fileHeader := &multipart.FileHeader{
		Filename: "testfile.jpg",
	}

	mockFile := new(mocksRW.FileData)

	mockFileOpener := new(mocksRW.FileOpener)
	mockFileOpener.On("Open").Return(mockFile, errors.New("failed to open file"))

	mockReadFile := new(mocksRW.Reader)
	mockUploader := new(mocksService.FileUploader)

	service := mockS3FileService(
		mockRepo,
		mockCfg,
		mockUploader,
		mockReadFile.ReadAll,
		mockImgDecoder.Decode,
		mockImgCompressor.CompressImage,
	)

	completion := make(chan model.Completion)

	go service.UploadFile(context.Background(), mockFileOpener, fileHeader.Filename, completion)

	result := <-completion

	assert.Error(t, result.Error)
	assert.Equal(t, "", result.FileURL)

	mockRepo.AssertExpectations(t)
	mockImgDecoder.AssertExpectations(t)
	mockImgCompressor.AssertExpectations(t)
	mockFile.AssertExpectations(t)
	mockFileOpener.AssertExpectations(t)
	mockReadFile.AssertExpectations(t)
	mockUploader.AssertExpectations(t)
}

func TestFileService_UploadFile_FileReadError(t *testing.T) {
	mockRepo := new(mocksRepo.FileRepository)
	mockCfg := mockCFG()
	mockImgDecoder := new(mocksImage.ImageDecoder)
	mockImgCompressor := new(mocksImage.ImageCompressor)

	fileHeader := &multipart.FileHeader{
		Filename: "testfile.jpg",
	}

	mockFile := new(mocksRW.FileData)
	mockFile.On("Close").Return(nil)

	mockFileOpener := new(mocksRW.FileOpener)
	mockFileOpener.On("Open").Return(mockFile, nil)

	fileBytes := []byte("fake file content")
	mockReadFile := new(mocksRW.Reader)
	mockReadFile.On("ReadAll", mock.Anything).Return(fileBytes, errors.New("failed to read file"))

	mockUploader := new(mocksService.FileUploader)

	service := mockS3FileService(
		mockRepo,
		mockCfg,
		mockUploader,
		mockReadFile.ReadAll,
		mockImgDecoder.Decode,
		mockImgCompressor.CompressImage,
	)

	completion := make(chan model.Completion)

	go service.UploadFile(context.Background(), mockFileOpener, fileHeader.Filename, completion)

	result := <-completion

	assert.Error(t, result.Error)
	assert.Equal(t, "", result.FileURL)

	mockRepo.AssertExpectations(t)
	mockImgDecoder.AssertExpectations(t)
	mockImgCompressor.AssertExpectations(t)
	mockFile.AssertExpectations(t)
	mockFileOpener.AssertExpectations(t)
	mockReadFile.AssertExpectations(t)
	mockUploader.AssertExpectations(t)
}

func TestFileService_UploadFile_UploadError(t *testing.T) {
	mockRepo := new(mocksRepo.FileRepository)
	mockCfg := mockCFG()
	mockImgDecoder := new(mocksImage.ImageDecoder)
	mockImgCompressor := new(mocksImage.ImageCompressor)

	fileHeader := &multipart.FileHeader{
		Filename: "testfile.jpg",
	}

	mockFile := new(mocksRW.FileData)
	mockFile.On("Close").Return(nil)

	mockFileOpener := new(mocksRW.FileOpener)
	mockFileOpener.On("Open").Return(mockFile, nil)

	fileBytes := []byte("fake file content")
	mockReadFile := new(mocksRW.Reader)
	mockReadFile.On("ReadAll", mock.Anything).Return(fileBytes, nil)

	mockUploader := new(mocksService.FileUploader)
	mockUploader.On("Upload", mock.Anything, mock.Anything).Return(nil, errors.New("failed upload to bucket"))

	service := mockS3FileService(
		mockRepo,
		mockCfg,
		mockUploader,
		mockReadFile.ReadAll,
		mockImgDecoder.Decode,
		mockImgCompressor.CompressImage,
	)

	completion := make(chan model.Completion)

	go service.UploadFile(context.Background(), mockFileOpener, fileHeader.Filename, completion)

	result := <-completion

	assert.Error(t, result.Error)
	assert.Equal(t, "", result.FileURL)

	mockRepo.AssertExpectations(t)
	mockImgDecoder.AssertExpectations(t)
	mockImgCompressor.AssertExpectations(t)
	mockFile.AssertExpectations(t)
	mockFileOpener.AssertExpectations(t)
	mockReadFile.AssertExpectations(t)
	mockUploader.AssertExpectations(t)
}

func TestFileService_UploadFile_ContextCanceled(t *testing.T) {
	mockRepo := new(mocksRepo.FileRepository)
	mockCfg := mockCFG()
	mockImgDecoder := new(mocksImage.ImageDecoder)
	mockImgCompressor := new(mocksImage.ImageCompressor)

	fileHeader := &multipart.FileHeader{
		Filename: "testfile.jpg",
	}

	mockFile := new(mocksRW.FileData)
	mockFileOpener := new(mocksRW.FileOpener)
	mockReadFile := new(mocksRW.Reader)
	mockUploader := new(mocksService.FileUploader)

	service := mockS3FileService(
		mockRepo,
		mockCfg,
		mockUploader,
		mockReadFile.ReadAll,
		mockImgDecoder.Decode,
		mockImgCompressor.CompressImage,
	)

	completion := make(chan model.Completion)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	go service.UploadFile(ctx, mockFileOpener, fileHeader.Filename, completion)

	result := <-completion

	assert.Error(t, result.Error)
	assert.Equal(t, "", result.FileURL)

	mockRepo.AssertExpectations(t)
	mockImgDecoder.AssertExpectations(t)
	mockImgCompressor.AssertExpectations(t)
	mockFile.AssertExpectations(t)
	mockFileOpener.AssertExpectations(t)
	mockReadFile.AssertExpectations(t)
	mockUploader.AssertExpectations(t)
}
