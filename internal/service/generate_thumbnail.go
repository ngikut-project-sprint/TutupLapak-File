package service

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"github.com/ngikut-project-sprint/TutupLapak-File/internal/model"
)

func (s *s3FileService) GenerateThumbnail(ctx context.Context, file *multipart.FileHeader, fileName string, completion chan model.Completion) {
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

	fileBytes, buffErr := s.readFile(src)
	if buffErr != nil {
		completion <- model.Completion{FileURL: "", Error: ErrReadBuffer(buffErr)}
		return
	}

	img, imageErr := s.decodeImage(bytes.NewReader(fileBytes))
	if imageErr != nil {
		completion <- model.Completion{FileURL: "", Error: ErrDecodeThumbnail(imageErr)}
		return
	}

	thumbBuff := &bytes.Buffer{}
	if err := s.compressImage(thumbBuff, img, s.thumbnailMaxSize); err != nil {
		completion <- model.Completion{FileURL: "", Error: ErrEncodeThumbnail(err)}
		return
	}

	thumbKey := fmt.Sprintf("%s/%s/thumbnails/%s.jpeg", s.team, s.project, fileName)
	thumbParams := &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(thumbKey),
		Body:        bytes.NewReader(thumbBuff.Bytes()),
		ContentType: aws.String("image/jpeg"),
		ACL:         types.ObjectCannedACLPublicRead,
	}

	_, err := s.uploader.Upload(ctx, thumbParams)
	if err != nil {
		completion <- model.Completion{FileURL: "", Error: ErrUploadThumbnail(err)}
		return
	}

	thumbURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, thumbKey)
	completion <- model.Completion{FileURL: thumbURL, Error: nil}
}
