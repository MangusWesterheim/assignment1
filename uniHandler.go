package assignment1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func UniHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUni(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func getUni(w http.ResponseWriter, r *http.Request) {
	//Puts api urls in variables
	// uniURL := "http://universities.hipolabs.com/"

	// Ensure interpretation as HTML by client ()
	w.Header().Set("content-type", "application/json")

	//Creates a new client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	parts := strings.Split(r.URL.Path, "/")
	fmt.Println(parts[1]) //prints the parts of the url

	if len(parts) != 5 || parts[4] == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//Gets the response from the api
	uniResp, err := client.Get("http://universities.hipolabs.com/search?name=" + parts[4])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(w, "An unexpected error occurred while processing the request.")
	}
	defer uniResp.Body.Close()

	var uni []University
	err = json.NewDecoder(uniResp.Body).Decode(&uni)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(w, "An unexpected error occurred while processing the request.")
	}

	//Writes the uni response to the client (skriver ut midlertidig slett etterpå)
	err = json.NewEncoder(w).Encode(uni)
	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
	}

}
