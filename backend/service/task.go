package service

import (
	"todo/global"
	"todo/model"
	"todo/model/request"
	"todo/model/response"
)
func GetTasks (claims *model.CustomClaims)(tasks model.Tasks,err error){
	userId:=claims.ID

	tasks=make(model.Tasks,0)
	err=global.DB.Where("user_id = ? ",userId).Find(&tasks).Error

	return
}

func CreateTask(req request.CreateTask, claims *model.CustomClaims) (model.Task, error) {
	falsePtr := false
	task := model.Task{
		Text:     req.Text,
		DDL:      req.DDL,
		Finished: &falsePtr,
		UserId:   claims.ID,
	}
	return task, global.DB.Create(&task).Error
}

func UpdateTask(req request.UpdateTask, claims *model.CustomClaims) error {
	task :=model.Task{
		BasicModel: model.BasicModel{
			ID: req.ID,
		},
		Text:       req.Text,
		DDL:        req.DDL,
		Finished:   req.Finished,
	}
	return global.DB.Where("id = ? AND user_id = ?",req.ID,claims.ID).Updates(task).Error
}

func NewTaskResp(task *model.Task) (resp response.TaskResp) {

	resp = response.TaskResp{
		ID:       task.ID,
		Text:     task.Text,
		Finished: *task.Finished,
		CreatedAt: task.CreatedAt,
	}

	if task.DDL != nil {
		resp.DDL = task.DDL.Local().Format("2006-01-02 15:04")
	}

	return resp
}
