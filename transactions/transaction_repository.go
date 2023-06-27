package transactions

import (
	"database/sql"
	"fmt"

	dbConfig "github.com/lcslima45/transactions/dbConfig"
	_ "github.com/lib/pq"
)

// struct that implements the Repository interface
type transactionRepository struct {
	db *sql.DB
}

// Constructor of the Repository
func NewTransactionRepository(db *sql.DB) Repository {
	return &transactionRepository{
		db: db,
	}
}

// Routine to save a new authorized transaction
func (repo *transactionRepository) SaveTransaction(transaction Transaction) error {
	sqlStatement := fmt.Sprintf("INSERT INTO %s (cardholder, merchant, acquirer, brand, issuer) VALUES ($1,$2,$3,$4,$5)", dbConfig.TableName)
	insert, err := repo.db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	defer insert.Close()
	//Inserting the new transaction on the database
	result, err := insert.Exec(transaction.Cardholder,
		transaction.Merchant,
		transaction.Acquirer,
		transaction.Brand,
		transaction.Issuer,
	)
	if err != nil {
		return err
	}
	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println("Rows affected:", affect)
	return nil
}

// Routine to do an authorized deletion
func (repo *transactionRepository) DeleteTransaction(id int) error {
	sqlStatement := fmt.Sprintf("Delete from %s where id=$1", dbConfig.TableName)
	delete, err := repo.db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	defer delete.Close()
	result, err := delete.Exec(id)
	if err != nil {
		return err
	}
	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println("Rows affected:", affect)
	return nil
}

// Routine to do an authorized update
func (repo *transactionRepository) UpdateTransaction(id int, transaction Transaction) error {
	sqlStatement := fmt.Sprintf("UPDATE %s SET cardholder=$1, merchant=$2, acquirer=$3, brand=$4, issuer=$5 WHERE id=$6", dbConfig.TableName)
	update, err := repo.db.Prepare(sqlStatement)
	if err != nil {
		return err
	}
	defer update.Close()
	result, err := update.Exec(transaction.Cardholder,
		transaction.Merchant,
		transaction.Acquirer,
		transaction.Brand,
		transaction.Issuer,
		id,
	)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	fmt.Println("Rows affected:", rows)
	return nil
}

// Routine that returns the pages of the http Request on the listing handler
func (repo *transactionRepository) GetTransactionsPagination(limit, offset int) ([]Transaction, error) {
	sqlStatement := fmt.Sprintf(`SELECT * FROM %s OFFSET $1 LIMIT $2`, dbConfig.TableName)

	rows, err := repo.db.Query(sqlStatement, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	transactions := []Transaction{}

	for rows.Next() {
		var transaction Transaction

		err = rows.Scan(&transaction.ID,
			&transaction.Cardholder,
			&transaction.Merchant,
			&transaction.Acquirer,
			&transaction.Brand,
			&transaction.Issuer,
		)

		transactions = append(transactions, transaction)
	}

	err = rows.Err()

	if err != nil {
		return nil, err
	}

	return transactions, nil
}
