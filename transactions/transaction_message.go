package transactions

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"
)

const (
	//URL of the Authorization microservice to do requests
	authorizationURL = "http://localhost:8888/authorization"
)

// Do http request to the authorization microservice and returns the authorization flag
// to perform the creation of the new transaction table on the Handler that creates the
// new transaction on the database
func authorizeCreationRequest(transaction Transaction) (bool, error) {
	var authorized bool
	transactioJSON, err := json.Marshal(transaction)
	if err != nil {
		return authorized, err
	}

	client := http.Client{}
	urlAuthoADD := authorizationURL + "/add/"
	requestAuthorization, err := http.NewRequest("POST", urlAuthoADD, bytes.NewBuffer(transactioJSON))
	if err != nil {
		return authorized, nil
	}
	requestAuthorization.Header.Set("Content-Type", "application/json")
	respAuthorization, err := client.Do(requestAuthorization)
	if err != nil {
		return authorized, err
	}
	defer respAuthorization.Body.Close()

	if respAuthorization.StatusCode != http.StatusOK {
		return authorized, nil
	}

	type AuthorizationResponse struct {
		Authorized bool `json:"authorized"`
	}

	var responseAuthorized AuthorizationResponse
	err = json.NewDecoder(respAuthorization.Body).Decode(&responseAuthorized)
	authorized = responseAuthorized.Authorized

	return authorized, err
}

// Do the http request to the authorization microservice and returns the authorization flag
// to delete a transaction table in the database
func authorizeDeleteRequest(id int) (bool, error) {
	var authorized bool

	client := http.Client{}
	idStr := strconv.Itoa(id)
	urlAuthoDelete := authorizationURL + "/delete/" + idStr

	requestAuthorization, err := http.NewRequest("GET", urlAuthoDelete, nil)

	if err != nil {
		return authorized, err
	}

	requestAuthorization.Header.Set("Content-Type", "application/json")
	respAuthorization, err := client.Do(requestAuthorization)
	if err != nil {
		return authorized, err
	}
	defer respAuthorization.Body.Close()

	type AuthorizationResponse struct {
		Authorized bool `json:"authorized"`
	}

	var responseAuthorized AuthorizationResponse

	err = json.NewDecoder(respAuthorization.Body).Decode(&responseAuthorized)
	authorized = responseAuthorized.Authorized

	return authorized, err
}

// Do http request to the authorization microservice and returns the authorization flag
// to update a transaction table in the database

func authorizeUpdateRequest(transaction Transaction) (bool, error) {
	var authorized bool
	var err error

	transactionJson, err := json.Marshal(transaction)

	if err != nil {
		return authorized, err
	}

	urlAuthoDelete := authorizationURL + "/update/"

	requestAuthorization, err := http.NewRequest("POST", urlAuthoDelete, bytes.NewBuffer(transactionJson))

	if err != nil {
		return authorized, err
	}

	client := http.Client{}
	requestAuthorization.Header.Set("Content-Type", "application/json")
	respAuthorization, err := client.Do(requestAuthorization)

	if err != nil {
		return authorized, err
	}
	defer respAuthorization.Body.Close()

	type AuthorizationResponse struct {
		Authorized bool `json:"authorized"`
	}

	var responseAuthorized AuthorizationResponse

	err = json.NewDecoder(respAuthorization.Body).Decode(&responseAuthorized)
	authorized = responseAuthorized.Authorized

	return authorized, err
}
