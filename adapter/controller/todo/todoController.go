package todo

import (
	"context"
	"graphql/graph/model"
	repositoryimpl "graphql/infrastructure/repositoryImpl"
	"graphql/usecase/todo"
)

type TodoController struct{}

func (ctrl TodoController) Show(ctx context.Context)([]*model.Todo,error){
	//ここに処理を書く
	todoRepository := repositoryimpl.NewTodoRepositoryImpl()
	result,err := todo.NewGetAllTodoUseCaseImpl(todoRepository).Get()
	return result,err
}