package utils

/*func CreateToken(c *gin.Context, Id uint, NickName string, Role int) string {
	//生成token信息
	j := middleware.NewJWT()
	claims := middleware.CustomClaims{
		ID:          Id,
		Username:    NickName,
		AuthorityId: uint(Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			// TODO 设置token过期时间
			ExpiresAt: time.Now().Unix() + global.GlobConfig.JWT.JwtTTL, //token -->1天过期
			Issuer:    "services-login",
		},
	}
	//生成token
	token, err := j.CreateToken(claims)
	if err != nil {
		response.Success(c, 401, "Fail to generate the token", "")
		return ""
	}
	return token
}*/
