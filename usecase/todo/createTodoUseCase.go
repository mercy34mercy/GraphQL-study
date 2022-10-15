package todo

import (
	"graphql/domain/repository"
	"graphql/graph/model"
)

type CreateTodoUseCaseImpl struct {
	NewTodo *model.NewTodo
	TodoRepository repository.TodoRepository
}

type CreateTodoUseCase interface {
	Create()(*model.Todo,error)
}

func NewCreateTodoUseCaseImpl(input *model.NewTodo, todorepository repository.TodoRepository) CreateTodoUseCase {
	return CreateTodoUseCaseImpl{
		NewTodo: input,
		TodoRepository: todorepository,
	}
}

func(impl CreateTodoUseCaseImpl) Create() (*model.Todo,error){
	//ここに処理を記述

}
