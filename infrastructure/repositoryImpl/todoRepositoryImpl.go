package repositoryimpl

import (
	"fmt"
	"graphql/domain/repository"
	"graphql/graph/model"
	"graphql/infrastructure"
	"graphql/infrastructure/dto"
	"strconv"
)

type TodoRepositoryImpl struct {}

func NewTodoRepositoryImpl() repository.TodoRepository{
	return &TodoRepositoryImpl{}
}

func (repositoryImpl *TodoRepositoryImpl) FindTodo()([]*model.Todo,error) {
	db := infrastructure.GetDB()
	var err error
	var Modeltodo []*model.Todo = []*model.Todo{}
	var todo []*dto.Todo = []*dto.Todo{}
	var User *model.User = &model.User{}

	if err = db.Find(&todo).Error; err != nil {
		//エラーハンドリング
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}

	for i := 0 ; i <  len(todo); i++ {
		if err = db.Where("id = ?",todo[i].UserID).First(&User).Error; err != nil {
			//エラーハンドリング
			fmt.Printf("db select Error!!!! err:%v\n", err)
		}
		Modeltodo = append(Modeltodo, 
			&model.Todo{
				ID:strconv.FormatUint(uint64(todo[i].ID),10),
				Text: todo[i].Text,
				Done: todo[i].Done,
				User: User,
			},
		)
	}

	return Modeltodo,err
}


func (repositoryImpl *TodoRepositoryImpl) CreateTodo(input model.NewTodo)(*model.Todo,error){
	db := infrastructure.GetDB()
	var err error
	var user *model.User = &model.User{}

	UserID_parseuint,_ := strconv.ParseUint(input.UserID,10,64)

	//gormを記述
	if err = db.Create(&dto.Todo{UserID: uint(UserID_parseuint),Text: input.Text,Done: false}).Error; err != nil {
		fmt.Printf("todo create Error!!!! err:%v\n", err)
	}

	if err = db.Where("id = ?",input.UserID).First(&user).Error; err != nil {
		fmt.Printf("db select Error!!!! err:%v\n", err)
	}

	return &model.Todo{
		ID: input.UserID,
		Text: input.Text,
		Done: false,
		User: user,
	},err
}