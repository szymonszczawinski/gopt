package auth

import (
	"context"
	"errors"
	"gosi/coreapi/service"

	"golang.org/x/sync/errgroup"
)

type IAuthService interface {
	service.IComponent
	login(username, password string) (AuthCredentials, error)
}

type authenticationService struct {
	ctx        context.Context
	eg         *errgroup.Group
	repository IAuthRepository
}

func NewAuthenticationService(eg *errgroup.Group, ctx context.Context, repository IAuthRepository) *authenticationService {
	instance := authenticationService{
		ctx:        ctx,
		eg:         eg,
		repository: repository,
	}
	return &instance
}

func (service *authenticationService) StartComponent() {

}

func (service authenticationService) login(username, pass string) (AuthCredentials, error) {
	userName, err := NewUserName(username)
	if err != nil {
		return AuthCredentials{}, err
	}
	password, err := NewPassword(pass)
	if err != nil {
		return AuthCredentials{}, err
	}
	authCredentials := NewAuthCredentials(userName, password)
	return authCredentials, errors.New("not implemented")
}
