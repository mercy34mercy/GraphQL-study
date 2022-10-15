package todo

import (
	"context"
	"graphql/graph/model"
)

type TodoController struct{}

func (ctrl TodoController) Show(ctx context.Context)([]*model.Todo){
	//ここに処理を書く
}