package authorization

type Authorization struct {
	Authorized bool `json:"authorized"`
}

type Transaction struct {
	ID         int
	Cardholder string
	Merchant   string
	Acquirer   string
	Brand      string
	Issuer     string
}

type AuthorizationService interface {
	AuthorizeCreation(transaction Transaction) (bool, error)
	AuthorizeDelete(id int) (bool, error)
	AuthorizeUpdate(transaction Transaction) (bool, error)
}

type AuthorizationRepository interface {
	AuthorizeCreation(transaction Transaction) (bool, error)
	AuthorizeDelete(id int) (bool, error)
	AuthorizeUpdate(transaction Transaction) (bool, error)
}
