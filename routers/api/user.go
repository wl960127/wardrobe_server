package api

import (
	// "wardrobe_server/middleware/jwt"
	"fmt"
	"log"
	"net/http"
	"wardrobe_server/pkg/app"
	"wardrobe_server/pkg/e"
	"wardrobe_server/pkg/utils"
	userservice "wardrobe_server/service/user_service"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

type auth struct {
	ID       int    `form:"id" `
	Mobile   string `form:"mobile" binding:"required"`
	Password string `form:"password" binding:"required"` //required 表示必填
	Username string `form:"username" `
	Sex      int    `form:"sex" `
}

// Auth  获取TOKEN 需要先注册 .
func Auth(c *gin.Context) {
	appG := app.Gin{C: c}
	var reqInfo auth

	// err := c.ShouldBind(&reqInfo)
	if c.ShouldBind(&reqInfo) != nil {
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	log.Printf("请求数据  %s %s", reqInfo.Mobile, reqInfo.Password)

	valid := validation.Validation{}

	valid.MaxSize(reqInfo.Mobile, 100, "mobile").Message("最长为100字符")
	valid.MaxSize(reqInfo.Password, 100, "password").Message("最长为100字符")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, valid.Errors)
		return
	}

	authService := userservice.User{Mobile: reqInfo.Mobile, Password: reqInfo.Password}
	isExist,data, err := authService.Check()


	if !isExist {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST, map[string]string{
			"token": "",
		})
		return
	}

	authService.ID = data["userId"].(int)



	user, err := authService.Get()

	fmt.Printf("请求Token %d %s %s", user.UserID, user.Mobile, user.Password)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}
	
	token, err := utils.GenerateToken(user.UserID, user.Mobile, user.Password)

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

// AddUser .
func AddUser(c *gin.Context) {
	appG := app.Gin{C: c}
	var reqInfo auth

	if err := c.ShouldBind(&reqInfo); err != nil {
		fmt.Printf("  %s %s %s %s", reqInfo.Mobile, reqInfo.Username, reqInfo.Password, err)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	fmt.Printf("  请求数据  %s %s", reqInfo.Mobile, reqInfo.Password)

	valid := validation.Validation{}
	valid.MaxSize(reqInfo.Mobile, 11, "mobile").Message("最长为11字符")
	valid.MaxSize(reqInfo.Password, 15, "password").Message("最长为15字符")
	valid.MaxSize(reqInfo.Username, 20, "username").Message("最长为20字符")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, valid.Errors)
		return
	}

	userSercice := userservice.User{
		Mobile:   reqInfo.Mobile,
		Username: reqInfo.Username,
		Password: reqInfo.Password,
	}

	if err := userSercice.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_FAIL, err)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, "注册成功")
}

// QureyUser  .
func QureyUser(c *gin.Context) {
	// var claims utils.Claims
	appG := app.Gin{C: c}

	userID := c.GetInt("claimsID")

	fmt.Printf("token   %d", userID)
	userService := userservice.User{ID: userID}

	user, err := userService.Get()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_NOT_EXIST, err.Error)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, gin.H{
		// "info": auth{Username: user.Username,Sex: user.Sex},
		"info": gin.H{
			"username": user.Username,
			"sex":      user.Sex,
			"id":       user.UserID,
		},
	})
}
