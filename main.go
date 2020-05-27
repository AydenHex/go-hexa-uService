package main

import (
	"fmt"
	"os"

	"github.com/aydenhex/go-hexa-uservice/shortener"
	rr "github.com/aydenhex/go-hexa-uservice/shortener/repository/redis"
)

func main() {

}

func httpPort() string {
	port := "8000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func chooseRepo() shortener.RedirectRepository {
	switch os.Getenv("URL_DB") {
	case "redis":
		redisURL := os.Getenv("REDIS_URL")
		repo, err := rr.New
	}
}
