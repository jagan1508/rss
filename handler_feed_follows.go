package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	_ "github.com/jagan1508/rss/internal/auth"
	"github.com/jagan1508/rss/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID string `json:"feedid"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing json: %v", err))
		return
	}
	uuid_st := uuid.NewString()
	err = apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid_st,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		FeedID:    params.FeedID,
		UserID:    user.ID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Cannot create feed follows: %v", err))
		return
	}

	feed_follows, err := apiCfg.DB.GetFeedFollows(r.Context(), uuid_st)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Cannot create feed follows %v", err))
		return
	}
	respondJson(w, 201, databaseFeedFollowsToFeedFollows(feed_follows))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feeds_follows, err := apiCfg.DB.GetFeedsFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Cannot get feed follows %v", err))
		return
	}
	respondJson(w, 201, databaseFeedsFollowsToFeedsFollows(feeds_follows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowId := chi.URLParam(r, "feedFollowId")
	err := apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Cannot delete feed follows %v", err))
		return
	}
	respondJson(w, 200, struct{}{})
}
