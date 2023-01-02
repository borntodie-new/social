package controllers

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "登录")
}
func Register(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "注册")
}