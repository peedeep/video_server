package handler

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"os"
	"time"
	"io/ioutil"
	"log"
	"io"
	"html/template"
)

func TestPageHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./videos/upload.html")
	t.Execute(w, nil)
}

func StreamHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	vid := p.ByName("video-id")
	vl := VideoPath + vid
	video, err := os.Open(vl)
	if err != nil {
		log.Printf("Error when try to open file: %v", err)
		SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
	}
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "", time.Now(), video)
	defer video.Close()
}

func UploadHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxUploadSize); err != nil {
		SendErrorResponse(w, http.StatusBadRequest, "File is too big")
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("Read file error: %v", err)
		SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	fileName := p.ByName("video-id")
	err = ioutil.WriteFile(VideoPath+fileName, data, 0666)
	if err != nil {
		log.Printf("Write file error: %v", err)
		SendErrorResponse(w, http.StatusInternalServerError, "Internal Error")
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Uploaded successfully")
}
