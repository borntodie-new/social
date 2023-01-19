package routers

import (
	"github.com/jason/social/pkg/controllers"
	"github.com/jason/social/pkg/utils"
	"net/http"
)

func registerUserRouter(mux *http.ServeMux) {
	mux.HandleFunc("/user", utils.CORS(controllers.UpdateUser))
	mux.HandleFunc("/user/find/", utils.CORS(controllers.GetUser))
}