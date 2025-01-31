package service_test

import (
	"context"
	"image"
	"image/color"
	"image/draw"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/stretchr/testify/require"

	mockConfig "github.com/ngikut-project-sprint/TutupLapak-File/config"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/repository"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/service"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/utils/imageutil"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/utils/rwutil"
)

func mockAWSConfig(t *testing.T) aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-1"))
	require.NoError(t, err)
	return cfg
}

func mockCFG() *mockConfig.Config {
	return &mockConfig.Config{
		Project: mockConfig.ProjectConfig{
			Team: "team",
			Name: "name",
		},
		AWS: mockConfig.AWSConfig{
			BucketName: "bucket",
			Region:     "ap-southeast-1",
		},
		File: mockConfig.FileConfig{
			ThumbnailMaxSize: 10240,
		},
	}
}

func mockS3FileService(
	repo repository.FileRepository,
	cfg *mockConfig.Config,
	uploader service.FileUploader,
	readFile rwutil.ReadFile,
	decodeImage imageutil.DecodeImage,
	compressImage imageutil.ImageCompress,
) service.FileService {
	return &service.S3FileService{
		Repo:             repo,
		Uploader:         uploader,
		Team:             cfg.Project.Team,
		Project:          cfg.Project.Name,
		Bucket:           cfg.AWS.BucketName,
		Region:           cfg.AWS.Region,
		ThumbnailMaxSize: cfg.File.ThumbnailMaxSize,
		ReadFile:         readFile,
		DecodeImage:      decodeImage,
		CompressImage:    compressImage,
	}
}

func createDummyImage(width, height int) image.Image {
	// Create an RGBA image
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Fill it with a solid color (e.g., red)
	red := color.RGBA{255, 0, 0, 255}
	draw.Draw(img, img.Bounds(), &image.Uniform{red}, image.Point{}, draw.Src)

	return img
}
