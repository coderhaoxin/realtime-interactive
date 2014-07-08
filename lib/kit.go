package lib

import (
	"errors"
	"strconv"
	"strings"
)

func ParseFileName(filename string) (oriFilename, formate string, width, height int, err error) {
	var strs []string
	var extname string

	debug("parse file name, filename:", filename)

	if strings.Contains(filename, ".png") {
		strs = strings.Split(filename, ".png")
		extname = "png"
	} else if strings.Contains(filename, ".jpg") {
		strs = strings.Split(filename, ".jpg")
		extname = "jpg"
	} else if strings.Contains(filename, ".jpeg") {
		strs = strings.Split(filename, ".jpeg")
		extname = "jpeg"
	} else {
		return "", "", 0, 0, errors.New("unknown formate")
	}

	debug("parse file name, strs:", strs)

	if len(strs) != 3 {
		return "", "", 0, 0, errors.New("unknown formate")
	}

	width, height = ParseSize(strings.Replace(strs[1], "_", "", 1))

	return strs[0] + "." + extname, extname, width, height, nil
}

func ParseSize(s string) (width, height int) {
	strs := strings.Split(s, "x")

	if len(strs) != 2 {
		return 100, 100
	}

	var err error

	width, err = strconv.Atoi(strs[0])

	if err != nil {
		width = 100
	}

	height, err = strconv.Atoi(strs[1])

	if err != nil {
		height = 100
	}

	return width, height
}
