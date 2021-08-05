package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nc-77/gtils"
	"net/http"
	"strings"
	"time"
	"todo/global"
	"todo/model"
	"todo/service"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// jwt鉴权取Authorization头部信息  登录时回返回token信息
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			gtils.FailResponse(c, http.StatusUnauthorized, global.ErrAuth)
			c.Abort()
			return
		}
		// 剥离Bear令牌头
		parts := strings.SplitN(token, " ", 2)
		if len(parts) < 2 {
			gtils.FailResponse(c, http.StatusUnauthorized, global.ErrAuth)
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(parts[1])
		if err != nil {
			gtils.FailResponse(c, http.StatusUnauthorized, err)
			c.Abort()
			return
		}
		if err, _ = service.FindUserById(claims.ID); err != nil {
			gtils.FailResponse(c, http.StatusUnauthorized, global.ErrAuth)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

type JWT struct {
	SigningKey  []byte
	Issuer      string
	ExpiresTime time.Duration
}

//goland:noinspection ALL
var (
	tokenExpired     = errors.New("Token is expired")
	tokenNotValidYet = errors.New("Token not active yet")
	tokenMalformed   = errors.New("That's not even a token")
	tokenInvalid     = errors.New("Couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		SigningKey:  []byte("TodoList"),
		Issuer:      "NiC",
		ExpiresTime: time.Hour * 24,
	}

}

// 创建 token
func (j *JWT) CreateToken(claims model.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*model.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, tokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, tokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, tokenNotValidYet
			} else {
				return nil, tokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*model.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, tokenInvalid

	} else {
		return nil, tokenInvalid

	}

}
