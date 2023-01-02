package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/jason/social/pkg/models"
	"github.com/jason/social/pkg/utils"
	"net/http"
)

// http://127.0.0.1:8080/login POST
func Login(w http.ResponseWriter, r *http.Request) {
	// 1. 判断请求方式
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// 2. 获取参数
	u := models.User{}
	utils.ParseBody(r, &u)

	// 3. 调用controller
	user, _, err := models.GetUserByUsername(u.Username)
	if err != nil {
		http.Error(w, "user not found!", http.StatusBadRequest)
		return
	}
	// 密码验证
	if !VerifyPassword(u.Password, user.Password) {
		http.Error(w, "Wrong password or username!", http.StatusBadRequest)
		return
	}
	// 4. 签发token
	cookie := http.Cookie{
		Name:  "username",
		Value: user.Username,
	}
	w.Header().Set("accessToken", cookie.String()) // 将用户信息存入cookie中
	// 5. 返回请求结果
	data := model2Map(user)
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func Register(w http.ResponseWriter, r *http.Request) {
	// 1. 判断请求方式
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// 2. 获取参数
	u := models.User{}
	utils.ParseBody(r, &u) // zs

	// 3. 调用controller,判断用户名是否已经注册了
	uu, _, _ := models.GetUserByUsername(u.Username) // zs
	if uu.Username == u.Username {
		http.Error(w, "User already exists!", http.StatusBadRequest)
		return
	}
	// 调用controller，执行注册操作
	// 完善用户信息
	u.Password = EncryptPassword(u.Password)
	u.CoverPic = "./assets/upload/default-cover-pic.jpg"
	u.ProfilePic = "./assets/upload/default-profile-pic.jpg"
	user, _, _ := models.CreateUser(u)

	// 5. 返回请求结果
	data := model2Map(user)
	res, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("accessToken", "")
	w.WriteHeader(http.StatusOK)
}

func EncryptPassword(password string) string {
	// password是明文的密码
	md := md5.New()
	md.Write([]byte(password))
	return hex.EncodeToString(md.Sum(nil))
}

func VerifyPassword(password, encryptPassword string) bool {
	// password是明文的密码， encryptPassword是加密后的密码
	return encryptPassword == EncryptPassword(password)
}

func model2Map(user *models.User) map[string]interface{} {
	return map[string]interface{}{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"name":       user.Name,
		"profilePic": user.ProfilePic,
		"coverPic":   user.CoverPic,
		"city":       user.City,
		"website":    user.WebSite,
	}
}
