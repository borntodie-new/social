package controllers

import (
	"github.com/jason/social/pkg/models"
	"github.com/jason/social/pkg/utils"
	"net/http"
	"time"
)

/*
http://127.0.0.1:8080/comments?postId=20 	GET  获取某条动态下所有的评论信息
http://127.0.0.1:8080/comments				POST 给某条动态评论
*/

func DispatchComment(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getComments(w, r) //
	} else if r.Method == "POST" {
		addComments(w, r) //
	} else {
		http.Error(w, "method not allowed!", http.StatusMethodNotAllowed)
		return
	}
}

//http://127.0.0.1:8080/comments?postId=20 	GET  获取某条动态下所有的评论信息
func getComments(w http.ResponseWriter, r *http.Request) {
	// 1. 获取查询参数
	postId := r.URL.Query().Get("postId")
	// 2. 与数据库做交互，获取数据
	comments, err := models.GetComments(postId)
	if err != nil {
		http.Error(w, "get comments failed!", http.StatusBadRequest)
		return
	}
	// 3. 返回结果
	utils.JSON(w, comments)
}

//http://127.0.0.1:8080/comments				POST 给某条动态评论
/*
{
    "postId": 20,
    "desc": "tank你个傻逼，知道我是谁吗？"
}
*/
func addComments(w http.ResponseWriter, r *http.Request) {
	user, exist := utils.CheckToken(r)
	if !exist {
		http.Error(w, "not log in!", http.StatusUnauthorized)
		return
	}
	comment := models.Comment{}
	utils.ParseBody(r, &comment)
	comment.UserId = user.ID
	// 创建时间
	comment.CreateAt = time.Now().Format("2006-01-02 15:04:05")
	createdComment, _, err := models.CreateComment(comment)
	if err != nil {
		http.Error(w, "create comment failed!", http.StatusBadRequest)
		return
	}
	utils.JSON(w, createdComment)
}
