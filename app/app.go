package app

import (
	"context"
	"log"
	"s3_multiclient/config"
	"s3_multiclient/file/minio"
	"s3_multiclient/load"
	"s3_multiclient/server"
)

func Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatalf("%s", err.Error())
		return err
	}

	minioLoader, err := minio.Init(cfg.S3Cfg)
	if err != nil {
		return err
	}

	if err = minioLoader.CreateBucket(ctx, cfg.S3Cfg.Region); err != nil {
		return err
	}

	loader := load.Init(minioLoader)

	server := server.Init(ctx, loader)

	if err := server.MustStart(cfg.SrvCfg); err != nil { // тут внутри горутина
		return err
	}

	return nil
}
