package handlers

import (
	"encoding/json"
	"errors"

	"github.com/a11dev/go-gen-backend/internal/data"
	"github.com/a11dev/go-gen-backend/internal/model"
	"github.com/gin-gonic/gin"
)

type URI struct {
	TaskKey int64 `json:"taskkey" uri:"id"`
}

func getFromTaskList(key int64) (model.ResultTaskDetail, error) {
	var info model.ResultTaskDetail
	for _, task := range data.TaskDetails {
		if task.Key == key {
			var f interface{}
			err := json.Unmarshal(task.Body, &f)
			if err != nil {
				return info, err
			}
			info.Key = task.Key
			info.Body = f
			return info, nil
		}

	}
	return info, errors.New("not found")
}

func GetTask(c *gin.Context) {

	uri := URI{}
	if err := c.BindUri(&uri); err != nil {
		c.AbortWithError(400, err)
		return
	}

	info, err := getFromTaskList(uri.TaskKey)
	if err != nil {
		c.JSON(500, err)
		return
	}

	c.JSON(200, info)
}

func GetTasks(c *gin.Context) {
	// name := c.Param("id")
	// j, _ := json.Marshal()

	c.JSON(200,
		data.TaskList,
	)
}
