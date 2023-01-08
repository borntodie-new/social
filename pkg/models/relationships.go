package models

import (
	"database/sql"
	"fmt"
	"github.com/jason/social/pkg/config"
)

type Relationship struct {
	ID             int `json:"id" form:"id"`
	FollowerUserId int `json:"followerUserId" form:"followerUserId"` // 关注人
	FollowedUserId int `json:"followedUserId" form:"followedUserId"` // 被关注人
}

/*
张三：1
李四：2
王五：3
赵六：4
李四关注了张三 FollowerUserId=2；FollowedUserId=1
王五关注了张三 FollowerUserId=3；FollowedUserId=1
赵六关注了张三 FollowerUserId=4；FollowedUserId=1

*/
func init() {
	db = config.GetDB()
}

// 1. 获取某个用户的所有关注人
func GetAllRelationship(followedUserId string) ([]Relationship, error) {
	sqlStr := "select * from relationships where followedUserId = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(followedUserId)
	if err != nil {
		return nil, err
	}
	relationships := rows2Relationship(rows)
	return relationships, nil
}

// 2. 关注某人
func CreateRelationship(relationship Relationship) (*Relationship, *sql.DB, error) {
	sqlStr := "insert into relationships(followerUserId, followedUserId) value(?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return nil, nil, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(relationship.FollowerUserId, relationship.FollowedUserId)
	if err != nil {
		return nil, nil, err
	}
	if affected, err := result.RowsAffected(); affected != 1 || err != nil {
		return nil, nil, err
	}

	id, _ := result.LastInsertId()
	relationship.ID = int(id)
	return &relationship, db, nil
}

// 3. 取消关注
// 李四关注了张三 FollowerUserId=2；FollowedUserId=1
// 李四取关张三
func DeleteRelationship(relationship Relationship) error {
	sqlStr := "delete from relationships where followerUserId = ? and followedUserId = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, _ := stmt.Exec(relationship.FollowerUserId, relationship.FollowedUserId)
	if affected,_ := result.RowsAffected(); affected != 1 {
		return fmt.Errorf("unfollowed failed")
	}
	return nil
}
func rows2Relationship(rows *sql.Rows) []Relationship {
	relationships := make([]Relationship, 0)
	for rows.Next() {
		relationship := Relationship{}
		_ = rows.Scan(&relationship.ID, &relationship.FollowerUserId, &relationship.FollowedUserId)
		relationships = append(relationships, relationship)
	}
	return relationships
}
