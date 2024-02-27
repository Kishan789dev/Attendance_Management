package login

type LoginSvc interface {
	LoginSvc(useremail string) (error, string)
}

type LoginSvcImpl struct {
	loginrepo LoginRepo
}

func NewLoginSvc(loginrepo LoginRepo) *LoginSvcImpl {
	return &LoginSvcImpl{
		loginrepo: loginrepo,
	}
}

func (impl *LoginSvcImpl) LoginSvc(useremail string) (error, string) {
	return impl.loginrepo.LoginRepo(useremail)

}
