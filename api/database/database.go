package database

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var Ctx = context.Background()

func CreateClient(dbno int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("DATABASE_ADRR"),
		Password: os.Getenv("DATABASE_PASS"),
		DB:       dbno,
	})
	if err := rdb.Ping(Ctx).Err(); err != nil {
		log.Fatal(err)
	}
	return rdb
}
