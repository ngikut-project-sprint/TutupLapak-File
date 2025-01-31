package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/ngikut-project-sprint/TutupLapak-File/internal/model"
	"github.com/ngikut-project-sprint/TutupLapak-File/internal/repository"
)

func TestFileRepository_AddFile_Success(t *testing.T) {
	db, mockDB, mock := setupTestDB(t)
	defer func() {
		assert.NoError(t, mock.ExpectationsWereMet())
		mockDB.Close()
	}()

	repo := repository.NewFileRepository(*db)

	ctx := context.Background()
	file := model.File{
		FileURI:          "https://example.com/file.jpg",
		FileThumbnailURI: "https://example.com/thumbnail.jpg",
	}

	mock.ExpectBegin()
	mock.MatchExpectationsInOrder(false)
	mock.ExpectQuery(`INSERT INTO "files" (.+) VALUES (.+) RETURNING "file_id"`).
		WithArgs(
			file.FileURI,
			file.FileThumbnailURI,
			sqlmock.AnyArg(), // CreatedAt
			sqlmock.AnyArg(), // UpdatedAt
			sqlmock.AnyArg(), // DeletedAt
		).
		WillReturnRows(sqlmock.NewRows([]string{"file_id"}).AddRow(1))
	mock.ExpectCommit()

	fileID, err := repo.AddFile(ctx, file)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), fileID)
}

func TestFileRepository_AddFile_TimeoutError(t *testing.T) {
	db, mockDB, mock := setupTestDB(t)
	defer func() {
		assert.NoError(t, mock.ExpectationsWereMet())
		mockDB.Close()
	}()

	repo := repository.NewFileRepository(*db)

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	file := model.File{
		FileURI:          "https://example.com/file.jpg",
		FileThumbnailURI: "https://example.com/thumbnail.jpg",
	}

	mock.ExpectBegin()
	mock.ExpectQuery("INSERT INTO \"files\"").WillReturnError(context.DeadlineExceeded)
	mock.ExpectRollback()

	_, err := repo.AddFile(ctx, file)

	if err == nil {
		t.Fatalf("Expected timeout error, but got no error")
	}

	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("Expected context.DeadlineExceeded error, but got: %v", err)
	}
}

func setupTestDB(t *testing.T) (*gorm.DB, *sql.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)

	dialector := postgres.New(postgres.Config{
		Conn:                 mockDB,
		PreferSimpleProtocol: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(t, err)

	return db, mockDB, mock
}
