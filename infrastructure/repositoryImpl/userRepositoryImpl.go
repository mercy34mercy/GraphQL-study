package repositoryimpl

import (
	"fmt"
	"graphql/domain/repository"
	"graphql/graph/model"
	"graphql/infrastructure"
	"graphql/infrastructure/dto"
	"strconv"
)

type UserRepositoryImpl struct {}

func NewUserRepositoryImpl() repository.UserRepository {
	return &UserRepositoryImpl{}
}

func (repositoryImpl *UserRepositoryImpl) CreateUser(input model.NewUser)(*model.User,error){
	db := infrastructure.GetDB()
	var err error
	var user *dto.User = &dto.User{}

	//gormを記述
	if err = db.Create(&dto.User{Name: input.Name}).Error; err != nil {
		fmt.Printf("user create Error!!!! err:%v\n", err)
	}

	if err = db.Where("name = ?",input.Name).Order("created_at").Find(&user).Error; err != nil {
		fmt.Printf("user create Error!!!! err:%v\n", err)
	}

	

	return &model.User{
		Name: input.Name,
		ID: strconv.FormatUint(uint64(user.ID),10),
	},err
}

