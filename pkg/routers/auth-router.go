package routers

import (
	"github.com/jason/social/pkg/controllers"
	"net/http"
)

func registerAuthRouter(mux *http.ServeMux) {
	mux.HandleFunc("/login", controllers.Login)
	mux.HandleFunc("/register", controllers.Register)
	mux.HandleFunc("/logout", controllers.Logout)
}
