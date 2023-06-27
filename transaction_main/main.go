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
	//Connect with database Postgres
	db, err := sql.Open(dbConfig.PostgresDriver, dbConfig.DataSourceName)
	if err != nil {
		log.Fatal("Failed to connect with databse:", err)
	}
	defer db.Close()
	//Instantiating the repository
	transactionRepo := transactions.NewTransactionRepository(db)
	//Instantiating the service
	transactionService := transactions.NewTransactionService(transactionRepo)
	//Instantiating the handler
	transactionHandler := transactions.NewTransactionHandler(transactionService)
	//Route to the routine of creation of a new transaction
	http.HandleFunc("/transactions/add", transactionHandler.CreateTransaction)
	//Route to the routine of deletion of a transaction
	http.HandleFunc("/transactions/delete/", transactionHandler.DeleteTransaction)
	//Route to the routine of update of a transaction
	http.HandleFunc("/transactions/update/", transactionHandler.UpdateTransaction)
	//Route to the routine of listing the paginated transactions
	http.Handle("/transactions/list/", http.HandlerFunc(transactionHandler.ListTransactionPagination))
	http.ListenAndServe(":8080", nil)
}
