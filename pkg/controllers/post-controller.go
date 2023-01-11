package controllers

import (
	"github.com/jason/social/pkg/models"
	"github.com/jason/social/pkg/utils"
	"net/http"
	"strconv"
	"time"
)

/*
1. 获取所有的动态数据 http://127.0.0.1:8080/posts 		GET
2. 发布一条动态      http://127.0.0.1:8080/posts 		POST 携带动态信息
3. 删除一条动态	   http://127.0.0.1:8080/posts/1	DELETE
*/

func DispatchPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getPosts(w, r) // 获取所有动态
	} else if r.Method == "POST" {
		addPost(w, r) // 发布动态
	} else {
		http.Error(w, "method not allowed!", http.StatusMethodNotAllowed)
		return
	}
}

// 删除一条动态	   http://127.0.0.1:8080/posts/1	DELETE
func DelPost(w http.ResponseWriter, r *http.Request) {
	user, exist := utils.CheckToken(r)
	if !exist {
		http.Error(w, "not log in!", http.StatusNotFound)
		return
	}
	postId := utils.ParseParams(r)
	postIdInt, _ := strconv.Atoi(postId)
	post := models.Post{
		ID:          postIdInt,
		UserId:      user.ID,
	}
	if err := models.DeletePost(post); err != nil {
		http.Error(w, "delete post failed!", http.StatusNotFound)
		return
	}
	// 3. 返回删除结果
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte("delete post success!"))
}

// http://127.0.0.1:8080/posts 		GET
func getPosts(w http.ResponseWriter, r *http.Request) {
	// 先想想，那个userId怎么来？我们是不是可以通过token中获取？
	userId := 0
	user, exist := utils.CheckToken(r)
	if exist { // 表示用户已经登录，我们需要携带userId进行查询
		userId = user.ID
	}
	postResults, _, err := models.GetPosts(userId)
	if err != nil {
		http.Error(w, "get all posts failed!", http.StatusNotFound)
		return
	}
	utils.JSON(w, postResults)
}

// http://127.0.0.1:8080/posts 		POST 携带动态信息
/*
{
    "desc": "这是我jason的第一个动态",
    "img": "http://www.baidu.com/image/large-avatar.png"
}
*/
func addPost(w http.ResponseWriter, r *http.Request) {
	user, exist := utils.CheckToken(r)
	if !exist {
		http.Error(w, "not log in!", http.StatusNotFound)
		return
	}
	// 获取到了数据
	post := models.Post{}
	utils.ParseBody(r, &post)
	post.UserId = user.ID
	// 设置创建时间
	cTime := time.Now().Format("2006-01-02 15::04:05")
	post.CreateAt = cTime

	createdPost, _, err := models.CreatePost(post)
	if err != nil {
		http.Error(w, "create post failed", http.StatusBadRequest)
		return
	}
	utils.JSON(w, createdPost)
}
