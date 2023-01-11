package routers

import (
	"github.com/jason/social/pkg/controllers"
	"net/http"
)

/*
1. 获取所有的动态数据 http://127.0.0.1:8080/posts 		GET
2. 发布一条动态      http://127.0.0.1:8080/posts 		POST 携带动态信息
3. 删除一条动态	    http://127.0.0.1:8080/posts/1	DELETE
*/

func registerPostRouter(mux *http.ServeMux)  {
	mux.HandleFunc("/posts", controllers.DispatchPost)
	mux.HandleFunc("/posts/", controllers.DelPost)
}