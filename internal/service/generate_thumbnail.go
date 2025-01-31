package service

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"github.com/ngikut-project-sprint/TutupLapak-File/internal/model"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/utils/rwutil"
)

func (s *S3FileService) GenerateThumbnail(ctx context.Context, file rwutil.FileOpener, fileName string, completion chan model.Completion) {
	defer close(completion)

	if err := ctx.Err(); err != nil {
		completion <- model.Completion{FileURL: "", Error: err}
		return
	}

	src, fileErr := file.Open()
	if fileErr != nil {
		completion <- model.Completion{FileURL: "", Error: fileErr}
		return
	}
	defer src.Close()

	fileBytes, buffErr := s.ReadFile(src)
	if buffErr != nil {
		completion <- model.Completion{FileURL: "", Error: ErrReadBuffer(buffErr)}
		return
	}

	img, imageErr := s.DecodeImage(bytes.NewReader(fileBytes))
	if imageErr != nil {
		completion <- model.Completion{FileURL: "", Error: ErrDecodeThumbnail(imageErr)}
		return
	}

	thumbBuff := &bytes.Buffer{}
	if err := s.CompressImage(thumbBuff, img, s.ThumbnailMaxSize); err != nil {
		completion <- model.Completion{FileURL: "", Error: ErrEncodeThumbnail(err)}
		return
	}

	thumbKey := fmt.Sprintf("%s/%s/thumbnails/%s.jpeg", s.Team, s.Project, fileName)
	thumbParams := &s3.PutObjectInput{
		Bucket:      aws.String(s.Bucket),
		Key:         aws.String(thumbKey),
		Body:        bytes.NewReader(thumbBuff.Bytes()),
		ContentType: aws.String("image/jpeg"),
		ACL:         types.ObjectCannedACLPublicRead,
	}

	_, err := s.Uploader.Upload(ctx, thumbParams)
	if err != nil {
		completion <- model.Completion{FileURL: "", Error: ErrUploadThumbnail(err)}
		return
	}

	thumbURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.Bucket, s.Region, thumbKey)
	completion <- model.Completion{FileURL: thumbURL, Error: nil}
}
