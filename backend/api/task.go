package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nc-77/gtils"
	_ "gorm.io/gorm"
	"net/http"
	"time"
	"todo/global"
	"todo/model"
	"todo/model/request"
	"todo/model/response"
	"todo/service"
)

// @Tags Task
// @Summary 创建任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.CreateTask true "任务内容,截止日期"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tasks [put]
func CreateTask(c *gin.Context) {
	var taskReq request.CreateTask
	if err := c.ShouldBindJSON(&taskReq); err != nil {
		gtils.FailResponse(c, http.StatusBadRequest, err)
		return
	}

	customClaims, _ := c.Get("claims")
	claims := customClaims.(*model.CustomClaims)

	_, err := service.CreateTask(taskReq, claims)
	if err != nil {
		gtils.FailResponse(c, http.StatusInternalServerError, err)
	} else {
		gtils.SucResponse(c, "创建成功", "")
	}

}

// @Tags Task
// @Summary 删除任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DeleteTask true "任务id"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tasks [delete]
func DeleteTask(c *gin.Context) {
	var task request.DeleteTask
	if err := c.ShouldBindJSON(&task); err != nil {
		gtils.FailResponse(c, http.StatusBadRequest, err)
		return
	}

	customClaims, _ := c.Get("claims")
	claims := customClaims.(*model.CustomClaims)

	if err := global.DB.Where("user_id = ?", claims.ID).Delete(&model.Task{}, task.IDs).Error; err != nil {
		gtils.FailResponse(c, http.StatusInternalServerError, global.ErrDeleted)
	} else {
		gtils.SucResponse(c, "删除成功", "")
	}

}

// @Tags Task
// @Summary 更新任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UpdateTask true "任务id,任务内容,截止日期，完成状态"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tasks [post]
func UpdateTask(c *gin.Context) {
	var taskReq request.UpdateTask
	if err := c.ShouldBindJSON(&taskReq); err != nil {
		gtils.FailResponse(c, http.StatusBadRequest, err)
		return
	}

	customClaims, _ := c.Get("claims")
	claims := customClaims.(*model.CustomClaims)

	if err := service.UpdateTask(taskReq, claims); err != nil {
		gtils.FailResponse(c, http.StatusInternalServerError, global.ErrUpdated)
	} else {
		gtils.SucResponse(c, "更新成功", "")
	}
}
// @Tags Task
// @Summary 查询任务
// @Description 查询未完成任务以及当天更新的任务
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tasks [get]
func GetTasks(c *gin.Context) {
	customClaims, _ := c.Get("claims")
	claims := customClaims.(*model.CustomClaims)

	tasks,err:=service.GetTasks(claims)
	if err!=nil{
		gtils.FailResponse(c,http.StatusInternalServerError,global.ErrGet)
		return
	}

	retTasks := make([]response.TaskResp,0)
	for _,task := range tasks {
		if  time.Now().Sub(task.UpdatedAt)<=time.Hour*24 {
			retTasks = append(retTasks, service.NewTaskResp(&task))
		}
	}
	gtils.SucResponse(c,"查询成功",retTasks)
}
