package main

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
)

func handlerChirp(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		// these tags indicate how the keys in the JSON should be mapped to the struct fields
		// the struct fields must be exported (start with a capital letter) if you want them parsed
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		// an error will be thrown if the JSON is invalid or has the wrong types
		// any missing fields will simply have their values in the struct set to their zero value
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}
	// params is a struct with data populated successfully
	if len(params.Body) <= 140 {
		respondValid(w, params.Body)
	} else {
		respondInvalid(w)
	}
}

func respondValid(w http.ResponseWriter, body string) {

	type valid struct {
		// the key will be the name of struct field unless you give it an explicit JSON tag
		CleanedBody string `json:"cleaned_body"`
	}

	respBody := valid{
		CleanedBody: cleanJsonBody(body),
	}
	dat, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}

func respondInvalid(w http.ResponseWriter) {

	type valid struct {
		// the key will be the name of struct field unless you give it an explicit JSON tag
		Error string `json:"error"`
	}

	respBody := valid{
		Error: "Chirp is too long",
	}
	dat, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	w.Write(dat)
}

func cleanJsonBody(body string) string {
	replacements := [3]string{"kerfuffle", "sharbert", "fornax"}

	for _, replacement := range replacements {
		pattern := "(?i)" + regexp.QuoteMeta(replacement)
		re := regexp.MustCompile(pattern)
		body = re.ReplaceAllString(body, "****")

	}

	return body
}
