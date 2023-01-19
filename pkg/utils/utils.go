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

func JSON(w http.ResponseWriter, x interface{})  {
	// 4. 返回结果
	// 4.1 响应数据做序列化
	res, _ := json.Marshal(x)
	// 4.2 设置content-type
	w.Header().Set("Content-Type", "application/json")
	// 4.3 设置状态码
	w.WriteHeader(http.StatusOK)
	// 4.4 返回结果。
	w.Write(res)
}

// 跨域处理
func CORS(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("access-control-allow-origin", "*")
		w.Header().Set("access-control-allow-headers", "content-type,accesstoken,x-xsrf-token,authorization,token")
		w.Header().Set("access-control-allow-credentials", "true")
		w.Header().Set("access-control-allow-methods", "post,get,delete,put,options")
		w.Header().Set("access-type", "application/json;charset=utf8-8")
		if r.Method == "OPTIONS"{
			w.WriteHeader(http.StatusNoContent)
			return
		}
		f(w, r)
	}
}