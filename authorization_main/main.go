package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/lcslima45/transactions/authorization"
	dBConfig "github.com/lcslima45/transactions/dbConfig"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Initializing authorization server...")
	//Connecting with the database
	db, err := sql.Open(dBConfig.PostgresDriver, dBConfig.DataSourceName)
	if err != nil {
		fmt.Println("Failed to connect with database:", err)
	}
	//Instatiating the repository of authorization
	repoAuthorization := authorization.NewAuthorizationRepository(db)
	//Instatiating the service of authorization
	servAuthorization := authorization.NewAuthorizationService(repoAuthorization)
	//Instatiating the handler of authorization
	authoHandler := authorization.NewAuthorizationHandler(servAuthorization)
	//Handle the authorization of creating new table
	http.HandleFunc("/authorization/add/", authoHandler.AuthorizeCreation)
	//Handle the authorization of deleting a table
	http.HandleFunc("/authorization/delete/", authoHandler.AuthorizeDelete)
	//Handle the authorization of updating a table
	http.HandleFunc("/authorization/update/", authoHandler.AuthorizeUpdate)
	http.ListenAndServe(":8888", nil)
}
