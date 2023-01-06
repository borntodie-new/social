package utils

import (
	"encoding/json"
	"github.com/jason/social/pkg/models"
	"io/ioutil"
	"net/http"
	"strings"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

// http://127.0.0.1:8080
func ParseParams(r *http.Request) string {
	tempPath := strings.Split(r.URL.Path, "/") // /user/find/12
	if len(tempPath) > 0 {
		return tempPath[len(tempPath)-1]
	}
	return ""
}

func CheckToken(r *http.Request) (*models.User, bool) {
	// 获取cookie
	cookie := r.Header.Get("AccessToken")
	temp := strings.Split(cookie, "=")
	if len(temp) <= 1 {
		return nil, false
	}
	username := strings.ReplaceAll(temp[1], ";", "")
	// 根据cookie里的用户信息查询当前用户
	user, _, err := models.GetUserByUsername(username)
	if err != nil {
		return nil, false
	}
	// 返回结果
	return user, true
}
