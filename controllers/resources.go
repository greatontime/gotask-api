package controllers

import (
	"github.com/greatontime/gotask-api/models"
)

type (
	UserResource struct {
		Data models.User `json:"data"`
	}
	LoginResource struct {
		Data LoginModel `json:"data"`
	}
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}
	TaskResource struct {
		Data models.Task `json:"data"`
	}
	TasksResource struct {
		Data []models.Task `json:"data"`
	}
	NoteResource struct {
		Data NoteModel `json:"data"`
	}
	NotesResource struct {
		Data []models.TaskNote `json:"data"`
	}
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
	NoteModel struct {
		TaksId      string `json:"taksid"`
		Description string `json:"description"`
	}
)
