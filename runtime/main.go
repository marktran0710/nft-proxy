package main

import (
	"log"

	services "github.com/alphabatem/nft-proxy/service"
	"github.com/babilu-online/common/context"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	ctx, err := context.NewCtx(
		&services.SqliteService{},
		&services.StatService{},
		&services.ResizeService{},
		&services.SolanaService{},
		&services.SolanaImageService{},
		&services.ImageService{},
		&services.HttpService{},
	)

	if err != nil {
		log.Fatalf("Loading context error: %s", err)
		return
	}

	err = ctx.Run()
	if err != nil {
		log.Fatalf("Running failed: %s", err)
	}

	log.Printf("Running server!")
}
