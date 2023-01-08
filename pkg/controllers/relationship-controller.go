package controllers

import (
	"github.com/jason/social/pkg/models"
	"github.com/jason/social/pkg/utils"
	"net/http"
	"strconv"
)

// restful
// 方法   	 路由	  	含义
// get 		/user/2 	获取ID为2的用户信息
// post 	/user   	新增用户
// delete 	/user/2 	删除ID为2的用户信息
// put 		/user/2 	修改ID为2的用户信息

func DispatchRelationship(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getRelationship(w, r)
	} else if r.Method == "POST" {
		addRelationship(w, r)
	} else if r.Method == "DELETE" {
		delRelationship(w, r)
	} else {
		http.Error(w, "method not allowed!", http.StatusMethodNotAllowed)
		return
	}
}

// http://127.0.0.1:8080/relationships?followedUserId=6  followedUserId就表示要查谁的粉丝
func getRelationship(w http.ResponseWriter, r *http.Request) {
	followedUserId := r.URL.Query().Get("followedUserId")
	relationships, err := models.GetAllRelationship(followedUserId)
	if err != nil {
		http.Error(w, "get all relationship failed!", http.StatusBadRequest)
		return
	}
	utils.JSON(w, relationships)
}

// http://127.0.0.1:8080/relationships POST
// 李四关注了张三，所以当前应该登录的用户是李四 -> followerUserId 是李四；followedUserId 是张三
func addRelationship(w http.ResponseWriter, r *http.Request) {
	// 验证权限，登录判断
	user, exist := utils.CheckToken(r) // 这里的user是当前登录的用户，就是李四
	if !exist {
		http.Error(w, "Not log in!", http.StatusUnauthorized)
		return
	}
	// 获取数据
	relationship := models.Relationship{}
	utils.ParseBody(r, &relationship)
	relationship.FollowerUserId = user.ID

	// 和数据库做交互
	createRelationship, _, err := models.CreateRelationship(relationship)
	if err != nil {
		http.Error(w, "create relationship failed", http.StatusBadRequest)
		return
	}

	//// 4. 返回结果
	//// 4.1 响应数据做序列化
	//res, _ := json.Marshal(createRelationship)
	//// 4.2 设置content-type
	//w.Header().Set("Content-Type", "application/json")
	//// 4.3 设置状态码
	//w.WriteHeader(http.StatusOK)
	//// 4.4 返回结果。
	//w.Write(res)
	utils.JSON(w, createRelationship)

}

// http://127.0.0.1:8080/relationships?userId=5
// 李四取关了张三 李四的信息从token中获得，张三的信息userId获取
func delRelationship(w http.ResponseWriter, r *http.Request) {
	// 验证权限，登录判断
	user, exist := utils.CheckToken(r) // 这里的user是当前登录的用户，就是李四
	if !exist {
		http.Error(w, "Not log in!", http.StatusUnauthorized)
		return
	}
	// 获取数据
	userId := r.URL.Query().Get("userId")
	followedUserId, _ := strconv.Atoi(userId)

	relationship := models.Relationship{
		FollowerUserId: user.ID,        // 李四的信息
		FollowedUserId: followedUserId, // 张三的信息
	}
	// 和数据库做交互
	err := models.DeleteRelationship(relationship)
	if err != nil {
		http.Error(w, "delete relationship failed!", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Unfollowed"))

}
