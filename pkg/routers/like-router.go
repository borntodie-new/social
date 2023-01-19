package routers

import (
	"github.com/jason/social/pkg/controllers"
	"github.com/jason/social/pkg/utils"
	"net/http"
)

// http://127.0.0.1:8080/likes?postId=19	GET		获取某条动态下所有的收藏信息
// http://127.0.0.1:8080/likes				POST	给某条动态收藏
// http://127.0.0.1:8080/likes?postId=1		DELETE 	删除收藏
func registerLikeRouter(mux *http.ServeMux)  {
	mux.HandleFunc("/likes", utils.CORS(controllers.DispatchLike))
}