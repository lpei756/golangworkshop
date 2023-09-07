package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func colors() []string {
	var result []string

	// Step 1: Return an array of 25x the same color
	// for i := 0; i < 25; i++ {
	// 	result = append(result, "blue")
	// }

	// Step 2: Alternate between 2 colors
	// for i := 0; i < 25; i++ {
	// 	if i%2 == 0 {
	// 		result = append(result, "yellow")
	// 	} else {
	// 		result = append(result, "pink")
	// 	}
	// }

	// Step 3: Make it a gradient
	for i := 0; i < 25; i++ {
		color := fmt.Sprintf("rgb(%d, %d, %d)", i*10, i*10, i*10)
		result = append(result, color)
	}

	// Step 4: Make it a rainbow
	// rainbowColors := []string{"red", "orange", "yellow", "green", "blue", "indigo", "violet"}
	// for i := 0; i < 25; i++ {
	// 	result = append(result, rainbowColors[i%len(rainbowColors)])
	// }

	return result
}

// No need to edit below this line

type colorsResponse struct {
	Colors []string `json:"colors"`
}

func Level1Handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(colorsResponse{
		Colors: colors(),
	})
}
