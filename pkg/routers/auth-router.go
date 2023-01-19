package routers

import (
	"github.com/jason/social/pkg/controllers"
	"github.com/jason/social/pkg/utils"
	"net/http"
)

func registerAuthRouter(mux *http.ServeMux) {
	mux.HandleFunc("/login", utils.CORS(controllers.Login))
	mux.HandleFunc("/register", utils.CORS(controllers.Register))
	mux.HandleFunc("/logout", utils.CORS(controllers.Logout))
}
