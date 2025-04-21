package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-sql-driver/mysql"
	"github.com/jagan1508/rss/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	godotenv.Load()

	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT not found in the environment")
	}

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DB_USER")
	cfg.Passwd = os.Getenv("DB_PASS")
	cfg.Net = "tcp"
	cfg.Addr = "172.23.112.1:3306"
	cfg.DBName = "rssagg"
	cfg.ParseTime = true

	conn, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	apiCfg := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowCredentials: false,
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerErr)
	v1Router.Post("/users", apiCfg.handlerUser)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portStr,
	}
	fmt.Println("Listening on Port: ", portStr)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
