package db

import (
	"errors"
)

type LoginInfo struct {
	Id 			int64 `json:"id"`
	Mobile 		string `json:"mobile"`
	Pwd 		string `json:"-"`
	Vip 		int `json:"vip"`
	Token 		string `json:"token"`
}

func GetLoginInfo(mobile, pwd string) (*LoginInfo, error) {
	rows, err :=
		db.Query(
			"SELECT id, mobile, pwd, vip FROM user WHERE mobile=?", mobile)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	i := LoginInfo{}
	if rows.Next() {
		err = rows.Scan(&i.Id, &i.Mobile, &i.Pwd, &i.Vip)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("该用户未注册")
	}
	if i.Pwd != pwd {
		return nil, errors.New("密码错误")
	}
	return &i, nil
}