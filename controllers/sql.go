package controllers

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type imgtable struct {
	Id       int    `db:"id"`
	Userid   int    `db:"userid"`
	Username string `db:"username"`
	ImgPath  string `db:"imgpath"`
	InUse    int    `db:"in_use"`
}

func uploadImageData(userid int, username string, filename string) bool {
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed    :", err, "\n")
		return false
	}
	sql := "INSERT INTO head_table(id,userid,username,imgpath)values (DEFAULT,?,?,?,0)"
	//执行SQL语句
	r, err := db.Exec(sql, userid, username, filename)
	if err != nil {
		fmt.Println("exec failed   :", err, "\n")
		conn.Rollback()
		return false
	}
	//查询最后一天用户ID，判断是否插入成功
	_, err = r.LastInsertId()
	if err != nil {
		fmt.Println("exec failed   :", err, "\n")
		conn.Rollback()
		return false
	}
	conn.Commit()
	return true
}

func getImagePathByUserId(userid int) []string {
	var imgs []imgtable
	var re []string
	sql := "SELECT id,userid,username,imgpath FROM head_table WHERE userid=? "
	err := db.Select(&imgs, sql, userid)
	if err != nil {
		fmt.Println("exec failed   :", err, "\n")
		return nil
	}
	if imgs == nil {
		return nil
	}
	for key, _ := range imgs {
		re = append(re, imgs[key].ImgPath)
	}
	return re
}

func getImagePathByUserIdInUse(userid int) []string {
	var imgs []imgtable
	var re []string
	sql := "SELECT id,userid,username,imgpath FROM head_table WHERE userid=? AND in_use=1"
	err := db.Select(&imgs, sql, userid)
	if err != nil {
		fmt.Println("exec failed   :", err, "\n")
		return nil
	}
	if imgs == nil {
		return nil
	}
	for key, _ := range imgs {
		re = append(re, imgs[key].ImgPath)
	}
	return re
}

func useHeadImage(imgpath string, userid int) bool {
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin failed    :", err, "\n")
		return false
	}
	sql := "UPDATE head_table SET in_use=? WHERE userid = ?"
	res, err := db.Exec(sql, 0, userid)
	if err != nil {
		fmt.Println("exec failed   :", err, "\n")
		conn.Rollback()
		return false
	}
	_, err = res.RowsAffected()
	if err != nil {
		fmt.Println("exec failed   :", err, "\n")
		conn.Rollback()
		return false
	}
	sql = "UPDATE head_table SET in_use=? WHERE imgpath = ? AND userid=?"
	res, err = db.Exec(sql, 1, imgpath, userid)
	if err != nil {
		fmt.Println("exec failed   :", err, "\n")
		conn.Rollback()
		return false
	}
	_, err = res.RowsAffected()
	if err != nil {
		fmt.Println("exec failed   :", err, "\n")
		conn.Rollback()
		return false
	}
	conn.Commit()
	return true
}
