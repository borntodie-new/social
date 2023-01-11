package routers

import (
	"github.com/jason/social/pkg/controllers"
	"net/http"
)


/*
http://127.0.0.1:8080/comments?postId=20 	GET  获取某条动态下所有的评论信息
http://127.0.0.1:8080/comments				POST 给某条动态评论
*/
func registerCommentRouter(mux *http.ServeMux)  {
	mux.HandleFunc("/comments", controllers.DispatchComment)
}