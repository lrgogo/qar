package db

import (
	"encoding/json"
	"time"
	"log"
	"app/config"
	"app/util"
)

type UserInfo struct {
	Id 			int64 `json:"id"`
	User_id 	int64 `json:"user_id"`
	Nickname 	string `json:"nickname"`
	Age 		int `json:"age"`
	Sex 		int `json:"sex"`
	Head 		string `json:"head"`
	Interest 	string `json:"interest"`
	Intro 		string `json:"intro"`
}

func GetUserInfo(uid string) (*UserInfo, error) {
	key := config.REDIS_USER_INFO + uid
	val, err := client.Get(key).Result()
	if err == nil { //Redis存在这个key
		info := &UserInfo{}
		err = json.Unmarshal([]byte(val), info)
		if err != nil {
			return nil, err
		}
		log.Println("redis hit uid=", uid)
		return info, err
	}

	rows, err :=
		db.Query(
			"SELECT id, user_id, nickname, age, sex, head, interest, intro FROM user_info WHERE user_id=?", uid)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	u := UserInfo{}
	if rows.Next() {
		err = rows.Scan(&u.Id, &u.User_id, &u.Nickname, &u.Age, &u.Sex, &u.Head, &u.Interest, &u.Intro)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, util.Error(util.USER_UNEXIST, "用户不存在")
	}
	bts, err := json.Marshal(&u)
	if err != nil {
		 return nil, err
	}
	err = client.Set(key, string(bts), 1 * time.Minute).Err()
	if err != nil {
		return nil, err
	}
	return &u, nil
}