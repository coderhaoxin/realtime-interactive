package lib

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	. "github.com/visionmedia/go-debug"
)

var debug = Debug("lib")

func Service(w http.ResponseWriter, r *http.Request) {
	debug("RequestURI: %s", r.RequestURI)

	strs := strings.Split(r.RequestURI, "/")

	if len(strs) == 0 {
		Res400(w)
		return
	}

	filename := strs[len(strs)-1]

	if FromCache(w, filename) {
		// cached
		debug("cached")
		return
	}

	debug("filename: %s", filename)

	var parsedFilename, formate string
	var width, height int
	var err error
	if strings.Contains(filename, ".png") {
		parsedFilename, formate, width, height, err = ParseFileName(filename)
		if err != nil {
			Res400(w)
			return
		}
		filename, err = Resize("fixture/"+parsedFilename, "fixture/"+filename, formate, uint(width), uint(height), false)
		if err != nil {
			debug("resize error: %s", err)
			Res500(w)
			return
		}
	} else if strings.Contains(filename, ".jpg") || strings.Contains(filename, ".jpeg") {
		parsedFilename, formate, width, height, err = ParseFileName(filename)
		if err != nil {
			Res400(w)
			return
		}
		filename, err = Resize("fixture/"+parsedFilename, "fixture/"+filename, formate, uint(width), uint(height), false)
		if err != nil {
			debug("resize error: %s", err)
			Res500(w)
			return
		}
	} else {
		Res400(w)
		return
	}

	debug("resized filenpath: %s", filename)
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		Res500(w)
		return
	}

	// response image file
	w.WriteHeader(200)
	w.Header().Set("content-type", "image/"+formate)
	w.Write(bytes)

	return
}

func FromCache(w http.ResponseWriter, filename string) bool {
	fullpath, err := filepath.Abs("fixture/" + filename)

	if err != nil {
		return false
	}

	if _, err := os.Stat(fullpath); os.IsNotExist(err) {
		// response
		bytes, err := ioutil.ReadFile(fullpath)

		if err != nil {
			return false
		}

		_, formate, _, _, err := ParseFileName(filename)

		if err != nil {
			return false
		}

		w.WriteHeader(200)
		w.Header().Set("content-type", "image/"+formate)
		w.Write(bytes)
	}

	return false
}

func Res400(w http.ResponseWriter) {
	w.WriteHeader(400)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "invalid request"}"`))
}

func Res500(w http.ResponseWriter) {
	w.WriteHeader(500)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(`{"message": "server error"}"`))
}
