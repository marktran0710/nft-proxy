package services

import (
	"log"
	"testing"

	"github.com/gagliardetto/solana-go"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env") // Not implement ENV file
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestSolanaImageService_FetchMetadata(t *testing.T) {
	pk := solana.MustPublicKeyFromBase58("CJ9AXYbSUPoR95oMvWzgCV3GbG3ZubQjFUpRHN7xqAVb")

	svc := SolanaService{}
	svc.Start()

	d, _, err := svc.TokenData(pk)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v\n", d)
}
