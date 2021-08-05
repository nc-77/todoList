package api

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nc-77/gtils"
	"net/http"
	"time"
	"todo/global"
	"todo/model"
	"todo/model/request"
	"todo/model/response"
	"todo/router/middleware"
	"todo/service"
)

// @Tags User
// @Summary 用户注册账号
// @Produce  application/json
// @Param data body request.Register true "用户名, 密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"注册成功"}"
// @Router /user/register [post]
func Register(c *gin.Context) {
	var r request.Register
	if err := c.ShouldBindJSON(&r); err != nil {
		gtils.FailResponse(c, http.StatusBadRequest, global.ErrRegister)
		return
	}
	user := model.User{
		Username: r.Username,
		Password: r.Password,
	}

	if err := service.Register(&user); err != nil {
		gtils.FailResponse(c, http.StatusBadRequest, err)
	} else {
		gtils.SucResponse(c, "创建成功", struct{}{})
	}
}

// @Tags User
// @Summary 用户登录
// @Produce  application/json
// @Param data body request.Login true "用户名, 密码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /user/login [post]
func Login(c *gin.Context) {
	var l request.Login
	if err := c.ShouldBindJSON(&l); err != nil {
		gtils.FailResponse(c, http.StatusBadRequest, global.ErrLogin)
		return
	}
	user := model.User{
		Username: l.Username,
		Password: l.Password,
	}
	if err := service.Login(&user); err != nil {
		gtils.FailResponse(c, http.StatusBadRequest, err)
	} else {
		tokenNext(c, &user)
	}
}

// 登录以后签发token
func tokenNext(c *gin.Context, user *model.User) {
	j := middleware.NewJWT()
	claims := model.CustomClaims{
		ID:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(j.ExpiresTime).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		gtils.FailResponse(c, http.StatusUnauthorized, err)
		return
	}
	loginResp := response.LoginResp{
		ID:        user.ID,
		Username:  user.Username,
		Token:     token,
		ExpiresAt: time.Unix(claims.ExpiresAt, 0).Format("2006-01-02 15:04"),
	}
	gtils.SucResponse(c, "登录成功", loginResp)

}
