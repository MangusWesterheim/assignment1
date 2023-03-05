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

	// Ensure interpretation as HTML by client ()
	w.Header().Set("content-type", "application/json")

	//Creates a new client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	//test for array
	parts := strings.Split(r.URL.Path, "/")
	//fmt.Println(parts[1]) //prints the parts of the url

	if len(parts) != 5 || parts[4] == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//Gets the response from the api
	uniResp, err := client.Get("http://universities.hipolabs.com/search?name=" + parts[4])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uniResp.Body.Close()

	var universities []University
	err = json.NewDecoder(uniResp.Body).Decode(&universities)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(w, "An unexpected error occurred while processing the request.")
		return
	}
	/*
		// Writes the uni response to the client (skriver ut midlertidig slett etterpå)
		err = json.NewEncoder(w).Encode(universities)
		if err != nil {
			http.Error(w, "Error when returning output", http.StatusInternalServerError)
			return
		}

	*/

	//starts handling the country api

	//puts all the isocodes in a slice
	var isocodes []string
	for _, uni := range universities {
		isocodes = append(isocodes, uni.Isocode)

	}
	//printer ut alle isokoder i slice
	fmt.Println("Her burde det være isokoder .... ->")
	fmt.Println(isocodes)

	//Gets the response from the api
	fmt.Println("Før client get...")

	something := "https://restcountries.com/v3.1/alpha?codes=" + strings.Join(isocodes, ",")
	ctryResp, err := client.Get(something)
	fmt.Println("Etter ...")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(w, "An unexpected error occurred while processing the request.")
		return
	}
	defer ctryResp.Body.Close()

	urlTest := "https://restcountries.com/v3.1/alpha?codes=" + strings.Join(isocodes, ",")
	fmt.Println(urlTest)
	fmt.Println(ctryResp.StatusCode)

	//skriver ut countries med matchende isocode funker ikke
	var countries []Countries
	err = json.NewDecoder(ctryResp.Body).Decode(&countries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(w, "An unexpected error occurred while processing the request.")
		return
	}
	/*
		// Writes the uni response to the client (skriver ut midlertidig slett etterpå)
		err = json.NewEncoder(w).Encode(countries)
		if err != nil {
			http.Error(w, "Error when returning output", http.StatusInternalServerError)
			return
		}

	*/
	var response []Response

	for _, uni := range universities {
		for _, ctry := range countries {
			if ctry.Isocode == uni.Isocode {
				response = append(response, Response{
					Name:      uni.Name,
					Country:   uni.Country,
					Isocode:   uni.Isocode,
					Webpages:  uni.WebPages,
					Languages: ctry.Languages,
					Maps:      ctry.Maps,
				})

			}

		}

	}
	JasonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(JasonData)
}
