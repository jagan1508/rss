package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter,code int , msg string){
	if code >499 {
		log.Println("Responding wwith 5xx error: ",msg)
	}
	type errResponse struct{
		Error string `json:"error"`
	}
	respondJson(w,code,errResponse{
		Error: msg,
	})
}

func respondJson(w http.ResponseWriter, code int, payload interface{}){
	data, err := json.Marshal(payload)
	if err !=nil {
		log.Printf("Failed to marshall JSON data %v", payload)
		w.WriteHeader(500)
		return 
	}
 	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(code)
	w.Write(data)
}
