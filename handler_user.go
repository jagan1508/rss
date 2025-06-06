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

func (apiCfg *apiConfig) handlerUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error Parsing json: %v", err))
		return
	}
	uuid_st := uuid.NewString()
	err = apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid_st,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Cannot create user: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUser(r.Context(), uuid_st)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Cannot create user %v", err))
		return
	}
	respondJson(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	respondJson(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetPostsForUser(w http.ResponseWriter, r *http.Request, user database.User) {
	posts,err :=apiCfg.DB.GetPostsForUser(r.Context(),database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err !=nil{
		respondWithError(w,400,fmt.Sprintf("Couldn't get posts : %v",err))
		return
	}
	respondJson(w,200,databasePostsToPosts(posts))
}

