package auth

import (
	"context"
	"gosi/coreapi/service"

	"golang.org/x/sync/errgroup"
)

type IAuthService interface {
	service.IComponent
	login(CredentialsData) (UserCredentials, error)
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

func (self *authenticationService) StartComponent() {

}

func (self authenticationService) login(credentialData CredentialsData) (UserCredentials, error) {
	panic("Not Implemented")
}
