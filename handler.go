package main

import "net/http" 

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondJson(w,200,struct{}{})
}