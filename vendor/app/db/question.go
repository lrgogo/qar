package db

type Question struct {
	Id          int64 `json:"id"`
	User_id     int64 `json:"user_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Update_time string `json:"update_time"`
}

func GetQuestions() []Question {
	rows, err :=
		db.Query(
			"SELECT question.id, question.user_id, question.title, question.content, question.update_time FROM question ORDER BY question.create_time DESC LIMIT 20")
	defer rows.Close()
	checkErr(err)

	var list []Question
	for rows.Next() {
		q := Question{}
		err = rows.Scan(&q.Id, &q.User_id, &q.Title, &q.Content, &q.Update_time)
		checkErr(err)

		list = append(list, q)
	}
	return list
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
