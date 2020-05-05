package streaming

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

var (
	ID = regexp.MustCompile("streaming/[0-9]+/")
)

func StreamHandler(w http.ResponseWriter, r *http.Request) {
	req := ID.FindString(r.URL.String())
	if req == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	url := strings.Split(r.URL.String(), "/")
	mediaBase := "video/" + url[5]
	if url[len(url)-1] == "stream" {
		serveHlsM3u8(w, r, mediaBase, "index.m3u8")
	} else {
		serveHlsTs(w, r, mediaBase, url[len(url)-1])
	}
}

func serveHlsM3u8(w http.ResponseWriter, r *http.Request, mediaBase, m3u8Name string) {
	mediaFile := fmt.Sprintf("%s/%s", mediaBase, m3u8Name)
	w.Header().Set("Content-Type", "application/x-mpegURL")
	http.ServeFile(w, r, mediaFile)
}

func serveHlsTs(w http.ResponseWriter, r *http.Request, mediaBase, segName string) {
	mediaFile := fmt.Sprintf("%s/%s", mediaBase, segName)
	w.Header().Set("Content-Type", "video/MP2T")
	http.ServeFile(w, r, mediaFile)
}
