package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"os"
	"time"
	"fmt"
	"io/ioutil"
	"log"
	"io"
)

func StreamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("vid-id")
	vl := VIDEO_DIR + vid
	fmt.Println(vl)
	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Open file err :%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	video.Close()
}

func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//校验文件大小
	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		sendErrorResponse(w, http.StatusBadRequest, "File is too large")
		return
	}
	//表单里拿出文件
	file, _, err := r.FormFile("file") // <form> name="file" ... </form>
	if err != nil {
		log.Printf("Form file err :%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file err :%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fn, data, 0666)
	if err != nil {
		log.Printf("Write file err:%v", err)
		sendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}
