package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	_ "github.com/jagan1508/rss/internal/auth"
	"github.com/jagan1508/rss/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing json: %v", err))
		return
	}
	uuid_st := uuid.NewString()
	err = apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid_st,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Cannot create feed: %v", err))
		return
	}

	feed, err := apiCfg.DB.GetFeed(r.Context(), uuid_st)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Cannot create feed %v", err))
		return
	}
	respondJson(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Couldn't get feeds %v",err))
		return
	}
	respondJson(w, 200, databaseFeedsToFeeds(feeds))
}
