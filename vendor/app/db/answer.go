package db

import "fmt"

type Answer struct {
	Id          int64 `json:"id"`
	User_id     int64 `json:"user_id"`
	Question_id int64 `json:"question_id"`
	Content     string `json:"content"`
	Update_time string `json:"update_time"`
}

func GetAnswers(currentpage, pagesize int) ([]Answer, error) {
	totalItem, totalPage, rows, err :=
		GetPagesInfo(
			"SELECT id, user_id, question_id, content, update_time FROM",
			"answer",
			"ORDER BY create_time DESC",
			currentpage,
			pagesize,
		)
	if rows != nil {
		defer rows.Close()
	}
	if err != nil {
		return nil, err
	}
	fmt.Println(totalItem, totalPage)

	var list []Answer
	for rows.Next() {
		a := Answer{}
		err = rows.Scan(&a.Id, &a.User_id, &a.Question_id, &a.Content, &a.Update_time)
		if err != nil {
			return nil, err
		}

		list = append(list, a)
	}
	return list, nil
}
