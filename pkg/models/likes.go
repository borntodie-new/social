package models

import (
	"database/sql"
	"github.com/jason/social/pkg/config"
)

type Like struct {
	ID     int `json:"id" form:"id"`
	UserId int `json:"userId" form:"userId"`
	PostId int `json:"postId" form:"postId"`
}

func init() {
	db = config.GetDB()
}

// http://127.0.0.1:8080/likes?postId=19	GET		获取某条动态下所有的收藏信息
func GetLikes(postId string) ([]Like, error) {
	sqlStr := "select * from likes where postId = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(postId)
	if err != nil {
		return nil, err
	}
	likes := rows2Likes(rows)
	return likes, nil
}
// http://127.0.0.1:8080/likes				POST	给某条动态收藏
func CreateLike(like Like)(*Like, *sql.DB, error) {
	sqlStr := "insert into likes(userId, postId) value(?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil,nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(like.UserId, like.PostId)
	if err != nil {
		return nil, nil, err
	}
	if affected, err := result.RowsAffected(); affected != 1{
		return nil, nil, err
	}
	id, _ := result.LastInsertId()
	like.ID = int(id)
	return &like, db, nil
}
// http://127.0.0.1:8080/likes?postId=1		DELETE 	删除收藏
func DeleteLike(like Like) error {
	sqlStr := "delete from likes where userId = ? and postId = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return  err
	}
	defer stmt.Close()
	result, err := stmt.Exec(like.UserId, like.PostId)
	if err != nil {
		return  err
	}
	if affected, err := result.RowsAffected(); affected != 1{
		return err
	}
	id, _ := result.LastInsertId()
	like.ID = int(id)
	return nil
}

func rows2Likes(rows *sql.Rows) []Like {
	likes := make([]Like, 0)
	for rows.Next() {
		like := Like{}
		_ = rows.Scan(&like.ID, &like.UserId, &like.PostId)
		likes = append(likes, like)
	}
	return likes
}
