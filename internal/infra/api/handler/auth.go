package handler

import "github.com/gin-gonic/gin"

type IAuthHandler interface {
	Login(c *gin.Context)
	Logout(c *gin.Context)
	RecoverPassword(c *gin.Context)
}

type AuthHandler struct {
}

func NewAuthHandler() IAuthHandler {
	return AuthHandler{}
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
