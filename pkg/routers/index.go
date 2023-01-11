package routers

import "net/http"

var RegisterSocialRouter = func(mux *http.ServeMux) {
	// 注册认证模块路由
	registerAuthRouter(mux)
	// 注册用户模块路由
	registerUserRouter(mux)
	// 注册关系模块路由
	registerRelationshipRouter(mux)
	// 注册动态模块路由
	registerPostRouter(mux)
	// 注册评论模块路由
	registerCommentRouter(mux)
}
