package service

import (
	"bytes"
	"context"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	"github.com/ngikut-project-sprint/TutupLapak-File/internal/model"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/utils/rwutil"
)

func (s *s3FileService) UploadFile(ctx context.Context, file *multipart.FileHeader, fileName string, completion chan model.Completion) {
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

	contentType := http.DetectContentType(fileBytes)
	path := rwutil.GetFileFormat(fileName, contentType)
	key := fmt.Sprintf("%s/%s/images/%s", s.team, s.project, path)
	params := &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(contentType),
		ACL:         types.ObjectCannedACLPublicRead,
	}

	_, err := s.uploader.Upload(ctx, params)
	if err != nil {
		completion <- model.Completion{FileURL: "", Error: ErrUploadImage(err)}
		return
	}

	fileURL := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, key)
	completion <- model.Completion{FileURL: fileURL, Error: nil}
}
