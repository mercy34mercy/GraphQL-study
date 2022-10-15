package repository

import "graphql/graph/model"

type TodoRepository interface{
	FindTodo()([]*model.Todo,error)
	CreateTodo(input model.NewTodo)(*model.Todo,error)
}