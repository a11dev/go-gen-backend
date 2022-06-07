package handlers

import (
	"encoding/json"
	"strconv"

	"github.com/a11dev/go-gen-backend/internal/model"
	"github.com/gin-gonic/gin"
)

var taskList = []model.Task{
	{978897386, 707427, "Assigned", "AA", "REFID_000001"},
	{978897369, 707426, "Working", "AB", "REFID_000002"},
	{978897180, 707406, "Working", "AB", "REFID_000003"},
	{978897047, 707389, "Deleted", "AA", "REFID_000004"},
	{978897038, 707388, "Deleted", "AB", "REFID_000005"},
	{978896992, 707387, "Deleted", "AC", "REFID_000006"},
	{978896993, 707386, "Deleted", "AA", "REFID_000007"},
	{978888483, 707378, "Deleted", "AC", "REFID_000008"},
	{978888466, 707377, "Assigned", "AC", "REFID_000009"},
	{978888413, 707376, "Working", "AA", "REFID_000010"},
	{978886577, 707374, "Assigned", "AC", "REFID_000011"},
	{978886576, 707373, "Assigned", "AC", "REFID_000012"},
	{978886571, 707372, "Assigned", "AA", "REFID_000013"},
	{978886570, 707371, "Assigned", "AB", "REFID_000014"},
}

func GetTask(c *gin.Context) {
	number, _ := strconv.ParseInt(c.Param("id"), 10, 0)
	var info model.Task
	for _, task := range taskList {
		if task.Number == number {
			info = task
		}
	}

	c.JSON(200, info)
}

func GetTasks(c *gin.Context) {
	// name := c.Param("id")
	j, _ := json.MarshalIndent(taskList, "", "  ")

	c.JSON(200,
		string(j),
	)
}
