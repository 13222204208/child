package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	req "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		authHeader := c.Request.Header.Get("Authorization")
		fmt.Println("authHeader:", authHeader)
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "请求头中Token为空",
			})
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "请求头中Token格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("id", mc.ID)
		fmt.Println("id:", mc.ID)
		c.Next() // 后续的处理函数可以用过c.Get("id")来获取当前请求的用户信息
	}
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*req.MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &req.MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return req.MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*req.MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
