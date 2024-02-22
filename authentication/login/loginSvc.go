package login

func LoginSvc(useremail string) (error, string) {
	return LoginRepo(useremail)

}
