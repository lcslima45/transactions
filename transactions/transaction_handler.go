package transactions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Service to handle the http requests on the Transactions CRUD
type transactionHandler struct {
	service Service
}

// Constructor of the handler
func NewTransactionHandler(service Service) *transactionHandler {
	handler := &transactionHandler{
		service: service,
	}
	return handler
}

// Handler of the routine to add a new table on the database
func (h *transactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//HTTP request to authorize the adding of a new table
	authorized, err := authorizeCreationRequest(transaction)

	if err != nil {
		http.Error(w, "Error in authorization", http.StatusInternalServerError)
		return
	}

	if !authorized {
		http.Error(w, "Transaction unauthorized", http.StatusForbidden)
		return
	}
	//Call the service to the repository add the new table
	err = h.service.CreateTransaction(transaction)
	if err != nil {
		http.Error(w, "Failed in create transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Transaction succesfully inserted!!!"))
}

// Handler of the routine to delete a table on the database
func (h *transactionHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed should be DELETE", http.StatusMethodNotAllowed)
		return
	}
	//search the Id of the table that the user wants to delete in the URL
	urlParts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(urlParts[len(urlParts)-1])
	if err != nil {
		http.Error(w, "Error on the path id", http.StatusBadRequest)
		return
	}

	//Do the HTTP request to delete the table of the database
	authorized, err := authorizeDeleteRequest(id)

	if !authorized {
		http.Error(w, "Deletion unauthorized", http.StatusForbidden)
		return
	}

	if err != nil {
		http.Error(w, "Error in authorization", http.StatusInternalServerError)
		return
	}

	//Call the service to delete the transaction
	err = h.service.DeleteTransaction(id)

	if err != nil {
		http.Error(w, "Failed in delete transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transaction succesfully deleted!!!"))
}

// Handler of the routine to UPDATE a table on the database
func (h *transactionHandler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var id int
	var transaction Transaction
	if r.Method != "PUT" {
		http.Error(w, "Method not allowed should be UPDATE", http.StatusMethodNotAllowed)
		return
	}
	//search the ID of the transaction in the URL
	urlParts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(urlParts[len(urlParts)-1])
	if err != nil {
		http.Error(w, "Error on the path id", http.StatusBadRequest)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	//Do the request to authorize the update
	authorized, err := authorizeUpdateRequest(transaction)

	if !authorized {
		http.Error(w, "Update unauthorized", http.StatusForbidden)
		return
	}

	if err != nil {
		http.Error(w, "Error in authorization", http.StatusInternalServerError)
		return
	}
	//Do the update on the database
	err = h.service.UpdateTransaction(id, transaction)
	if err != nil {
		http.Error(w, "Failed in update transaction", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transaction succesfully updated!!!"))
}

// Handler of the routine to list a table on the database
func (h *transactionHandler) ListTransactionPagination(w http.ResponseWriter, r *http.Request) {
	var limit, offset int
	//search the page and limit in the url
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)

	if err != nil {
		http.Error(w, "Invalid page", http.StatusBadRequest)
		return
	}

	limit, err = strconv.Atoi(limitStr)
	if err != nil {
		http.Error(w, "Invalid page", http.StatusBadRequest)
		return
	}
	//Operation to generate the offset of pagination
	offset = (page - 1) * limit
	//Calls the service to get the list of transactions by pages
	transactions, err := h.service.GetTransactionsPagination(limit, offset)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error listing pagination", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		http.Error(w, "Error failed to encode response", http.StatusInternalServerError)
		return
	}
}
