package transactions

type transactionService struct {
	repository Repository
}

func NewTransactionService(repo Repository) Service {
	return &transactionService{
		repository: repo,
	}
}

func (s *transactionService) CreateTransaction(transaction Transaction) error {
	return s.repository.SaveTransaction(transaction)
}

func (s *transactionService) DeleteTransaction(id int) error {
	return s.repository.DeleteTransaction(id)
}

func (s *transactionService) UpdateTransaction(id int, transaction Transaction) error {
	return s.repository.UpdateTransaction(id, transaction)
}

func (s *transactionService) GetTransactionsPagination(limit, offset int) ([]Transaction, error) {
	return s.repository.GetTransactionsPagination(limit, offset)
}
