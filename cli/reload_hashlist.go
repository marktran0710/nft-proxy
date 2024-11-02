package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	nft_proxy "github.com/alphabatem/nft-proxy"
	"github.com/alphabatem/nft-proxy/constants"
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

type HashList []string

func main() {
	mainContext, err := context.NewCtx(
		&services.SqliteService{},
		&services.SolanaImageService{},
		&services.ImageService{},
		&services.ResizeService{},
		&services.SolanaService{},
	)
	if err != nil {
		log.Fatal("Reloading hashlist failed: %s", err)
	}

	err = mainContext.Run()
	if err != nil {
		log.Panic("Reload hashlist runs context failed: %s", err)
	}

	err = run(mainContext)
	if err != nil {
		log.Panic("Reload hashlist error: %s", err)
	}
}

func run(ctx *context.Context) error {
	db := ctx.Service(services.SQLITE_SVC).(*services.SqliteService)

	hashes, err := loadHashlist("./hashlist.json")
	if err != nil {
		return err
	}

	log.Printf("Mints: %v", len(hashes))

	var amount int64
	err = db.Db().Model(&nft_proxy.SolanaMedia{}).Count(&amount).Error
	if err != nil {
		return err
	}
	log.Printf("Before: %v", amount)

	err = db.Db().
		Where(`mint IN ("` + strings.Join(hashes, `","`) + `")`).
		Delete(&nft_proxy.SolanaMedia{}).Error
	if err != nil {
		return err
	}

	err = db.Db().Model(&nft_proxy.SolanaMedia{}).Count(&amount).Error
	if err != nil {
		return err
	}
	log.Printf("After: %v", amount)

	//Use to reload the remote DB
	//reloadRemote(hashes)

	//Use to reload the local DB
	//img := ctx.Service(services.SOLANA_IMG_SVC).(*services.SolanaImageService)
	//reloadLocally(img, hashes)

	return nil
}

func reloadRemote(hashes HashList) error {
	c := &http.Client{Timeout: constants.MAX_TIMEOUT * time.Second}
	for _, h := range hashes {
		log.Printf("Loading hash: %s", h)

		_, err := c.Get(fmt.Sprintf("https://api.degencdn.com/v1/nfts/%s/image.jpg", h))
		if err != nil {
			log.Printf("Failed media: %s - %s", h, err)
		}
	}
	return nil
}

func reloadLocally(img *services.SolanaImageService, hashes HashList) error {
	for _, h := range hashes {
		log.Printf("Loading hash: %s", h)
		_, err := img.Media(h, true)
		if err != nil {
			log.Printf("Failed media: %s - %s", h, err)
		}
	}
	return nil
}

func loadHashlist(location string) (HashList, error) {
	data, err := os.ReadFile(location)
	if err != nil {
		return nil, err
	}

	var hashList HashList
	err = json.Unmarshal(data, &hashList)
	if err != nil {
		return nil, err
	}

	return hashList, nil
}
