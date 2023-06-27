package transactions

type Transaction struct {
	ID         int
	Cardholder string
	Merchant   string
	Acquirer   string
	Brand      string
	Issuer     string
}

type Service interface {
	CreateTransaction(transaction Transaction) error
	DeleteTransaction(id int) error
	UpdateTransaction(id int, transaction Transaction) error
	GetTransactionsPagination(limit, offset int) ([]Transaction, error)
}

type Repository interface {
	SaveTransaction(transaction Transaction) error
	DeleteTransaction(id int) error
	UpdateTransaction(id int, transaction Transaction) error
	GetTransactionsPagination(limit, offset int) ([]Transaction, error)
}
