package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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
	feed, err := urlToFeed("https://www.wagslane.dev/index.xml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(feed)
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

	db := database.New(conn)
	apiCfg := apiConfig{
		DB: db,
	}

	go startScraping(db, 10, time.Minute)
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
	v1Router.Get("/users", apiCfg.middleware_auth(apiCfg.handlerGetUser))
	v1Router.Post("/feeds", apiCfg.middleware_auth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
	v1Router.Post("/feed_follows", apiCfg.middleware_auth(apiCfg.handlerCreateFeedFollows))
	v1Router.Get("/feed_follows", apiCfg.middleware_auth(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowId}", apiCfg.middleware_auth(apiCfg.handlerDeleteFeedFollows))
	v1Router.Get("/posts", apiCfg.middleware_auth(apiCfg.handlerGetPostsForUser))

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
