package controllers

import (
	"encoding/json"
	"github.com/jason/social/pkg/models"
	"github.com/jason/social/pkg/utils"
	"net/http"
)

// http://127.0.0.1:8080/user POST
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// 1. 判断请求方式
	if r.Method != "POST" {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	// 2. 获取参数
	user := models.User{}
	utils.ParseBody(r, &user)

	// 判断是否登录
	loginUser, exist := utils.CheckToken(r)
	if !exist {
		http.Error(w, "Not log in or not exists", http.StatusUnauthorized)
		return
	}
	user.ID = loginUser.ID
	// 4. 和数据库做交互，修改用户信息
	_, _, err := models.UpdateUser(user)
	if err != nil {
		http.Error(w, "Update user information failed", http.StatusBadRequest)
		return
	}
	updatedUser, _, _ := models.GetUserByUsername(loginUser.Username)
	// 4. 返回结果
	// 4.1 响应数据做序列化
	res, _ := json.Marshal(updatedUser)
	// 4.2 设置content-type
	w.Header().Set("Content-Type", "application/json")
	// 4.3 设置状态码
	w.WriteHeader(http.StatusOK)
	// 4.4 返回结果。
	w.Write(res)
}

// http://127.0.0.1:8080/user/find/12 GET
func GetUser(w http.ResponseWriter, r *http.Request) {
	// 1. 判断请求方式
	if r.Method != "GET" {
		http.Error(w, "Method not allowed!", http.StatusMethodNotAllowed)
		return
	}
	// 2. 获取参数
	userId := utils.ParseParams(r)

	// 3. 与数据库做交互
	user, _, err := models.GetUserById(userId)
	if err != nil {
		http.Error(w, "User not found!", http.StatusBadRequest)
		return
	}

	// 4. 返回结果
	// 4.1 响应数据做序列化
	res, _ := json.Marshal(user)
	// 4.2 设置content-type
	w.Header().Set("Content-Type", "application/json")
	// 4.3 设置状态码
	w.WriteHeader(http.StatusOK)
	// 4.4 返回结果。
	w.Write(res)
}
