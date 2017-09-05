package db

import (
	"strconv"
	"database/sql"
	"app/util"
)

/**
 * 分页函数，适用任何表
 * 参数：表名，查询条件，当前页数，页面大小
 * 返回 总记录条数,总页数,以及当前请求的数据*Rows
 */
func GetPagesInfo(sel string, tableName string, conditions string, currentpage int, pagesize int) (int, int, *sql.Rows, error) {
	if currentpage <= 1 {
		currentpage = 1
	}
	if pagesize == 0 {
		pagesize = 20
	}

	var totalItem, totalpages int = 0, 0                                                          //总条数,总页数
	rows, err := db.Query("SELECT COUNT(id) FROM " + tableName + " " + conditions)
	if err != nil {
		return 0, 0, nil, err
	}

	if rows.Next()  {
		err = rows.Scan(&totalItem)//获取总条数
		if err != nil {
			return 0, 0, nil, err
		}
	}
	if totalItem <= pagesize {
		totalpages = 1
	} else if totalItem > pagesize {
		temp := totalItem / pagesize
		if (totalItem % pagesize) != 0 {
			temp = temp + 1
		}
		totalpages = temp
	}
	if currentpage > totalpages {
		return totalItem, totalpages, nil, util.Error(util.PARAMS_ERROR, "超过总页数")
	}
	rows, err = db.Query(sel + " " + tableName + " " + conditions + " LIMIT " + strconv.Itoa((currentpage-1)*pagesize) + "," + strconv.Itoa(pagesize))
	if err != nil {
		return 0, 0, nil, err
	}
	return totalItem, totalpages, rows, nil
}



