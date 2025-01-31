package service

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"github.com/ngikut-project-sprint/TutupLapak-File/internal/model"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/utils/rwutil"
)

func (s *S3FileService) UploadFile(ctx context.Context, file rwutil.FileOpener, fileName string, completion chan model.Completion) {
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

	contentType := http.DetectContentType(fileBytes)
	path := rwutil.GetFileFormat(fileName, contentType)
	key := fmt.Sprintf("%s/%s/images/%s", s.Team, s.Project, path)
	params := &s3.PutObjectInput{
		Bucket:      aws.String(s.Bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(contentType),
		ACL:         types.ObjectCannedACLPublicRead,
	}

	_, err := s.Uploader.Upload(ctx, params)
	if err != nil {
		completion <- model.Completion{FileURL: "", Error: ErrUploadImage(err)}
		return
	}

	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.Bucket, s.Region, key)
	completion <- model.Completion{FileURL: fileURL, Error: nil}
}
