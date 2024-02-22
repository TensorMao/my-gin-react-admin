package controllers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-backend/common/request"
	"go-backend/common/response"
	"go-backend/models"
	"go-backend/services"
	"net/http"
	"strconv"
)

func CreateUserHandler(ctx *gin.Context) {
	var user models.User
	// 解析请求体
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用服务层创建用户
	if err := services.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func GetUserByIDHandler(ctx *gin.Context) {
	// 从路径参数中获取用户ID
	userID := ctx.Param("id")

	// 将用户ID转换为uint
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// 调用服务层获取用户
	user, err := services.UserService.GetUserByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func UpdateUserHandler(ctx *gin.Context) {
	var user models.User

	// 解析请求体
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用服务层更新用户
	if err := services.UpdateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUserHandler(ctx *gin.Context) {
	// 从路径参数中获取用户ID
	userID := ctx.Param("id")

	// 将用户ID转换为uint
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// 调用服务层删除用户
	if err := services.DeleteUser(uint(id)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func Login(ctx *gin.Context) {
	var input request.Login
	// 解析请求体
	if err := ctx.ShouldBindJSON(&input); err != nil {
		//utils.HandleValidatorError(ctx, err)
		response.Err(ctx, http.StatusBadRequest, 40000, "Failed to bind login request", nil)
		return
	}
	user, err := services.UserService.Login(input)
	if err != nil {
		response.Err(ctx, http.StatusInternalServerError, 50000, err.Error(), nil)
		return
	} else {
		tokenData, _, err := services.JwtService.CreateToken(services.AppGuardName, user)
		if err != nil {
			response.TokenFail(ctx, err)
			return
		}
		response.Success(ctx, "Login successfully", tokenData)
	}

	/*if !store.Verify(input.CaptchaId, input.Captcha, true) {
		response.Err(ctx, 400, 400, "Captcha err", "")
		return
	}
	*/

}

func Register(ctx *gin.Context) {
	var form request.Register
	if err := ctx.ShouldBindJSON(&form); err != nil {
		response.Err(ctx, http.StatusBadRequest, 40000, request.GetErrorMsg(form, err), nil)
		return
	}

	if user, err := services.UserService.Register(form); err != nil {
		response.Err(ctx, http.StatusInternalServerError, 50000, err.Error(), user)
	} else {
		response.Success(ctx, "Create User Successfully", user)
	}
}

func Logout(c *gin.Context) {
	err := services.JwtService.JoinBlackList(c.Keys["token"].(*jwt.Token))
	if err != nil {
		response.Err(c, http.StatusInternalServerError, 50000, "登出失败", nil)
		return
	}
	response.Success(c, "logout success", nil)
}

func Info(c *gin.Context) {
	user, err := services.UserService.GetUserByID(c.Keys["id"].(uint))
	if err != nil {
		response.Err(c, 500, 50000, "get info fail", nil)
		return
	}
	response.Success(c, "get info success", user)

}
