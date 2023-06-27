package authorization

type authorizationService struct {
	repo AuthorizationRepository
}

func NewAuthorizationService(repo AuthorizationRepository) AuthorizationService {
	return &authorizationService{
		repo: repo,
	}
}

func (service *authorizationService) AuthorizeCreation(transaction Transaction) (bool, error) {
	return service.repo.AuthorizeCreation(transaction)
}

func (service *authorizationService) AuthorizeDelete(id int) (bool, error) {
	return service.repo.AuthorizeDelete(id)
}

func (service *authorizationService) AuthorizeUpdate(transaction Transaction) (bool, error) {
	return service.repo.AuthorizeUpdate(transaction)
}
