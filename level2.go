package main

import (
	"encoding/json"
	"net/http"
)

// Level2 is about using structs and storing state.
// Use the struct below to store a state and make colors move!
//
// Step 1: make a square move at every invocation (learn about structs)
// Add a field to the `Level2` struct and update it in the `colors` function.
//
// Step 2: use defer to change the state (learn about defer).
// defer is an effective way of executing logic at the end of functions, try
// using it to update your state!
//
// Step 3: get the color as a constructor parameter (learn how to add parameters to a struct)
// Update the `NewLevel2` function to accept a color parameter.
type Level2 struct {
	Position int
	Color    string
}

// Step 3: Update the NewLevel2 function to accept a color parameter
func NewLevel2(color string) *Level2 {
	return &Level2{
		Color: color,
	}
}

func (l *Level2) colors() []string {
	colorsArray := make([]string, 25)
	for i := range colorsArray {
		if i == l.Position {
			colorsArray[i] = l.Color
		} else {
			colorsArray[i] = "pink"
		}
	}

	// Step 2: Use defer to update the position at the end of the function
	defer func() {
		l.Position = (l.Position + 1) % 25
	}()

	return colorsArray
}

// No need to edit below this line

func (l *Level2) Handler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(colorsResponse{
		Colors: l.colors(),
	})
}
