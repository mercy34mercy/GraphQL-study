package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	todoCtrl "graphql/adapter/controller/todo"
	userCtrl "graphql/adapter/controller/user"
	"graphql/graph/generated"
	"graphql/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	userCtrl := userCtrl.UserCreateController{}
	User, err := userCtrl.Create(ctx, &input)
	return User, err
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todoCtrl := todoCtrl.TodoCreateController{}
	Todo, err := todoCtrl.Create(ctx, &input)
	return Todo, err
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	todoCtrl := todoCtrl.TodoController{}
	Todo, err := todoCtrl.Show(ctx)
	return Todo, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
