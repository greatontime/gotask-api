package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	"gopkg.in/mgo.v2/bson"

	"github.com/greatontime/gotask-api/common"
	"github.com/greatontime/gotask-api/data"
	"github.com/greatontime/gotask-api/models"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var dataSource NoteResource

	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Note Data", 500)
		return
	}
	noteModel := dataResource.Data
	note := &models.TaskNote{
		TaskId:      bson.ObjectIdHex(noteModel.TaskId),
		Description: noteModel.Description,
	}
	context := NewContext()
	defer context.Close()
	col := context.DbCollection("notes")

	repo := &data.NoteRepository{C: col}
	repo.Create(note)
	j, err := json.Marshal(note)
	if err != nil {
		common.DisplayAppError(w, err, "An unpected error has occurred", 500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}
