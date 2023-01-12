package models

import (
	"database/sql"
	"github.com/jason/social/pkg/config"
)

type Comment struct {
	ID          int    `json:"id" form:"id"`
	Description string `json:"description" form:"description"`
	CreateAt    string `json:"createAt" form:"createAt"`
	UserId      int      `json:"userId" form:"userId"`
	PostId      int    `json:"postId" form:"postId"`
}
type CommentResult struct {
	Name       string `json:"name" form:"name"`
	ProfilePic string `json:"profilePic" form:"profilePic"`
	Comment
}

func init() {
	db = config.GetDB()
}

// 获取某条动态下的评论
func GetComments(postId string) ([]CommentResult, error) {
	sqlStr := "select c.*, u.name, u.profilePic from comments as c join users as u on (u.id = c.userId) where c.postId = ? order by c.createdAt desc"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(postId)
	if err != nil {
		return nil, err
	}
	comments := rows2Comments(rows)
	return comments, nil
}

// 给某条动态添加评论
func CreateComment(comment Comment) (*Comment, *sql.DB, error) {
	sqlStr := "insert into comments(`desc`, `createdAt`, `userId`, `postId`) value(?,?,?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(comment.Description, comment.CreateAt, comment.UserId, comment.PostId)
	if affected, err := result.RowsAffected(); affected != 1 || err != nil {
		return nil, nil, err
	}
	id, _ := result.LastInsertId()
	comment.ID = int(id)
	return &comment, db, nil
}

func rows2Comments(rows *sql.Rows) []CommentResult {
	comments := make([]CommentResult, 0)
	for rows.Next() {
		commentResult := CommentResult{}
		_ = rows.Scan(&commentResult.ID, &commentResult.Description, &commentResult.CreateAt, &commentResult.UserId, &commentResult.PostId, &commentResult.Name, &commentResult.ProfilePic)
		comments = append(comments, commentResult)
	}
	return comments
}
