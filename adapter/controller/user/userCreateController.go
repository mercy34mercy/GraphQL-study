package user

import (
	"context"
	"graphql/graph/model"
	repositoryimpl "graphql/infrastructure/repositoryImpl"
	"graphql/usecase/user"
)

type UserCreateController struct{}

func (ctrl UserCreateController) Create(ctx context.Context,input *model.NewUser)(*model.User,error){
	userRepository := repositoryimpl.NewUserRepositoryImpl()
	result,err := user.NewCreateUserUseCaseImpl(input,userRepository).Create()
	return result,err
}