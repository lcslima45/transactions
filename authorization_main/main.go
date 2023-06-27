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
	db, err := sql.Open(dBConfig.PostgresDriver, dBConfig.DataSourceName)
	if err != nil {
		fmt.Println("Failed to connect with database:", err)
	}
	repoAuthorization := authorization.NewAuthorizationRepository(db)
	servAuthorization := authorization.NewAuthorizationService(repoAuthorization)
	authoHandler := authorization.NewAuthorizationHandler(servAuthorization)
	http.HandleFunc("/authorization/add/", authoHandler.AuthorizeCreation)
	http.HandleFunc("/authorization/delete/", authoHandler.AuthorizeDelete)
	http.HandleFunc("/authorization/update/", authoHandler.AuthorizeUpdate)
	http.ListenAndServe(":8888", nil)
}
