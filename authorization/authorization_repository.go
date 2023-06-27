package authorization

import (
	"database/sql"
	"fmt"
	"strings"

	dbConfig "github.com/lcslima45/transactions/dbConfig"
)

// Struct that implements the Authorization Repository interface
type authorizationRepository struct {
	db *sql.DB
}

// Constructor of the authorization repository
func NewAuthorizationRepository(db *sql.DB) AuthorizationRepository {
	return &authorizationRepository{
		db: db,
	}
}

// Service to authorize creation
// If one Cardholder already exists with the same Brand and Issuer of the input the
// operation is not allowed.
func (repo *authorizationRepository) AuthorizeCreation(transaction Transaction) (bool, error) {
	var authorized bool
	sqlStatement := fmt.Sprintf("SELECT * FROM %s WHERE cardholder = $1 AND brand = $2 AND issuer = $3", dbConfig.TableName)
	query, err := repo.db.Query(sqlStatement, transaction.Cardholder, transaction.Brand, transaction.Issuer)

	if err != nil {
		return authorized, err
	}

	var count int
	fmt.Println(query)
	for query.Next() {
		err = query.Scan(&count)
		if err != nil {
			return authorized, err
		}
	}

	if count == 0 {
		authorized = true
	}

	return authorized, err

}

// Service to authorize deletion
// If the table is Brand == 'Visa"
// operation is not allowed.
func (repo *authorizationRepository) AuthorizeDelete(id int) (bool, error) {
	var authorize bool
	var err error
	sqlStatement := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", dbConfig.TableName)
	query, err := repo.db.Query(sqlStatement, id)
	var transaction Transaction
	for query.Next() {
		err = query.Scan(&transaction.ID,
			&transaction.Cardholder,
			&transaction.Merchant,
			&transaction.Acquirer,
			&transaction.Brand,
			&transaction.Issuer,
		)
	}

	if !(strings.ToLower(transaction.Brand) == "visa") {
		authorize = true
	}

	return authorize, err
}

// Service to authorize deletion
// If the table Cardholder is being altered
// operation is not allowed.
func (repo *authorizationRepository) AuthorizeUpdate(transaction Transaction) (bool, error) {
	var authorized bool
	var err error
	sqlStatement := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", dbConfig.TableName)
	query, err := repo.db.Query(sqlStatement, transaction.ID)
	if err != nil {
		return authorized, err
	}
	var transactionFromQuery Transaction
	for query.Next() {
		err = query.Scan(&transactionFromQuery.ID,
			&transactionFromQuery.Cardholder,
			&transactionFromQuery.Merchant,
			&transactionFromQuery.Acquirer,
			&transactionFromQuery.Brand,
			&transactionFromQuery.Issuer,
		)

		if err != nil {
			return authorized, nil
		}
	}

	if strings.ToLower(transactionFromQuery.Cardholder) == strings.ToLower(transaction.Cardholder) {
		authorized = true
	}

	return authorized, err
}
