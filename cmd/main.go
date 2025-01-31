package main

import (
	"context"
	"log"

	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/98prabowo/tutuplapak-file/cmd/api/rest"
	"github.com/98prabowo/tutuplapak-file/config"
	"github.com/98prabowo/tutuplapak-file/internal/model"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to open configuration")
	}

	awsCfg, err := awsConfig.LoadDefaultConfig(
		context.TODO(),
		awsConfig.WithRegion(cfg.AWS.Region),
		awsConfig.WithCredentialsProvider(cfg.AWS.GetCredential()),
	)
	if err != nil {
		log.Fatal("Failed to load AWS configuration")
	}

	db, err := initDB(&cfg.Database)
	if err != nil {
		log.Fatal("Failed to open connection to database")
	}

	// var wg sync.WaitGroup
	//
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	rest.StartRESTServer(db, cfg, awsCfg)
	// }()
	//
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	rpc.StartGRPCServer()
	// }()
	//
	// wg.Wait()
}

func initDB(cfg *config.DataBaseConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.File{}); err != nil {
		return nil, err
	}

	return db, nil
}
