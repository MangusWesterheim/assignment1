package assignment1

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var startTime = time.Now()

func init() {
	startTime = time.Now()
}

func DiagHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getDiag(w)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

}

func getDiag(w http.ResponseWriter) {

	//Puts api urls in variables
	uniURL := "http://universities.hipolabs.com/"
	ctryURL := "https://restcountries.com/"
	upTime := time.Since(startTime)

	//Creates a new client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	//Gets the response from the api
	uniResp, err := client.Get(uniURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(w, "An unexpected error occurred while processing the request.")
	}
	defer uniResp.Body.Close()

	//Gets the response from the api
	ctryResp, err := client.Get(ctryURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(w, "An unexpected error occurred while processing the request.")
	}
	defer ctryResp.Body.Close()

	//creates an instance of the diag struct
	diag := Diag{
		Universitiesapi: uniResp.Status,
		Countriespai:    ctryResp.Status,
		Version:         "v1.0",
		Uptime:          upTime.String(),
	}

	//Sets the content type to json
	w.Header().Set("content-type", "application/json")

	//Encodes the struct into json
	jsonData, err := json.Marshal(diag)
	/*
		//Checks for errors
		err = encoder.Encode(diag)
		if err != nil {
			http.Error(w, "Error during encoding", http.StatusInternalServerError)
			return
	*/

	//Writes the json data to the client
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}
