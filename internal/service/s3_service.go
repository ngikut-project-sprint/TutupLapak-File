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

type s3FileService struct {
	repo             repository.FileRepository
	uploader         *manager.Uploader
	team             string
	project          string
	bucket           string
	region           string
	thumbnailMaxSize int64
	readFile         rwutil.ReadFile
	decodeImage      imageutil.DecodeImage
	compressImage    imageutil.ImageCompress
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
	return &s3FileService{
		repo:             repo,
		uploader:         uploader,
		team:             cfg.Project.Team,
		project:          cfg.Project.Name,
		bucket:           cfg.AWS.BucketName,
		region:           cfg.AWS.Region,
		thumbnailMaxSize: cfg.File.ThumbnailMaxSize,
		readFile:         readFile,
		decodeImage:      decodeImage,
		compressImage:    compressImage,
	}
}
