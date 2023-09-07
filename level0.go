package main

import (
	"encoding/json"
	"net/http"
	// "os/user"
)

// Hello gopher!
//
// This is the level 0 of the Go workshop. You don't have to do anything here
// but if you wish you can update your name or avatar below!
//
// Don't forget to restart your program after each change!

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	// usr, _ := user.Current()

	info := Info{
		// Update your name and avatar here!
		Name:      "Lei",
		AvatarURL: "https://i.keaimeitu.com/up/e1/f4/a4/92d29b6e51af79ad175d0f0595a4f4e1.jpg",
	}

	json.NewEncoder(w).Encode(info)
}

type Info struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}
