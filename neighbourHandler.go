package assignment1

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func NeighbourHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetNeighbour(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func GetNeighbour(w http.ResponseWriter, r *http.Request) {

	// Ensure interpretation as HTML by client ()
	w.Header().Set("content-type", "application/json")

	client := &http.Client{}
	defer client.CloseIdleConnections()

	parts := strings.Split(r.URL.Path, "/")

	//Gets the response from the api
	ctryResp, err := client.Get("https://restcountries.com/v3.1/name/" + parts[4] + "?fullText=true")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer ctryResp.Body.Close()

	//creates a slice that stores the response from the api
	var countries []Countries
	//populates the slice with the response from the api
	err = json.NewDecoder(ctryResp.Body).Decode(&countries)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(w, "An unexpected error occurred while processing the request.")
		return
	}

	// Writes the uni response to the client (skriver ut midlertidig slett etterpå)
	//err = json.NewEncoder(w).Encode(countries)
	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
		return
	}
	//variable that stores the neighboring countries
	neighbourCodes := countries[0].Borders
	fmt.Println(neighbourCodes)

	uniresp, err := client.Get("http://universities.hipolabs.com/search?name=" + parts[5] + "&country=" + parts[4])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uniresp.Body.Close()
	//Gets the response from university api the api
	var universities []University
	err = json.NewDecoder(uniresp.Body).Decode(&universities)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(w, "An unexpected error occurred while processing the request.")
		return
	}
	// Writes the uni response to the client (skriver ut midlertidig slett etterpå)
	//err = json.NewEncoder(w).Encode(universities)
	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
		return
	}
	//variable that stores the neighbouring countries
	urlcodes := "https://restcountries.com/v3.1/alpha?codes="

	//Adds the neighbouring countries to the url
	for i := 0; i < len(neighbourCodes); i++ {
		if i == len(neighbourCodes)-1 {
			urlcodes += neighbourCodes[i]
		} else {
			urlcodes += neighbourCodes[i] + ","
		}
	}

	//Gets the response from the api for the neighbouring countries
	ctryResp, err = client.Get(urlcodes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer ctryResp.Body.Close()
	//Gets the response from the api for the neighbouring countries
	var borderCountry []Countries
	err = json.NewDecoder(ctryResp.Body).Decode(&borderCountry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	uniResp, err := client.Get("http://universities.hipolabs.com/search?name=" + parts[5])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer uniResp.Body.Close()

	//Creates a slice of the neighbouring universities
	var neighbourUniversities []University
	//Populates the slice with the neighbouring universities
	err = json.NewDecoder(uniResp.Body).Decode(&neighbourUniversities)
	//Checks for errors
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(w, "An unexpected error occurred while processing the request.")
		return
	}

	//creates a slice to store response
	var response []Response

	//loop to check if neighbouring universities are in neighbouring countries, matching them and adding them to the response
	for _, uni := range neighbourUniversities {
		for _, ctry := range borderCountry {
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
	//marshals the response to json
	JasonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error when returning output", http.StatusInternalServerError)
		return
	}
	//writes the response to the client
	w.WriteHeader(http.StatusOK)
	w.Write(JasonData)
}
