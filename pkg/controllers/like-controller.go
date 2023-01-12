package controllers

import (
	"github.com/jason/social/pkg/models"
	"github.com/jason/social/pkg/utils"
	"net/http"
	"strconv"
)

func DispatchLike(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getLikes(w, r)
	} else if r.Method == "POST" {
		addLike(w, r)
	} else if r.Method == "DELETE" {
		delLike(w, r)
	} else {
		http.Error(w, "method not allowed!", http.StatusMethodNotAllowed)
		return
	}
}

// http://127.0.0.1:8080/likes?postId=19	GET		获取某条动态下所有的收藏信息

func getLikes(w http.ResponseWriter, r *http.Request) {
	postId := r.URL.Query().Get("postId")
	likes, err := models.GetLikes(postId)
	if err != nil{
		http.Error(w, "get likes informations failed!", http.StatusBadRequest)
		return
	}
	utils.JSON(w, likes)
}
// http://127.0.0.1:8080/likes				POST	给某条动态收藏
/*
{
    "postId": 19
}
*/
func addLike(w http.ResponseWriter, r *http.Request) {
	user, exist := utils.CheckToken(r)
	if !exist{
		http.Error(w, "not log in!", http.StatusUnauthorized)
		return
	}
	like := models.Like{}
	utils.ParseBody(r, &like)

	like.UserId = user.ID

	createdLike, _, err := models.CreateLike(like)
	if err != nil {
		http.Error(w, "create like failed!", http.StatusBadRequest)
		return
	}
	utils.JSON(w, createdLike)
}
// http://127.0.0.1:8080/likes?postId=1		DELETE 	删除收藏

func delLike(w http.ResponseWriter, r *http.Request) {
	user, exist := utils.CheckToken(r)
	if !exist{
		http.Error(w, "not log in!", http.StatusUnauthorized)
		return
	}
	postId := r.URL.Query().Get("postId")
	postIdInt, _ := strconv.Atoi(postId)

	like := models.Like{
		UserId: user.ID,
		PostId: postIdInt,
	}

	err := models.DeleteLike(like)
	if err != nil {
		http.Error(w, "delete like failed!", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post has disliked!"))
}