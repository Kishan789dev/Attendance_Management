package tokenvalid

func ValidateTokenAndGetEmailSvc(email string) error {
	return ValidateTokenAndGetEmailRepo(email)
}
