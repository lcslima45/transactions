package authorization

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Struct of the handler of authorization
type authorizationHandler struct {
	service AuthorizationService
}

// Constructor of the handler of authorization
func NewAuthorizationHandler(authoService AuthorizationService) *authorizationHandler {
	return &authorizationHandler{
		service: authoService,
	}
}

// Handler of the creation of new table
func (handler *authorizationHandler) AuthorizeCreation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received:", r.Method)
	if r.Method != "POST" {
		http.Error(w, "Method not allowed should be POST", http.StatusMethodNotAllowed)
		return
	}

	var transaction Transaction

	err := json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	//Calls the service to manipulate the repository
	authorized, err := handler.service.AuthorizeCreation(transaction)

	if err != nil {
		http.Error(w, "An error ocurred during the authorization", http.StatusInternalServerError)
		return
	}
	//Returns the authorization flag to the transaction microservice
	responseAuthorized := &Authorization{
		Authorized: authorized,
	}

	w.Header().Set("Content-Type", "application-json")

	err = json.NewEncoder(w).Encode(responseAuthorized)

	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

}

// Handler of the deletion of a table
func (handler *authorizationHandler) AuthorizeDelete(w http.ResponseWriter, r *http.Request) {
	//searching the Id of the table that will be excluded
	urlParts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(urlParts[len(urlParts)-1])

	if err != nil {
		http.Error(w, "Error on authorization delete path id", http.StatusBadRequest)
	}
	//Calls the Service to manipulates the repository and do the deletion
	authorized, err := handler.service.AuthorizeDelete(id)

	if err != nil {
		http.Error(w, "An error ocurred during the authorization", http.StatusInternalServerError)
		return
	}
	//Returns the authorization flag to the transaction microservice

	responseAuthorized := &Authorization{
		Authorized: authorized,
	}

	w.Header().Set("Content-Type", "application-json")

	err = json.NewEncoder(w).Encode(responseAuthorized)

	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

// Handler of the update of a table

func (handler *authorizationHandler) AuthorizeUpdate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received:", r.Method)
	if r.Method != "POST" {
		http.Error(w, "Method not allowed should be POST", http.StatusMethodNotAllowed)
		return
	}

	var transaction Transaction

	err := json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	//Calls the service to manipulate the repository and do the new updating
	authorized, err := handler.service.AuthorizeUpdate(transaction)

	if err != nil {
		http.Error(w, "An error ocurred during the authorization", http.StatusInternalServerError)
		return
	}
	//Return the authorization flag as response to the transactions microservice
	responseAuthorized := &Authorization{
		Authorized: authorized,
	}

	w.Header().Set("Content-Type", "application-json")

	err = json.NewEncoder(w).Encode(responseAuthorized)

	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
