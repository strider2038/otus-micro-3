/*
 * User Service
 *
 * This is simple client API
 *
 * API version: 1.0.0
 * Contact: schetinnikov@gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"user-service/internal/di"

	"github.com/jackc/pgx/v4/pgxpool"
)

var (
	version = ""
)

const (
	defaultPort = "8000"
)

func main() {
	log.Printf("Server started")
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "unknown"
	}

	connection, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("failed to connect to postgres:", err)
	}
	router := di.NewRouter(connection, version, env)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
