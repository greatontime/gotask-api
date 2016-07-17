package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/greatontime/gotask-api/common"
	"github.com/greatontime/gotask-api/controllers"
)

func SetNoteRoutes(router *mux.Router) *mux.Router {
	noteRouter.HandleFunc("/notes", controllers.CreateNote).Method("POST")
	noteRouter.HandleFunc("/notes/{id}", controllers.UpdateNote).Method("PUT")
	noteRouter.HandleFunc("/notes/{id}", controllers.GetNoteByID).Methods("GET")
	noteRouter.HandleFunc("/notes", controllers.GetNotes).Methods("GET")
	noteRouter.HandleFunc("/notes/tasks/{id}", controllers.GetNotesByTask).Methods("GET")
	NoteRouter.HandleFunc("/notes/{id}", controllers.DeleteNote).Methods("DELETE")
	router.PathPrefix("/notes").Handler(negroni.New(
		negroni.HandlerFunc(common.Authorize),
		negroni.Wrap(noteRouter),
	))
	return router
}
