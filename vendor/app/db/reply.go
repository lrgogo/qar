package db

import (
	"fmt"
)

type Reply struct {
	Id          int64 `json:"id"`
	User_id     int64 `json:"user_id"`
	Answer_id int64 `json:"answer_id"`
	Content     string `json:"content"`
	Update_time string `json:"update_time"`
}

func GetReplys(currentpage, pagesize int) ([]Reply, error) {
	totalItem, totalPage, rows, err :=
		GetPagesInfo(
			"SELECT id, user_id, answer_id, content, update_time FROM",
			"reply",
			"ORDER BY id DESC",
			currentpage,
			pagesize,
		)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	fmt.Println(totalItem, totalPage)

	var list []Reply
	for rows.Next() {
		r := Reply{}
		err = rows.Scan(&r.Id, &r.User_id, &r.Answer_id, &r.Content, &r.Update_time)
		if err != nil {
			return nil, err
		}

		list = append(list, r)
	}
	return list, nil
}
