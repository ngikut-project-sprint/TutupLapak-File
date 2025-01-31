package controller

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/98prabowo/tutuplapak-file/internal/model"
	"github.com/98prabowo/tutuplapak-file/internal/utils/errorutil"
)

func (fc *FileController) Post(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return errorutil.SendErrorResponse(c, "File is required", http.StatusBadRequest)
	}

	if err := fc.validateFile(file, fc.maxFileSize); err != nil {
		return errorutil.SendErrorResponse(c, err.Error(), http.StatusBadRequest)
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), 2*time.Second)
	defer cancel()

	id := uuid.New().String()
	imageChan := make(chan model.Completion)
	thumbChan := make(chan model.Completion)

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer wg.Done()
		fc.service.UploadFile(ctx, file, id, imageChan)
	}()

	go func() {
		defer wg.Done()
		fc.service.GenerateThumbnail(ctx, file, id, thumbChan)
	}()

	imageComp := <-imageChan
	if imageComp.Error != nil {
		log.Printf("Failed to upload file: %s", imageComp.Error.Error())
		return errorutil.SendErrorResponse(c, "Server error", http.StatusInternalServerError)
	}

	thumbComp := <-thumbChan
	if thumbComp.Error != nil {
		log.Printf("Failed to upload file: %s", thumbComp.Error.Error())
		return errorutil.SendErrorResponse(c, "Server error", http.StatusInternalServerError)
	}

	wg.Wait()

	fileID, sqlErr := fc.service.AddFile(ctx, imageComp.FileURL, thumbComp.FileURL)
	if sqlErr != nil {
		log.Printf("Failed to add into sql: %s", sqlErr.Error())
		return errorutil.SendErrorResponse(c, "Server error", http.StatusInternalServerError)
	}

	fileRes := &model.AddFileResponse{
		FileID:           strconv.FormatUint(uint64(fileID), 10),
		FileURI:          imageComp.FileURL,
		FileThumbnailURI: thumbComp.FileURL,
	}

	return c.JSON(http.StatusOK, fileRes)
}
