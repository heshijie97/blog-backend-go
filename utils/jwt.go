package utils

import (
	"fmt"
	jwt "github.com/golang-jwt/jwt/v4"
	"time"
)

// 定义一个密钥，用于加密jwt
var jwtKey = []byte("119638hxx@")

// 生成jwt
func GenerateJWT(userID string) (string, error) {
	//创建一个新的token对象，指明签名方法和声明
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), //设置过期时间为24小时
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析并验证jwt
func ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保token的签名方法是我们预期的签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	// 获取token中的用户Id
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId := claims["user_id"].(string)
		return userId, nil
	}
	return "", fmt.Errorf("could not parse toekn claims")
}
