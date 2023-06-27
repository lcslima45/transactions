package transactions

//Model of the transaction that is in the database
type Transaction struct {
	ID         int
	Cardholder string
	Merchant   string
	Acquirer   string
	Brand      string
	Issuer     string
}

//Interface that implements the service that is called by the handler and manipulates the Repository
type Service interface {
	CreateTransaction(transaction Transaction) error
	DeleteTransaction(id int) error
	UpdateTransaction(id int, transaction Transaction) error
	GetTransactionsPagination(limit, offset int) ([]Transaction, error)
}

//Interface that implements the repository that manipulates the activity of the CRUD on the database
type Repository interface {
	SaveTransaction(transaction Transaction) error
	DeleteTransaction(id int) error
	UpdateTransaction(id int, transaction Transaction) error
	GetTransactionsPagination(limit, offset int) ([]Transaction, error)
}
