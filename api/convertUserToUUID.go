package api

import (
	"encoding/json"
	"net/http"

	"github.com/charmbracelet/log"
)

type UserResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ConvertUserToUUID(username string) string {
	// Get UUID from username
	url := "https://api.mojang.com/users/profiles/minecraft/" + username
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var user UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Fatal("Error decoding JSON:", "err", err, "resp", resp)
	}

	if user.ID == "" {
		log.Fatal("uuid cannot be empty")
	}

	return user.ID
}
