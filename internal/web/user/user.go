package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler 在此定义跟 user有关的路由
type UserHandler struct {
}

// 这种写法缺陷：容易被别人注册相同的路由
func (u *UserHandler) RegisterRoutesUser(server *gin.Engine) {
	// 统一处理前缀
	ug := server.Group("/user")
	ug.GET("/profile", u.Profile)
	ug.POST("/signup", u.SignUp)
	ug.POST("/login", u.Login)
	ug.POST("/edit", u.Edit)
}

func (u *UserHandler) SignUp(ctx *gin.Context) {

	type SignUpReq struct {
		Email           string `json:"emailInfo"`
		ConfirmPassword string `json: "confirmPassword"`
		Password        string `json: "passWord"`
	}

	var req SignUpReq
	// Bind 方法会根据 Content-Type 来解析你的数据到 req 里面
	// 解析错了，就会直接写会一个 400 的错误
	if err := ctx.Bind(&req); err != nil {
		return
	}
	ctx.String(http.StatusOK, "hello 你在注册")
	fmt.Printf("%v", req)
	// 这边就是数据库操作
}

func (u *UserHandler) Login(ctx *gin.Context) {

}

func (u *UserHandler) Edit(ctx *gin.Context) {

}

func (u *UserHandler) Profile(ctx *gin.Context) {

}
