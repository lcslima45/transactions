package authorization

//struct with the Authorization flag
type Authorization struct {
	Authorized bool `json:"authorized"`
}

//Struct to manipulate transactions inside the Authorization escope
type Transaction struct {
	ID         int
	Cardholder string
	Merchant   string
	Acquirer   string
	Brand      string
	Issuer     string
}

//Interface of the Authorization Service
type AuthorizationService interface {
	AuthorizeCreation(transaction Transaction) (bool, error)
	AuthorizeDelete(id int) (bool, error)
	AuthorizeUpdate(transaction Transaction) (bool, error)
}

//Interface of the Authorization Repository
type AuthorizationRepository interface {
	AuthorizeCreation(transaction Transaction) (bool, error)
	AuthorizeDelete(id int) (bool, error)
	AuthorizeUpdate(transaction Transaction) (bool, error)
}
