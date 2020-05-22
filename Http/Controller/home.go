package Controller

import (
	"fmt"

	"github.com/ZongzhiCui/go_gin/Http/Model"
	"github.com/ZongzhiCui/go_gin/common"
	"gopkg.in/go-playground/validator.v8"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required,min=6,max=12"`
}

// 绑定模型获取验证错误的方法
func (r *LoginForm) GetError(err validator.ValidationErrors) string {

	// 这里的 "LoginRequest.Mobile" 索引对应的是模型的名称和字段
	if val, ok := err["LoginForm.User"]; ok {
		if val.Field == "User" {
			switch val.Tag {
			case "required":
				return "请输入手机号码"
			}
		}
	}
	if val, ok := err["LoginForm.Password"]; ok {
		if val.Field == "Password" {
			switch val.Tag {
			case "required":
				return "请输入验证码"
			case "min":
				return "密码最少6位"
			case "max":
				return "密码最大12位"
			}

		}
	}
	return "参数错误"
}

func LoginEndpoint(c *gin.Context) {
	user := c.PostForm("user")
	password := c.DefaultPostForm("password", "15")

	data := map[string]interface{}{
		"param1": user,
		"param2": password,
	}
	var form LoginForm

	//fmt.Println(c.Request.Header.Get("Content-Type"))
	// 在这种情况下，将自动选择合适的绑定
	if e := c.ShouldBind(&form); e == nil {

		//连接数据库
		userInfo := Model.User_info()
		//fmt.Printf("%+v \n", userInfo)

		if form.User == userInfo.Name && form.Password == userInfo.Password {
			common.OutputJson(c, 200, "you are logged in", data)
		} else {
			common.OutputJson(c, 200, "unauthorized", nil)
		}
	} else {
		//fmt.Printf("c.ShouldBind(&form) error is : [%T] %v \n", e, e)
		common.OutputJson(c, 200, form.GetError(e.(validator.ValidationErrors)), nil)
	}
}

func Create_user(c *gin.Context) {
	user := c.PostForm("user")
	password := c.DefaultPostForm("password", "15")

	//test写入数据
	var u Model.User
	u.Name = user
	u.Password = password
	r := Model.User_create(u)
	fmt.Println(r)
	if r {
		common.OutputJson(c, 200, "数据写入失败", nil)
	}
	common.OutputJson(c, 200, "OK", u)
}

func Ping(c *gin.Context) {
	common.OutputJson(c, 200, "pong", []int{1, 2})
}
