package middleware

// IsAdminAuth 判断权限
/*func IsAdminAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取token信息
		claims, _ := ctx.Get("claims")
		// 获取现在用户信息
		currentUser := claims.(*services.CustomClaims)

		// 判断role权限
		if currentUser.AuthorityId != 1 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "用户没有权限",
			})
			//中断下面中间件
			ctx.Abort()
			return
		}
		//继续执行下面中间件
		ctx.Next()
	}
}*/
