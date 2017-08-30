package db

type Question struct {
	Id          int64 `json:"id"`
	User_id     int64 `json:"user_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Update_time string `json:"update_time"`
}

func GetQuestions() ([]Question, error) {
	rows, err :=
		db.Query(
			"SELECT question.id, question.user_id, question.title, question.content, question.update_time FROM question ORDER BY question.create_time DESC LIMIT 20")
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var list []Question
	for rows.Next() {
		q := Question{}
		err = rows.Scan(&q.Id, &q.User_id, &q.Title, &q.Content, &q.Update_time)
		if err != nil {
			return nil, err
		}

		list = append(list, q)
	}
	return list, nil
}

func GetQuestionsTotal() (int, error) {
	rows, err := db.Query("SELECT COUNT(*) FROM question")
	defer rows.Close()
	if err != nil {
		return 0, err
	}

	var total int
	for rows.Next()  {
		err = rows.Scan(&total)
		if err != nil {
			return 0, err
		}
	}
	return total, nil
}
