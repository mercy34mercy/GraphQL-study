package repository

import "graphql/graph/model"

type UserRepository interface{
	CreateUser(input model.NewUser)(*model.User,error)
}
