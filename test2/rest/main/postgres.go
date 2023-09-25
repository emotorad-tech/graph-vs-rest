package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func PostgresConnect() *pgxpool.Pool {
	encodedPassword := url.QueryEscape("3Motorad")
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		"postgres",
		encodedPassword,
		"backend-dev-1.clsrq9deqixe.ap-south-1.rds.amazonaws.com",
		5432,
		"postgres",
		"require",
	)
	count := 0
	for {
		db, err := pgxpool.New(context.Background(), connectionString)
		if err != nil {
			count++
		} else {
			log.Println("connected to postgres!")
			return db
		}
		if count == 5 {
			log.Println("unable to connect to postgres: ", err)
			log.Println("retying in 5 seconds...")
			time.Sleep(time.Second * 5)
			count = 0
		}
	}
}
