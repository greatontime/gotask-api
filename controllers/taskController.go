package controllers

import (
	"encoding/json"
	"net/http"

	httpcontext "github.com/gorilla/context"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"

	"github.com/greatontime/gotask-api/common"

	"github.com/greatontime/gotask-api/data"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var dataResource TaskResource
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Task Data",
			500,
		)
		return
	}
	task := &dataResource.Data
	context := NewContext()
	defer context.Close()
	if val, ok := httpcontext.GetOk(r, "user"); ok {
		context.User = val.(string)
	}
	task.CreatedBy = context.User
	col := context.DbCollection("tasks")
	repo := &data.TaskRepository{C: col}
	repo.Create(task)
	j, err := json.Marshal(TaskResource{Data: *task})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}
func GetTasks(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("tasks")
	repo := &data.TaskRepository{C: col}
	tasks := repo.GetAll()
	j, err := json.Marshal(TaskResource{Data: tasks})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
