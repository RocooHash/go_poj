package service

import (
	"github.com/gin-gonic/gin"
	"go_poj/handler"
	"go_poj/model"
	"go_poj/pkg/errno"
)

func AddUser(c *gin.Context) {
	var user model.User
	if err := c.Bind(&user); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	u := model.User{
		Username: user.Username,
		Password: user.Password,
	}
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}
	if _, err := u.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, user)
}

func SelectUser(c *gin.Context) {
	name := c.Query("Username")
	if name == "" {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}
	var user model.User
	if err := user.SelectUserByName(name); err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	if err := user.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}
	handler.SendResponse(c, nil, user)
}
