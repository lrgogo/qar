package util

import (
	"github.com/dgrijalva/jwt-go"
	"app/config"
	"time"
	"errors"
	"net/http"
)

func VerifyJWT(r *http.Request) (int64, error) {
	//解析token
	str := r.Header.Get("X-Access-Token")
	if str == "" {
		return 0, errors.New("没有token")
	}
	t, err := jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token无效")
		}
		return []byte(config.JWT_SECRET_KEY), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := t.Claims.(jwt.MapClaims); ok && t.Valid {
		uid := int64(claims["uid"].(float64))
		exp := int64(claims["exp"].(float64))
		if exp <= time.Now().Unix() {
			return 0, errors.New("token过期")
		}
		return uid, nil
	}
	return 0, errors.New("token无效")
}

func CreateAccessJWT(uid int64) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Hour * 24 * 3).Unix(), //3天
	})
	token, err := t.SignedString([]byte(config.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return token, nil
}

func CreateRefreshJWT(uid int64) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uid": uid,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(), //30天
	})
	token, err := t.SignedString([]byte(config.JWT_SECRET_KEY))
	if err != nil {
		return "", err
	}
	return token, nil
}
