package routers

import (
	"github.com/jason/social/pkg/controllers"
	"github.com/jason/social/pkg/utils"
	"net/http"
)


// restful
// 方法   	 路由	  	含义
// get 		/user/2 	获取ID为2的用户信息
// post 	/user   	新增用户
// delete 	/user/2 	删除ID为2的用户信息
// put 		/user/2 	修改ID为2的用户信息

func registerRelationshipRouter(mux *http.ServeMux)  {
	mux.HandleFunc("/relationships", utils.CORS(controllers.DispatchRelationship))
}