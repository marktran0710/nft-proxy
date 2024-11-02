package constants

import "github.com/gagliardetto/solana-go"

const (
	BASE64_PREFIX = ";base64,"
)

var (
	METAPLEX_CORE = solana.MustPublicKeyFromBase58("CoREENxT6tW1HoK8ypY1SxRMZTcVPm7R94rH4PZNhX7d")
	TOKEN_2022    = solana.MustPublicKeyFromBase58("TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb")
)

const MAX_TIMEOUT = 5

const (
	IMG_PNG  = "png"
	IMG_JPEG = "jpeg"
	IMG_JPG  = "jpg"
	IMG_GIF  = "gif"
	IMG_SVC  = "svc"
)
