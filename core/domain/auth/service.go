package auth

import (
	"context"
	"errors"
	"gopt/coreapi"
	"gopt/coreapi/service"

	"golang.org/x/sync/errgroup"
)

type IAuthService interface {
	service.IComponent
	login(username, password string) coreapi.Result[AuthCredentials]
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

func (service authenticationService) login(username, pass string) coreapi.Result[AuthCredentials] {
	userName, err := NewUserName(username)
	if err != nil {
		return coreapi.NewResult[AuthCredentials](AuthCredentials{}, err)
	}
	password, err := NewPassword(pass)
	if err != nil {
		return coreapi.NewResult[AuthCredentials](AuthCredentials{}, err)
	}
	authCredentials := NewAuthCredentials(userName, password)
	return coreapi.NewResult[AuthCredentials](authCredentials, errors.New("not implemented"))
}
