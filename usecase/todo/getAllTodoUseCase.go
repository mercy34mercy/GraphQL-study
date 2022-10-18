package todo

import (
	"graphql/domain/repository"
	"graphql/graph/model"
)


type getAllTodoUseCaseImpl struct {
	TodoRepository repository.TodoRepository
}

type getAllTodoUseCase interface {
	Get()([]*model.Todo,error)
}

func NewGetAllTodoUseCaseImpl(todorepository repository.TodoRepository) getAllTodoUseCase {
	return getAllTodoUseCaseImpl{
		TodoRepository: todorepository,
	}
}

func (impl getAllTodoUseCaseImpl) Get()([]*model.Todo,error){
	//ここに処理を書く
	todo,err := impl.TodoRepository.FindTodo()

	return todo,err
}