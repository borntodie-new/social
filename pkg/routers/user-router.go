package routers

import (
	"github.com/jason/social/pkg/controllers"
	"net/http"
)

func registerUserRouter(mux *http.ServeMux) {
	mux.HandleFunc("/user", controllers.UpdateUser)
	mux.HandleFunc("/user/find/", controllers.GetUser)
}