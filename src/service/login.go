package service

import (
	"crow-qin/src/config"
	"crow-qin/src/models"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/toolkits/pkg/logger"
	"strings"
	"time"
)

func Login(loginRequest LoginRequest) (signedToken string, err error) {

	//检查用户密码
	userId, err := CheckUserPassword(loginRequest)

	//创建token
	claims := &JWTClaims{
		UserID:      int(userId),
		Username:    loginRequest.UserName,
		Password:    loginRequest.Password,
		FullName:    loginRequest.UserName,
		Permissions: []string{},
	}
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	logger.Infof("%+v", claims)

	signedToken, err = createToken(claims)
	if err != nil {
		logger.Error(err)
		return signedToken, err
	}
	logger.Info(signedToken)

	//存入redis
	redisKey := fmt.Sprintf("%s%s", tokenKeyPrefix, signedToken)
	err = config.Rdb.Set(config.RedisCtx, redisKey, claims.Username, time.Second*tokenTimeOut).Err()
	if err != nil {
		logger.Info(err)
	}
	return signedToken, nil
}

func CheckUserPassword(loginRequest LoginRequest) (userId int64, err error) {
	passwordMd5 := EncryptWithMd5(loginRequest.Password)
	logger.Info("PasswordMd5 :", passwordMd5)

	userId, err = models.CheckUser(loginRequest.UserName, passwordMd5, config.MyTx)
	if err != nil {
		logger.Error(err)
		return userId, err
	}
	logger.Info(userId)
	return userId, nil
}

//身份验证

func Verify(c *gin.Context) (result bool, userName string, err error) {

	//获取token
	authorization := c.Request.Header.Get("Authorization")
	if len(authorization) == 0 {
		return result, userName, errors.New("获取token失败")
	}

	s2 := strings.SplitN(authorization, " ", 2)
	//fmt.Printf("===================%+v", s2)
	strToken := s2[1]
	//fmt.Printf("===================%+v", strToken)

	if strToken == "" {
		result = false
		return result, "", errors.New("获取token失败")
	}

	//验证token，并获取用户名
	claim, err := verifyAction(strToken)
	if err != nil {
		result = false
		return result, "", nil
	}
	logger.Infof("claim : %+v", claim)

	result = true
	userName = claim.Username

	//对比redis中的token
	redisValue := config.Rdb.Get(config.RedisCtx, tokenKeyPrefix+strToken)
	if redisValue.Val() != userName {
		err = errors.New("用户token验证失败：002")
		return false, userName, err
	}
	logger.Infof("redis中的用户名：%s，token的用户名：%s", redisValue.Val(), claim.Username)

	return result, userName, nil
}

// 解析token，获取用户信息
func verifyAction(strToken string) (claims *JWTClaims, err error) {

	token, err := jwt.ParseWithClaims(strToken, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	//fmt.Printf("==================\n%+v\n==================\n", token)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, err
	}
	err = token.Claims.Valid()
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// 生成token

func createToken(claims *JWTClaims) (signedToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(Secret))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func GetTokenInHeader(c *gin.Context) (strToken string, err error) {
	authorization := c.Request.Header.Get("Authorization")
	logger.Infof("authorization：%+v", authorization)

	s2 := strings.SplitN(authorization, " ", 2)
	logger.Infof("s2：%+v", s2)
	strToken = s2[1]
	if strToken == "" {
		err = errors.New("获取token失败")
		logger.Error(err)
		return strToken, err
	}
	logger.Infof("strToken：%s", strToken)
	return strToken, nil
}

func RefreshToken(strToken string) (newToken string, err error) {

	//从token获取用户信息
	claims, err := verifyAction(strToken)
	if err != nil {
		logger.Error(err)
		return newToken, err
	}
	claims.ExpiresAt = time.Now().Unix() + (claims.ExpiresAt - claims.IssuedAt)

	//生成新token
	newToken, err = createToken(claims)
	if err != nil {
		logger.Error(err)
		return newToken, err
	}

	//存入redis
	redisKey := fmt.Sprintf("%s%s", tokenKeyPrefix, newToken)
	err = config.Rdb.Set(config.RedisCtx, redisKey, claims.Username, time.Second*tokenTimeOut).Err()
	if err != nil {
		logger.Info(err)
	}

	//删除旧token
	err = config.Rdb.Del(config.RedisCtx, tokenKeyPrefix+strToken).Err()
	if err != nil {
		logger.Error(err)
	}
	return newToken, nil
}

func Logout(strToken string) error {

	//删除redis中的token
	err := config.Rdb.Del(config.RedisCtx, tokenKeyPrefix+strToken).Err()
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
