package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	dbConfig "github.com/lcslima45/transactions/dbConfig"
	transactions "github.com/lcslima45/transactions/transactions"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Inicializing transaction server...")
	db, err := sql.Open(dbConfig.PostgresDriver, dbConfig.DataSourceName)
	if err != nil {
		log.Fatal("Failed to connect with databse:", err)
	}
	defer db.Close()
	transactionRepo := transactions.NewTransactionRepository(db)
	transactionService := transactions.NewTransactionService(transactionRepo)
	transactionHandler := transactions.NewTransactionHandler(transactionService)
	http.HandleFunc("/transactions/add", transactionHandler.CreateTransaction)
	http.HandleFunc("/transactions/delete/", transactionHandler.DeleteTransaction)
	http.HandleFunc("/transactions/update/", transactionHandler.UpdateTransaction)
	http.Handle("/transactions/list/", http.HandlerFunc(transactionHandler.ListTransactionPagination))
	http.ListenAndServe(":8080", nil)
}
