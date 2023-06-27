package transactions

//Struct that implements the Service Interface
type transactionService struct {
	repository Repository
}

//Constructor of the Service
func NewTransactionService(repo Repository) Service {
	return &transactionService{
		repository: repo,
	}
}

//Returns the response of the routine to save a new transaction of the repository
func (s *transactionService) CreateTransaction(transaction Transaction) error {
	return s.repository.SaveTransaction(transaction)
}

//Returns the response of the routine to delete a transaction of the repository
func (s *transactionService) DeleteTransaction(id int) error {
	return s.repository.DeleteTransaction(id)
}

//Returns the response of the routine to update a transaction of the repository
func (s *transactionService) UpdateTransaction(id int, transaction Transaction) error {
	return s.repository.UpdateTransaction(id, transaction)
}

//Returns the response of the routine to list the paginated transactions of the repository
func (s *transactionService) GetTransactionsPagination(limit, offset int) ([]Transaction, error) {
	return s.repository.GetTransactionsPagination(limit, offset)
}
