package auth

import (
	"context"
	"gopt/coreapi"

	"golang.org/x/sync/errgroup"
)

type IAuthRepository interface {
	coreapi.IComponent
}

type authenticationService struct {
	ctx        context.Context
	eg         *errgroup.Group
	repository IAuthRepository
}

func NewAuthenticationService(eg *errgroup.Group, ctx context.Context, repository IAuthRepository) authenticationService {
	instance := authenticationService{
		ctx:        ctx,
		eg:         eg,
		repository: repository,
	}
	return instance
}

func (service authenticationService) StartComponent() {
}

func (service authenticationService) Login(username, pass string) coreapi.Result[AuthCredentials] {
	userName, err := NewUserName(username)
	if err != nil {
		return coreapi.NewResult(AuthCredentials{}, err)
	}
	password, err := NewPassword(pass)
	if err != nil {
		return coreapi.NewResult(AuthCredentials{}, err)
	}
	authCredentials := NewAuthCredentials(userName, password)
	// return coreapi.NewResult(authCredentials, errors.New("not implemented"))
	return coreapi.NewResult(authCredentials, nil)
}
