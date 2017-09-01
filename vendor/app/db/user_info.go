package db

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
	}
	return &u, nil
}