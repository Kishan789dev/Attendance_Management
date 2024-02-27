package tokenvalid

type AuthenticationSvc interface {
	ValidateTokenAndGetEmailSvc(email string) error
}

type AuthenticationSvcImpl struct {
	authenticationrepo AuthenticationRepo
}

func NewAuthenticationSvc(authenticationrepo AuthenticationRepo) *AuthenticationSvcImpl {
	return &AuthenticationSvcImpl{
		authenticationrepo: authenticationrepo,
	}
}

func (impl *AuthenticationSvcImpl) ValidateTokenAndGetEmailSvc(email string) error {
	return impl.authenticationrepo.ValidateTokenAndGetEmailRepo(email)
}
