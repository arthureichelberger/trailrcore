package exception

type CouldNotValidatePasswordError struct{}

func (cnvp CouldNotValidatePasswordError) Error() string {
	return "could not validate password"
}

type CouldNotValidateEmailError struct{}

func (cnve CouldNotValidateEmailError) Error() string {
	return "could not validate email"
}
