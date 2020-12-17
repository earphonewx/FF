package middleware

//func JWTAuthMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		token := c.Request.Header.Get("x-token")
//		if token == "" {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthenticated request"})
//			c.Abort()
//			return
//		}
//		// parseToken 解析token包含的信息
//		claims, err := utils.ParseToken(token)
//		if err != nil {
//			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
//			c.Abort()
//			return
//		}
//		c.Set("claims", claims)
//		c.Next()
//	}
//}
