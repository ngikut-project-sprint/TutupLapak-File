package service

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/ngikut-project-sprint/TutupLapak-File/config"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/repository"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/utils/imageutil"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/utils/rwutil"
)

type S3FileService struct {
	Repo             repository.FileRepository
	Uploader         FileUploader
	Team             string
	Project          string
	Bucket           string
	Region           string
	ThumbnailMaxSize int64
	ReadFile         rwutil.ReadFile
	DecodeImage      imageutil.DecodeImage
	CompressImage    imageutil.ImageCompress
}

func NewS3FileService(
	repo repository.FileRepository,
	cfg *config.Config,
	awsCfg aws.Config,
	readFile rwutil.ReadFile,
	decodeImage imageutil.DecodeImage,
	compressImage imageutil.ImageCompress,
) FileService {
	client := s3.NewFromConfig(awsCfg)
	uploader := manager.NewUploader(client)
	return &S3FileService{
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
