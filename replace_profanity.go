package main

import (
	"strings"
)

func replaceProfanity(text string) string {
	words := strings.Split(text, " ")
	for i := range words {
		if strings.ToLower(words[i]) == "kerfuffle" ||
			strings.ToLower(words[i]) == "sharbert" ||
			strings.ToLower(words[i]) == "fornax" {
			words[i] = "****"
		}
	}

	return strings.Join(words[:], " ")
}
