package db

import "fmt"

type TempAnswer struct {
	Id          int64 `json:"id"`
	C     string `json:"c"`
}

func GetTempAnswers(currentpage, pagesize int) ([]TempAnswer, error) {
	totalItem, totalPage, rows, err :=
		GetPagesInfo(
			"SELECT id, c FROM",
			"temp_answer",
			"ORDER BY id DESC",
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

	var list []TempAnswer
	for rows.Next() {
		a := TempAnswer{}
		err = rows.Scan(&a.Id, &a.C)
		if err != nil {
			return nil, err
		}

		list = append(list, a)
	}
	return list, nil
}

