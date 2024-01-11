package auth

import "errors"

var (
	ErrorInvalidUsername = errors.New("invalid username provided")
	ErrorInvalidPassword = errors.New("invalid password provided")
)

type AuthCredentials struct {
	Username Username
	Password Password
}

func NewAuthCredentials(username Username, password Password) AuthCredentials {
	return AuthCredentials{
		Username: username,
		Password: password,
	}
}

type Username string

func NewUserName(username string) (Username, error) {
	if len(username) == 0 {
		return "", errors.Join(errors.New("Username can not be epmty"), ErrorInvalidUsername)
	}
	return Username(username), nil
}

func (u Username) String() string {
	return string(u)
}

type Password string

func NewPassword(password string) (Password, error) {
	if len(password) == 0 {
		return "", errors.Join(errors.New("Password can not be epmty"), ErrorInvalidPassword)
	}
	return Password(password), nil
}

func (p Password) String() string {
	return string(p)
}
