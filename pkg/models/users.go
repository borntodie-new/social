package models

import (
	"database/sql"
	"github.com/jason/social/pkg/config"
	"strings"
)

var (
	db *sql.DB
)

type User struct {
	ID         int    `json:"id" form:"id"`                 // 用户编号
	Username   string `json:"username" form:"username"`     // 用户名
	Password   string `json:"password" form:"password"`     // 密码
	Email      string `json:"email" form:"email"`           // 邮箱
	Name       string `json:"name" form:"name"`             // 昵称
	CoverPic   string `json:"coverPic" form:"coverPic"`     // 背景图
	ProfilePic string `json:"profilePic" form:"profilePic"` // 头像
	City       string `json:"city" form:"city"`             //  城市
	WebSite    string `json:"webSite" form:"webSite"`       // 个人网站
}

func init() {
	db = config.GetDB()
}

func GetAllUser() ([]User, error) {
	// 1. 准备sql语句
	sqlStr := `select * from users`
	// 2. 预处理SQL
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	// 3. 执行SQL语句
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	users := rows2Users(rows)
	// 4. 返回执行结果
	return users, nil
}

func GetUserById(userId string) (*User, *sql.DB, error) {
	// 1. 准备SQL语句
	sqlStr := "select * from users where id = ?"

	// 2. 预执行SQL
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, nil, err
	}
	defer stmt.Close()
	// 3. 执行SQL
	row := stmt.QueryRow(userId)
	user, err := row2User(row)
	if err != nil {
		return nil, nil, err
	}
	// 4. 返回执行结果
	return user, db, nil
}

func GetUserByUsername(username string) (*User, *sql.DB, error) {
	// 1. 准备SQL语句
	sqlStr := "select * from users where username = ?"

	// 2. 预执行SQL
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, nil, err
	}
	defer stmt.Close()
	// 3. 执行SQL
	row := stmt.QueryRow(username)
	user, err := row2User(row)
	if err != nil {
		return nil, nil, err
	}
	// 4. 返回执行结果
	return user, db, nil
}

func CreateUser(user User) (*User, *sql.DB, error) {
	// 1. 准备SQL语句
	sqlStr := "insert into users(username, email, password, name, coverPic, profilePic, city, website) value(?,?,?,?,?,?,?,?)"
	// 2. 预执行SQL
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, nil, err
	}
	defer stmt.Close()
	// 3. 执行SQL
	result, err := stmt.Exec(user.Username, user.Email, user.Password, user.Name, user.CoverPic, user.ProfilePic, user.City, user.WebSite)
	if err != nil {
		return nil, nil, err
	}

	id, _ := result.LastInsertId()
	user.ID = int(id)
	// 4. 返回执行结果
	return &user, db, nil
}

func UpdateUser(user User) (*User, *sql.DB, error) {
	// 1. 准备SQL语句
	tempSql := "update users set"
	var params []interface{}
	if user.Name != "" {
		tempSql = tempSql + " ,name=?" // update user set name=?
		params = append(params, user.Name)
	}
	if user.City != "" {
		tempSql = tempSql + " ,city=?"
		params = append(params, user.City)
	}
	if user.WebSite != "" {
		tempSql = tempSql + " ,website=?"
		params = append(params, user.WebSite)
	}
	if user.ProfilePic != "" {
		tempSql = tempSql + " ,profilePic=?"
		params = append(params, user.ProfilePic)
	}
	if user.CoverPic != "" {
		tempSql = tempSql + " ,coverPic=?"
		params = append(params, user.CoverPic)
	}
	sqlStr := tempSql + " where id=?" //  update user set name=? where username=?
	params = append(params, user.ID)
	sqlStr = strings.Replace(sqlStr, ",", "", 1)

	// 2. 预执行SQL
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, nil, err
	}
	defer stmt.Close()
	// 3. 执行SQL
	result, err := stmt.Exec(params...)
	if err != nil {
		return nil, nil, err
	}

	id, _ := result.LastInsertId()
	user.ID = int(id)
	// 4. 返回执行结果
	return &user, db, nil
}

func rows2Users(rows *sql.Rows) []User {
	users := make([]User, 0)
	for rows.Next() {
		user := User{}
		_ = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Name, &user.CoverPic, &user.ProfilePic, &user.City, &user.WebSite)
		users = append(users, user)
	}
	return users
}

func row2User(row *sql.Row) (*User, error) {
	user := User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Name, &user.CoverPic, &user.ProfilePic, &user.City, &user.WebSite)
	return &user, err
}
