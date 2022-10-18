package todo

import (
	"context"
	"graphql/graph/model"
	repositoryimpl "graphql/infrastructure/repositoryImpl"
	"graphql/usecase/todo"
)

type TodoCreateController struct{}

func (ctrl TodoCreateController) Create(ctx context.Context,input *model.NewTodo)(*model.Todo,error){
	todoRepository := repositoryimpl.NewTodoRepositoryImpl()
	result,err := todo.NewCreateTodoUseCaseImpl(input,todoRepository).Create()
	return result,err
	
}