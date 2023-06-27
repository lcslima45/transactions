package authorization

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type authorizationHandler struct {
	service AuthorizationService
}

func NewAuthorizationHandler(authoService AuthorizationService) *authorizationHandler {
	return &authorizationHandler{
		service: authoService,
	}
}

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

	authorized, err := handler.service.AuthorizeCreation(transaction)

	if err != nil {
		http.Error(w, "An error ocurred during the authorization", http.StatusInternalServerError)
		return
	}

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

func (handler *authorizationHandler) AuthorizeDelete(w http.ResponseWriter, r *http.Request) {
	urlParts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(urlParts[len(urlParts)-1])

	if err != nil {
		http.Error(w, "Error on authorization delete path id", http.StatusBadRequest)
	}

	authorized, err := handler.service.AuthorizeDelete(id)

	if err != nil {
		http.Error(w, "An error ocurred during the authorization", http.StatusInternalServerError)
		return
	}

	responseAuthorized := &Authorization{
		Authorized: authorized,
	}

	w.Header().Set("Content-Type", "application-json")

	err = json.NewEncoder(w).Encode(responseAuthorized)

	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

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

	authorized, err := handler.service.AuthorizeUpdate(transaction)

	if err != nil {
		http.Error(w, "An error ocurred during the authorization", http.StatusInternalServerError)
		return
	}

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
