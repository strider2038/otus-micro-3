package di

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"user-service/internal/api"
	"user-service/internal/postgres"
	"user-service/internal/postgres/database"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewRouter(connection *pgxpool.Pool, version, env string) http.Handler {
	db := database.New(connection)

	userRepository := postgres.NewUserRepository(db)
	userApiService := api.NewUserApiService(userRepository)
	userApiController := api.NewUserApiController(userApiService)

	apiRouter := api.NewRouter(userApiController)
	metrics := api.NewMetrics("user_service")
	apiRouter.Use(func(handler http.Handler) http.Handler {
		return api.MetricsMiddleware(handler, metrics)
	})

	router := mux.NewRouter()
	router.PathPrefix("/api").Handler(apiRouter)
	router.Handle("/metrics", promhttp.Handler())

	router.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("content-type", "application/json")
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte(`{"status":"ok"}`))
	})

	router.HandleFunc("/ready", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("content-type", "application/json")
		err := connection.Ping(context.Background())
		if err == nil {
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte(`{"status":"ok"}`))
		} else {
			writer.WriteHeader(http.StatusServiceUnavailable)
			writer.Write([]byte(`{"status":"not available"}`))
		}
	})

	router.HandleFunc("/version", func(writer http.ResponseWriter, request *http.Request) {
		dbVersion, err := postgres.GetMigrationsVersion(connection)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Println("error:", err)
			return
		}

		writer.Header().Set("content-type", "application/json")
		writer.WriteHeader(http.StatusOK)
		response, _ := json.Marshal(struct {
			Environment        string `json:"environment"`
			ApplicationVersion string `json:"application_version"`
			DatabaseVersion    int64  `json:"database_version"`
		}{
			Environment:        env,
			ApplicationVersion: version,
			DatabaseVersion:    dbVersion,
		})
		writer.Write(response)
	})

	return router
}
