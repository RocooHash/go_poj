package model

import (
	"fmt"
	"log"
)

func Insert(sql string, args ...interface{}) (int64, error) {
	stmt, err := DB.Prepare(sql)
	defer stmt.Close()

	if err != nil {
		return 1, err
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		return 1, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 1, err
	}
	return id, nil
}

func Delete(sql string, args ...interface{}) {
	stmt, err := DB.Prepare(sql)
	defer stmt.Close()

	CheckErr(err, "Sql设置失败")
	res, err := stmt.Exec(args...)
	CheckErr(err, "参数设置错误")
	num, err := res.RowsAffected()
	CheckErr(err, "删除失败")
	fmt.Printf("成功删除%d行", num)
}

func Update(sql string, args ...interface{}) {
	stmt, err := DB.Prepare(sql)
	defer stmt.Close()

	CheckErr(err, "Sql设置失败")
	res, err := stmt.Exec(args...)
	CheckErr(err, "参数设置错误")
	num, err := res.RowsAffected()
	CheckErr(err, "更新失败")
	fmt.Printf("成功修改%d行", num)
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Panicln(msg, err)
	}
}
