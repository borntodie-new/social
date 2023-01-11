package models

import (
	"database/sql"
	"fmt"
	"github.com/jason/social/pkg/config"
)

type Post struct {
	ID          int    `json:"id" form:"id"`
	Description string `json:"description" form:"description"`
	Img         string `json:"img" form:"img"`
	UserId      int    `json:"userId" form:"userId"`
	CreateAt    string `json:"createAt" form:"createAt"`
}

type PostResult struct {
	UserId     int    `json:"userId" form:"userId"`
	Name       string `json:"name" form:"name"`
	ProfilePic string `json:"profilePic" form:"profilePic"`
	Post
	//post Post
}

func init() {
	db = config.GetDB()
}

// 1. 根据动态ID获取动态信息
func GetPostByID(postId string) (*Post, *sql.DB, error) {
	sqlStr := "select * from posts where id = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(postId)
	post, err := row2Post(row)
	if err != nil {
		return nil, nil, err
	}
	return &post, db, nil
}

// 2. 获取全部的动态信息：
// 2.1 如果传了用户ID的话，我们就需要查询这个用户发布的动态
// 2.2 如果没有传用户ID，那我们就直接查询全部的动态
func GetPosts(userId int) ([]PostResult, *sql.DB, error) {
	var params []interface{}
	var sqlStr string
	if userId == 0 { // 没有传userId信息，直接查
		sqlStr = "select p.*, u.id as userId, u.name, u.profilePic from posts as p join users as u on (p.userId = u.id) order by p.createAt desc"
	} else { // 查询当前用户发布的动态
		sqlStr = "select p.*, u.id as userId, u.name, u.profilePic from posts as p join users as u on (p.userId = u.id) where p.userId = ? order by p.createAt desc"
		params = append(params, userId)
	}
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(params...) // 这里的rows里面存的可不只是post表结构了的数据了
	if err != nil {
		return nil, nil, err
	}
	postResults := rows2Posts(rows)
	return postResults, db, nil
}

// 3. 发布动态
func CreatePost(post Post) (*Post, *sql.DB, error) { // 这里的post中，ID字段其实是0
	sqlStr := "insert into posts(`desc`, `img`, `userId`, `createAt`) value(?,?,?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(post.Description, post.Img, post.UserId, post.CreateAt)
	if err != nil {
		return nil, nil, err
	}
	id, _ := result.LastInsertId()
	post.ID = int(id)
	return &post, db, nil
}

// 4. 删除动态：根据动态ID和用户ID删除
func DeletePost(post Post) error {
	sqlStr := "delete from posts where id = ? and userId = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, _ := stmt.Exec(post.ID, post.UserId)
	if affected, _ := result.RowsAffected(); affected != 1 {
		return fmt.Errorf("delete error")
	}
	return nil
}

func row2Post(row *sql.Row) (Post, error) {
	post := Post{}
	err := row.Scan(&post.ID, &post.Description, &post.Img, &post.UserId, &post.CreateAt)
	return post, err
}

func rows2Posts(rows *sql.Rows) []PostResult {
	postResult := make([]PostResult, 0)
	for rows.Next() {
		post := PostResult{}
		_ = rows.Scan(&post.ID, &post.Description, &post.Img, &post.UserId, &post.CreateAt, &post.UserId, &post.Name, &post.ProfilePic)
		postResult = append(postResult, post)
	}
	return postResult
}
