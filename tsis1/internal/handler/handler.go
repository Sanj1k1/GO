package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

func GetPlayerByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	name := params["name"]
	for _, player := range players {
		if player.FullName == name {
			json.NewEncoder(w).Encode(player)
			return
		}
	}
	http.Error(w, "Player not found", http.StatusNotFound)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello"))
}
