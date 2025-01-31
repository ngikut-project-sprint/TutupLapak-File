package rest

import (
	"image/jpeg"
	"io"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/98prabowo/tutuplapak-file/config"
	"github.com/98prabowo/tutuplapak-file/internal/controller"
	"github.com/98prabowo/tutuplapak-file/internal/repository"
	"github.com/98prabowo/tutuplapak-file/internal/service"
	"github.com/98prabowo/tutuplapak-file/internal/utils/imageutil"
	"github.com/98prabowo/tutuplapak-file/internal/utils/validation"
)

func InitRouter(
	e *echo.Echo,
	db *gorm.DB,
	cfg *config.Config,
	awsCfg aws.Config,
) {
	fileRepository := repository.NewFileRepository(*db)
	fileService := service.NewS3FileService(
		fileRepository,
		cfg,
		awsCfg,
		io.ReadAll,
		jpeg.Decode,
		imageutil.CompressImage,
	)
	fileController := controller.NewFileController(
		fileService,
		cfg.File.FileMaxSize,
		validation.ValidateFile,
	)

	versionGroup := e.Group("/v1")
	fileGroup := versionGroup.Group("/file")

	fileGroup.POST("", fileController.Post)

	fileGroup.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "healthty",
			"time":   time.Now().String(),
		})
	})
}
