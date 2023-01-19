package main

import (
	"github.com/jason/social/pkg/routers"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", upload) // 文件的上传
	mux.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("upload")))) // 文件夹权限的开放
	routers.RegisterSocialRouter(mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func upload(w http.ResponseWriter, r *http.Request)  {
	if r.Method == "POST"{
		// 设置读取文件大小
		r.ParseMultipartForm(32 << 20)
		// 获取上传文件
		file, header, err := r.FormFile("file") // file：文件对象。header：文件信息。err
		if err != nil {
			http.Error(w, "get file error", http.StatusBadRequest)
			return
		}
		defer file.Close()
		log.Println(header.Header)
		// 创建文件夹
		if err := os.Mkdir("./upload", os.ModePerm); err != nil {
			http.Error(w, "create director failed!", http.StatusBadRequest)
			return
		}
		// 创建文件
		fp, err := os.Create("./upload/" + header.Filename)
		if err != nil {
			http.Error(w, "create file object failed!", http.StatusBadRequest)
			return
		}
		defer fp.Close()
		// 保存文件
		_, err = io.Copy(fp, file)
		if err != nil {
			http.Error(w, "upload file failed!", http.StatusBadRequest)
			return
		}
		w.Write([]byte("http://127.0.0.1:8080/image/"+header.Filename))
	}
}