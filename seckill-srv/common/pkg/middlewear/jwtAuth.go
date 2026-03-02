package middlewear

import "github.com/gin-gonic/gin"

func AuthToken(c *gin.Context) {
	token := c.Request.Header.Get("token")
	if token == "" {
		c.JSON(401, gin.H{
			"code": 401,
			"msg":  "用户认证失败",
		})
	}
	parseToken, err := ParseToken(token)
	if err != nil {
		return
	}
	c.Set("userId", parseToken["userId"].(int64))
}
