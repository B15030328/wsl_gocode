package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/julienschmidt/httprouter"
)

// 获取流视频处理函数
func streamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	log.Println(vid)
	video, err := os.Open(VIDEO_DIR + vid)

	if err != nil {
		sendErrorResponse(w, http.StatusInternalServerError, "get stream video error")
		return
	}
	log.Println("获取视频成功")
	// 设置http头部，用以浏览器根据头部的格式解析二进制文件
	w.Header().Set("Content_Type", "video/mp4")
	// 以二进制格式传输
	http.ServeContent(w, r, "", time.Now(), video)

	defer video.Close()
}

// 上传视频处理函数
func uploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("Error when try to get file: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}

	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fn, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}
