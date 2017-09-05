package db

type Question struct {
	Id          int64 `json:"id"`
	User_id     int64 `json:"user_id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Update_time string `json:"update_time"`
}

func GetQuestions(uid int64) ([]Question, error) {
	rows, err :=
		db.Query(
			"SELECT id, user_id, title, content, update_time FROM question WHERE user_id=? ORDER BY id DESC LIMIT 20", uid)
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
	rows, err := db.Query("SELECT COUNT(question.id) FROM question")
	defer rows.Close()
	if err != nil {
		return 0, err
	}

	var total int
	if rows.Next()  {
		err = rows.Scan(&total)
		if err != nil {
			return 0, err
		}
	}
	return total, nil
}
