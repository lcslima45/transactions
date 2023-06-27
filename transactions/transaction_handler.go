package transactions

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type transactionHandler struct {
	service Service
}

func NewTransactionHandler(service Service) *transactionHandler {
	handler := &transactionHandler{
		service: service,
	}
	return handler
}

func (h *transactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	//Decodificar o corpo da request que é um json onde está inserido a transação a ser inserida no banco
	var transaction Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	//Chama o serviço de transação para criar a transação
	authorized, err := authorizeCreationRequest(transaction)

	if err != nil {
		http.Error(w, "Error in authorization", http.StatusInternalServerError)
		return
	}

	if !authorized {
		http.Error(w, "Transaction unauthorized", http.StatusForbidden)
		return
	}

	err = h.service.CreateTransaction(transaction)
	if err != nil {
		http.Error(w, "Failed in create transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Transaction succesfully inserted!!!"))
}

func (h *transactionHandler) DeleteTransaction(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed should be DELETE", http.StatusMethodNotAllowed)
		return
	}

	urlParts := strings.Split(r.URL.Path, "/")
	id, err := strconv.Atoi(urlParts[len(urlParts)-1])
	if err != nil {
		http.Error(w, "Error on the path id", http.StatusBadRequest)
		return
	}

	authorized, err := authorizeDeleteRequest(id)

	if !authorized {
		http.Error(w, "Deletion unauthorized", http.StatusForbidden)
		return
	}

	if err != nil {
		http.Error(w, "Error in authorization", http.StatusInternalServerError)
		return
	}

	err = h.service.DeleteTransaction(id)

	if err != nil {
		http.Error(w, "Failed in delete transaction", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transaction succesfully deleted!!!"))
}

func (h *transactionHandler) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	var id int
	var transaction Transaction
	if r.Method != "PUT" {
		http.Error(w, "Method not allowed should be UPDATE", http.StatusMethodNotAllowed)
		return
	}
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

	authorized, err := authorizeUpdateRequest(transaction)

	if !authorized {
		http.Error(w, "Update unauthorized", http.StatusForbidden)
		return
	}

	if err != nil {
		http.Error(w, "Error in authorization", http.StatusInternalServerError)
		return
	}

	err = h.service.UpdateTransaction(id, transaction)
	if err != nil {
		http.Error(w, "Failed in update transaction", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transaction succesfully updated!!!"))
}

func (h *transactionHandler) ListTransactionPagination(w http.ResponseWriter, r *http.Request) {
	var limit, offset int
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

	offset = (page - 1) * limit

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
