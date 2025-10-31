package controller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kevinjosephdavis/chatroom/common/message"
	"github.com/kevinjosephdavis/chatroom/internal/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

// Register 处理用户注册请求
func (ctrl *UserController) Register(c *gin.Context) {
	var req message.RegisterRequest

	// 1. 绑定并验证请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, message.BaseResponse{
			Code:      400,
			Message:   "请求参数错误:" + err.Error(),
			Timestamp: time.Now().Unix(),
		})
		return
	}

	// 2. 调用业务逻辑
	user, err := ctrl.userService.Register(req.Username, req.Password, req.Email, req.Nickname)
	if err != nil {
		c.JSON(http.StatusBadRequest, message.BaseResponse{
			Code:      400,
			Message:   "注册失败:" + err.Error(),
			Timestamp: time.Now().Unix(),
		})
		return
	}

	// 3. 返回成功响应
	c.JSON(http.StatusOK, message.BaseResponse{
		Code:    200,
		Message: "注册成功",
		Data: message.RegisterResponse{
			AccountID: user.AccountID,
			Username:  user.Username,
			Email:     user.Email,
		},
		Timestamp: time.Now().Unix(),
	})
}

// Login 处理用户登录请求
func (ctrl *UserController) Login(c *gin.Context) {
	var req message.LoginRequest

	// 1. 绑定并验证请求参数
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, message.BaseResponse{
			Code:      400,
			Message:   "请求参数错误:" + err.Error(),
			Timestamp: time.Now().Unix(),
		})
		return
	}

	// 2. 调用业务逻辑
	user, err := ctrl.userService.Login(req.AccountID, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, message.BaseResponse{
			Code:      401,
			Message:   "登录失败:" + err.Error(),
			Timestamp: time.Now().Unix(),
		})
		return
	}

	// 3. 返回成功响应
	c.JSON(http.StatusOK, message.BaseResponse{
		Code:    200,
		Message: "登录成功",
		Data: message.LoginResponse{
			AccountID: user.AccountID,
			UserID:    user.ID,
			Username:  user.Username,
			Nickname:  user.Nickname,
			Avatar:    user.Avatar,
			//Token和ExpiredAt稍后添加
		},
		Timestamp: time.Now().Unix(),
	})
}
