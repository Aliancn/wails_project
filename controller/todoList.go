package controller

import (
	"changeme/Dao"
	"changeme/utils"
)

type TodoList struct {
	Id       uint   `json:"id"`
	TaskName string `json:"taskName"`
	Competed bool   `json:"competed"`
}

func (a *App) GetTodoList() *utils.Resp {
	var result []TodoList
	dao := Dao.TodoListDao{}

	dbresult, err := dao.GetAll(DB)
	if err != nil {
		return utils.Error("get todo list error")
	}
	for i := range dbresult {
		var ttmp = TodoList{
			Id:       dbresult[i].ID,
			TaskName: dbresult[i].TaskName,
			Competed: dbresult[i].Competed,
		}
		result = append(result, ttmp)
	}
	return utils.Success(result)
}

func (a *App) CreateTodoList(tl TodoList) *utils.Resp {
	var ttmp = Dao.TodoListDao{
		ID:       tl.Id,
		TaskName: tl.TaskName,
		Competed: tl.Competed,
	}
	err := ttmp.Create(DB)
	if err != nil {
		return utils.Error("create todo list error")
	}
	return utils.Success(nil)
}

func (a *App) UpdateTodoList(tl TodoList) *utils.Resp {
	var ttmp = Dao.TodoListDao{
		ID:       tl.Id,
		TaskName: tl.TaskName,
		Competed: tl.Competed,
	}
	// todo 能不能完成更新操作
	err := ttmp.Update(DB)
	if err != nil {
		return utils.Error("update todo list error")
	}
	return utils.Success(nil)
}

func (a *App) DeleteTodoList(tl TodoList) *utils.Resp {
	var ttmp = Dao.TodoListDao{
		ID: tl.Id,
	}
	err := ttmp.Delete(DB)
	if err != nil {
		return utils.Error("delete todo list error")
	}
	return utils.Success(nil)
}
