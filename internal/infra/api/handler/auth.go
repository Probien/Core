package handler

import (
	application "github.com/JairDavid/Probien-Backend/internal/app"
	"github.com/gin-gonic/gin"
)

type IAuthHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	RecoverPassword(c *gin.Context)
}

type AuthHandler struct {
	app application.AuthApp
}

func NewAuthHandler(app application.AuthApp) IAuthHandler {
	return AuthHandler{
		app: app,
	}
}

func (a AuthHandler) Login(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a AuthHandler) Logout(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (a AuthHandler) RecoverPassword(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
