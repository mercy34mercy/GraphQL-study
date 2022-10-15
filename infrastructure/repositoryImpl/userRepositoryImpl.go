package repositoryimpl

import (
	"graphql/domain/repository"
	"graphql/graph/model"
	"graphql/infrastructure"
)

type UserRepositoryImpl struct {}

func NewUserRepositoryImpl() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (repositoryImpl *UserRepositoryImpl) CreateUser(input model.NewUser)(*model.User,error){
	db := infrastructure.GetDB()
	var err error
	var user *model.User = &model.User{}

	//gormを記述

	return user,err
}