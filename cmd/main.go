package main

import (
	"github.com/jason/social/pkg/routers"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	routers.RegisterSocialRouter(mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
