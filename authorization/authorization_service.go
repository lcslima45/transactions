package authorization

//Struct that implements the Authorization service interface
type authorizationService struct {
	repo AuthorizationRepository
}

//Constructor of the AuthorizationService
func NewAuthorizationService(repo AuthorizationRepository) AuthorizationService {
	return &authorizationService{
		repo: repo,
	}
}

//Call to the repository to authorize creation
func (service *authorizationService) AuthorizeCreation(transaction Transaction) (bool, error) {
	return service.repo.AuthorizeCreation(transaction)
}

//Call to the repository to authorize deletion
func (service *authorizationService) AuthorizeDelete(id int) (bool, error) {
	return service.repo.AuthorizeDelete(id)
}

//Call to the repository to authorize update
func (service *authorizationService) AuthorizeUpdate(transaction Transaction) (bool, error) {
	return service.repo.AuthorizeUpdate(transaction)
}
